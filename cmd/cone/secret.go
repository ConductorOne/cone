package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"filippo.io/age"
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

const (
	displayNameFlag    = "display-name"
	contentFlag        = "content"
	fileFlag           = "file"
	outFlag            = "out"
	inputFormatFlag    = "input-format"
	allowedUserIdsFlag = "allowed-user-ids"
	allowedEmailsFlag  = "allowed-emails"
	expiresInFlag      = "expires-in"
	maxViewsFlag       = "max-views"
)

func secretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Create and retrieve secrets",
	}

	cmd.AddCommand(secretCreateCmd())
	cmd.AddCommand(secretGetCmd())
	cmd.AddCommand(secretViewCmd())
	cmd.AddCommand(secretDownloadCmd())

	return cmd
}

func secretCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a TEXT or FILE secret shared with team members or external recipients",
		Long: "Create a secret and encrypt its content client-side (Age) to the recipient the API returns.\n\n" +
			"Recipients (exactly one is required):\n" +
			"  --allowed-user-ids   share with C1 team members who sign in with SSO\n" +
			"  --allowed-emails     share with external recipients who verify via email\n\n" +
			"Content:\n" +
			"  --content   inline TEXT content; --input-format hints the viewer (plaintext, json, yaml, key-value)\n" +
			"  --file       path to a FILE to upload; its content type, size, and name are derived automatically\n" +
			"  (omit both)  create an empty secret and print the vault id, Age recipient, and upload URL for\n" +
			"               out-of-band upload\n\n" +
			"--content and --file are mutually exclusive, as are --allowed-user-ids and --allowed-emails.",
		RunE: secretCreateRun,
	}

	cmd.Flags().String(displayNameFlag, "", "A cleartext label for the secret, visible in the My Secrets view.")
	cmd.Flags().String(contentFlag, "", "The plaintext TEXT content to encrypt and store (max 64KB encrypted).")
	cmd.Flags().String(fileFlag, "", "Path to a file to upload as a FILE secret (max 1GB). Mutually exclusive with --content.")
	cmd.Flags().String(inputFormatFlag, "plaintext", "TEXT format hint for the viewer UI: plaintext, json, yaml, or key-value. Ignored for files.")
	cmd.Flags().StringSlice(allowedUserIdsFlag, nil, "C1 user IDs allowed to view this secret (1 to 128). Mutually exclusive with --allowed-emails.")
	cmd.Flags().StringSlice(allowedEmailsFlag, nil, "External email addresses allowed to view this secret (1 to 64). Mutually exclusive with --allowed-user-ids.")
	cmd.Flags().String(expiresInFlag, "", "Required. Duration after which the secret content expires (e.g. 3600s).")
	cmd.Flags().Int64(maxViewsFlag, 0, "Maximum number of views before the secret is burned (0 = unlimited).")

	return cmd
}

func secretCreateRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(0, args, cmd); err != nil {
		return err
	}

	userIDs := v.GetStringSlice(allowedUserIdsFlag)
	emails := v.GetStringSlice(allowedEmailsFlag)
	expiresIn := v.GetString(expiresInFlag)
	content := v.GetString(contentFlag)
	filePath := v.GetString(fileFlag)
	if err := validateSecretCreateInput(userIDs, emails, content, filePath, expiresIn); err != nil {
		return err
	}

	// Read the file up-front to fail fast and capture its original metadata. FileSize must
	// be the original (pre-encryption) size: the create request requires it, but content can
	// only be Age-encrypted after the API returns the recipient, so it is not the ciphertext size.
	var fileBytes []byte
	params := createSecretParams{
		userIDs:     userIDs,
		emails:      emails,
		expiresIn:   expiresIn,
		displayName: v.GetString(displayNameFlag),
		inputFormat: v.GetString(inputFormatFlag),
		isFile:      filePath != "",
	}
	if mv := v.GetInt64(maxViewsFlag); mv > 0 {
		params.maxViews = &mv
	}
	if params.isFile {
		fileBytes, err = os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read --%s: %w", fileFlag, err)
		}
		params.filename = filepath.Base(filePath)
		params.contentType = mime.TypeByExtension(filepath.Ext(filePath))
		if params.contentType == "" {
			params.contentType = "application/octet-stream"
		}
		params.fileSize = int64(len(fileBytes))
	}

	createResp, err := createSecret(ctx, c, params)
	if err != nil {
		return err
	}

	vaultID := client.StringFromPtr(createResp.VaultID)
	if vaultID == "" {
		return fmt.Errorf("secret was created but no vault id was returned")
	}

	outputManager := output.NewManager(ctx, v)

	// Empty TEXT secret: return the vault id, Age recipient, and upload URL for out-of-band upload.
	if !params.isFile && content == "" {
		resp := SecretCreateResult(*createResp)
		return outputManager.Output(ctx, &resp, output.WithTransposeTable())
	}

	recipientKey := client.StringFromPtr(createResp.AgeRecipient)
	if recipientKey == "" {
		return fmt.Errorf("secret %s was created but the API returned no Age recipient for encryption", vaultID)
	}

	if params.isFile {
		uploadURL := client.StringFromPtr(createResp.UploadURL)
		if uploadURL == "" {
			return fmt.Errorf("file secret %s was created but the API returned no upload URL", vaultID)
		}
		encrypted, err := encryptBytesToAgeRecipient(recipientKey, fileBytes)
		if err != nil {
			return fmt.Errorf("failed to encrypt file content: %w", err)
		}
		if err := c.UploadSecretFile(ctx, uploadURL, encrypted); err != nil {
			return fmt.Errorf("secret %s was created but uploading the file failed: %w", vaultID, err)
		}
	} else {
		encrypted, err := encryptToAgeRecipient(recipientKey, []byte(content))
		if err != nil {
			return fmt.Errorf("failed to encrypt secret content: %w", err)
		}
		if err := c.SetSecretTextContent(ctx, vaultID, encrypted, setContentInputFormat(params.inputFormat)); err != nil {
			return fmt.Errorf("secret %s was created but uploading content failed: %w", vaultID, err)
		}
	}

	// Re-fetch so the displayed metadata reflects the uploaded content.
	secret, err := c.GetSecret(ctx, vaultID)
	if err != nil {
		return err
	}

	resp := Secret(*secret)
	return outputManager.Output(ctx, &resp, output.WithTransposeTable())
}

// validateSecretCreateInput enforces the create command's flag rules: exactly one recipient
// set (team user IDs xor external emails), a required expiry, and at most one content source.
func validateSecretCreateInput(userIDs, emails []string, content, filePath, expiresIn string) error {
	switch {
	case len(userIDs) == 0 && len(emails) == 0:
		return fmt.Errorf("one of --%s or --%s is required", allowedUserIdsFlag, allowedEmailsFlag)
	case len(userIDs) > 0 && len(emails) > 0:
		return fmt.Errorf("--%s and --%s are mutually exclusive (team vs external sharing)", allowedUserIdsFlag, allowedEmailsFlag)
	}

	if expiresIn == "" {
		return fmt.Errorf("--%s is required (e.g. 3600s)", expiresInFlag)
	}

	if content != "" && filePath != "" {
		return fmt.Errorf("--%s and --%s are mutually exclusive", contentFlag, fileFlag)
	}
	return nil
}

// secretCreator is the subset of client.C1Client that createSecret needs, narrowed so the
// request-building logic can be exercised with a lightweight fake in tests.
type secretCreator interface {
	CreateInternalSecret(ctx context.Context, req *shared.PaperSecretServiceCreateInternalRequest) (*shared.PaperSecretServiceCreateResponse, error)
	CreateExternalSecret(ctx context.Context, req *shared.PaperSecretServiceCreateExternalRequest) (*shared.PaperSecretServiceCreateResponse, error)
}

// createSecretParams carries the inputs for building a create request. Exactly one of
// userIDs (internal/team sharing) or emails (external sharing) is set.
type createSecretParams struct {
	userIDs     []string
	emails      []string
	expiresIn   string
	displayName string
	maxViews    *int64
	inputFormat string
	isFile      bool
	filename    string
	contentType string
	fileSize    int64
}

// createSecret builds and sends the appropriate create request: external (allowedEmails)
// when emails are provided, otherwise internal (allowedUserIds). FILE secrets set the file
// metadata; TEXT secrets set the input-format hint.
func createSecret(ctx context.Context, c secretCreator, p createSecretParams) (*shared.PaperSecretServiceCreateResponse, error) {
	var displayName *string
	if p.displayName != "" {
		displayName = &p.displayName
	}

	if len(p.emails) > 0 {
		secretType := shared.PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeText
		if p.isFile {
			secretType = shared.PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeFile
		}
		req := &shared.PaperSecretServiceCreateExternalRequest{
			SecretType:    &secretType,
			AllowedEmails: p.emails,
			ExpiresIn:     &p.expiresIn,
			DisplayName:   displayName,
			MaxViews:      p.maxViews,
		}
		if p.isFile {
			req.ContentType = &p.contentType
			req.Filename = &p.filename
			req.FileSize = &p.fileSize
		} else {
			req.InputFormat = createExternalInputFormat(p.inputFormat)
		}
		return c.CreateExternalSecret(ctx, req)
	}

	secretType := shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeText
	if p.isFile {
		secretType = shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeFile
	}
	req := &shared.PaperSecretServiceCreateInternalRequest{
		SecretType:     &secretType,
		AllowedUserIds: p.userIDs,
		ExpiresIn:      &p.expiresIn,
		DisplayName:    displayName,
		MaxViews:       p.maxViews,
	}
	if p.isFile {
		req.ContentType = &p.contentType
		req.Filename = &p.filename
		req.FileSize = &p.fileSize
	} else {
		req.InputFormat = createInputFormat(p.inputFormat)
	}
	return c.CreateInternalSecret(ctx, req)
}

func secretGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <vault-id>",
		Short: "Get a secret's metadata by its vault id",
		RunE:  secretGetRun,
	}
}

func secretGetRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	secret, err := c.GetSecret(ctx, args[0])
	if err != nil {
		return err
	}

	resp := Secret(*secret)
	return output.NewManager(ctx, v).Output(ctx, &resp, output.WithTransposeTable())
}

// encryptBytesToAgeRecipient encrypts plaintext to an Age X25519 recipient and returns the
// raw ciphertext bytes, which begin with the Age header "age-encryption.org/v1". FILE secrets
// PUT these bytes verbatim to the upload URL.
func encryptBytesToAgeRecipient(recipientKey string, plaintext []byte) ([]byte, error) {
	recipient, err := age.ParseX25519Recipient(recipientKey)
	if err != nil {
		return nil, fmt.Errorf("invalid Age recipient: %w", err)
	}

	buf := &bytes.Buffer{}
	w, err := age.Encrypt(buf, recipient)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(plaintext); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// encryptToAgeRecipient encrypts plaintext to an Age X25519 recipient and returns the
// base64-encoded ciphertext. The TEXT content API stores content in a protobuf bytes field,
// which is base64-encoded over JSON; the decoded bytes are the raw Age format.
func encryptToAgeRecipient(recipientKey string, plaintext []byte) (string, error) {
	raw, err := encryptBytesToAgeRecipient(recipientKey, plaintext)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(raw), nil
}

func secretViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view <vault-id>",
		Short: "Fetch and decrypt a TEXT secret's content",
		Long: "Fetch a TEXT secret's content and decrypt it locally. A one-time Age identity is " +
			"generated for the request; the server re-encrypts the content to it. This may count " +
			"against the secret's view limit.",
		RunE: secretViewRun,
	}
}

func secretViewRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	// Generate an ephemeral identity; the server re-encrypts content to its recipient.
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return fmt.Errorf("failed to generate Age identity: %w", err)
	}

	content, err := c.GetSecretContent(ctx, args[0], identity.Recipient().String())
	if err != nil {
		return err
	}

	if content.DownloadURL != nil && *content.DownloadURL != "" {
		return fmt.Errorf("secret is a FILE secret; download and decrypt it from: %s", *content.DownloadURL)
	}

	encrypted := client.StringFromPtr(content.EncryptedContent)
	if encrypted == "" {
		return fmt.Errorf("secret has no text content to view")
	}

	plaintext, err := decryptFromAgeIdentity(identity, encrypted)
	if err != nil {
		return fmt.Errorf("failed to decrypt secret content: %w", err)
	}

	fmt.Println(plaintext)
	return nil
}

// decryptBytesFromAgeIdentity decrypts raw Age ciphertext bytes with the given identity.
func decryptBytesFromAgeIdentity(identity *age.X25519Identity, raw []byte) ([]byte, error) {
	r, err := age.Decrypt(bytes.NewReader(raw), identity)
	if err != nil {
		return nil, err
	}

	plaintext := &bytes.Buffer{}
	if _, err := io.Copy(plaintext, r); err != nil {
		return nil, err
	}
	return plaintext.Bytes(), nil
}

// decryptFromAgeIdentity base64-decodes the Age ciphertext and decrypts it with the given identity.
func decryptFromAgeIdentity(identity *age.X25519Identity, encrypted string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("invalid base64 content: %w", err)
	}

	plaintext, err := decryptBytesFromAgeIdentity(identity, raw)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func secretDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download <vault-id>",
		Short: "Download and decrypt a FILE secret's content to disk",
		Long: "Fetch a FILE secret's content and decrypt it locally. A one-time Age identity is " +
			"generated for the request; the server re-encrypts the file to it. Writes to --out, or to " +
			"the secret's original filename in the current directory. This may count against the " +
			"secret's view limit.",
		RunE: secretDownloadRun,
	}

	cmd.Flags().String(outFlag, "", "Output path for the decrypted file (default: the secret's original filename).")

	return cmd
}

func secretDownloadRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	// Generate an ephemeral identity; the server re-encrypts the file to its recipient.
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return fmt.Errorf("failed to generate Age identity: %w", err)
	}

	content, err := c.GetSecretContent(ctx, args[0], identity.Recipient().String())
	if err != nil {
		return err
	}

	downloadURL := client.StringFromPtr(content.DownloadURL)
	if downloadURL == "" {
		if client.StringFromPtr(content.EncryptedContent) != "" {
			return fmt.Errorf("secret is a TEXT secret; use `cone secret view` instead")
		}
		return fmt.Errorf("secret has no downloadable file content")
	}

	encrypted, err := c.DownloadSecretFile(ctx, downloadURL)
	if err != nil {
		return err
	}

	plaintext, err := decryptBytesFromAgeIdentity(identity, encrypted)
	if err != nil {
		return fmt.Errorf("failed to decrypt file content: %w", err)
	}

	outPath := v.GetString(outFlag)
	if outPath == "" {
		outPath = client.StringFromPtr(content.Filename)
	}
	if outPath == "" {
		return fmt.Errorf("could not determine output path; pass --%s", outFlag)
	}

	if err := os.WriteFile(outPath, plaintext, 0o600); err != nil {
		return fmt.Errorf("failed to write %s: %w", outPath, err)
	}

	fmt.Printf("Wrote %d bytes to %s\n", len(plaintext), outPath)
	return nil
}

func createInputFormat(name string) *shared.PaperSecretServiceCreateInternalRequestInputFormat {
	var format shared.PaperSecretServiceCreateInternalRequestInputFormat
	switch name {
	case "json":
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatJSON
	case "yaml":
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatYaml
	case "key-value":
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

func createExternalInputFormat(name string) *shared.PaperSecretServiceCreateExternalRequestInputFormat {
	var format shared.PaperSecretServiceCreateExternalRequestInputFormat
	switch name {
	case "json":
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatJSON
	case "yaml":
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatYaml
	case "key-value":
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

func setContentInputFormat(name string) *shared.PaperSecretServiceSetTextContentRequestInputFormat {
	var format shared.PaperSecretServiceSetTextContentRequestInputFormat
	switch name {
	case "json":
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatJSON
	case "yaml":
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatYaml
	case "key-value":
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

// SecretCreateResult renders the response from creating a secret: the vault id plus the
// Age recipient and upload URL needed to encrypt and upload content out-of-band.
type SecretCreateResult shared.PaperSecretServiceCreateResponse

func (r *SecretCreateResult) Header() []string {
	return []string{
		"Vault Id",
		"Age Recipient",
		"Upload URL",
	}
}

func (r *SecretCreateResult) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.VaultID),
		client.StringFromPtr(r.AgeRecipient),
		client.StringFromPtr(r.UploadURL),
	}}
}

type Secret shared.PaperSecret

func (r *Secret) Header() []string {
	return []string{
		"Display Name",
		"Type",
		"Status",
		"Share URL",
		"Max Views",
		"Current Views",
		"Created At",
		"Content Expires At",
	}
}

func (r *Secret) WideHeader() []string {
	return append([]string{"Vault Id"}, r.Header()...)
}

func (r *Secret) rows() []string {
	var secretType, status string
	if r.SecretType != nil {
		secretType = string(*r.SecretType)
	}
	if r.Status != nil {
		status = string(*r.Status)
	}

	maxViews := ""
	if r.MaxViews != nil {
		maxViews = strconv.FormatInt(*r.MaxViews, 10)
	}
	currentViews := ""
	if r.CurrentViews != nil {
		currentViews = strconv.FormatInt(*r.CurrentViews, 10)
	}

	return []string{
		client.StringFromPtr(r.DisplayName),
		secretType,
		status,
		client.StringFromPtr(r.ShareURL),
		maxViews,
		currentViews,
		output.FormatTime(r.CreatedAt),
		output.FormatTime(r.ContentExpiresAt),
	}
}

func (r *Secret) Rows() [][]string {
	return [][]string{r.rows()}
}

func (r *Secret) WideRows() [][]string {
	return [][]string{append([]string{client.StringFromPtr(r.VaultID)}, r.rows()...)}
}
