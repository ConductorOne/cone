package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
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
	inputFormatFlag    = "input-format"
	allowedUserIdsFlag = "allowed-user-ids"
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

	return cmd
}

func secretCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an internal TEXT secret",
		Long: "Create an internal TEXT secret. When --content is provided it is encrypted client-side " +
			"(Age) to the recipient the API returns and uploaded. Without --content the secret is created " +
			"empty and its vault id and Age recipient are returned for out-of-band upload.",
		RunE: secretCreateRun,
	}

	cmd.Flags().String(displayNameFlag, "", "A cleartext label for the secret, visible in the My Secrets view.")
	cmd.Flags().String(contentFlag, "", "The plaintext secret content to encrypt and store (max 64KB encrypted).")
	cmd.Flags().String(inputFormatFlag, "plaintext", "Format hint for the viewer UI: plaintext, json, yaml, or key-value.")
	cmd.Flags().StringSlice(allowedUserIdsFlag, nil, "Required. C1 user IDs allowed to view this secret (1 to 128).")
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

	// AllowedUserIds and ExpiresIn are required by the API.
	allowed := v.GetStringSlice(allowedUserIdsFlag)
	if len(allowed) == 0 {
		return fmt.Errorf("--%s is required (1 to 128 C1 user IDs)", allowedUserIdsFlag)
	}
	expiresIn := v.GetString(expiresInFlag)
	if expiresIn == "" {
		return fmt.Errorf("--%s is required (e.g. 3600s)", expiresInFlag)
	}

	secretType := shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeText
	req := &shared.PaperSecretServiceCreateInternalRequest{
		SecretType:     &secretType,
		InputFormat:    createInputFormat(v.GetString(inputFormatFlag)),
		AllowedUserIds: allowed,
		ExpiresIn:      &expiresIn,
	}

	if displayName := v.GetString(displayNameFlag); displayName != "" {
		req.DisplayName = &displayName
	}
	if maxViews := v.GetInt64(maxViewsFlag); maxViews > 0 {
		req.MaxViews = &maxViews
	}

	createResp, err := c.CreateInternalSecret(ctx, req)
	if err != nil {
		return err
	}

	vaultID := client.StringFromPtr(createResp.VaultID)
	if vaultID == "" {
		return fmt.Errorf("secret was created but no vault id was returned")
	}

	outputManager := output.NewManager(ctx, v)

	// Without content, return the vault id and Age recipient for out-of-band upload.
	content := v.GetString(contentFlag)
	if content == "" {
		resp := SecretCreateResult(*createResp)
		return outputManager.Output(ctx, &resp, output.WithTransposeTable())
	}

	recipientKey := client.StringFromPtr(createResp.AgeRecipient)
	if recipientKey == "" {
		return fmt.Errorf("secret %s was created but the API returned no Age recipient for encryption", vaultID)
	}

	encrypted, err := encryptToAgeRecipient(recipientKey, []byte(content))
	if err != nil {
		return fmt.Errorf("failed to encrypt secret content: %w", err)
	}

	if err := c.SetSecretTextContent(ctx, vaultID, encrypted, setContentInputFormat(v.GetString(inputFormatFlag))); err != nil {
		return fmt.Errorf("secret %s was created but uploading content failed: %w", vaultID, err)
	}

	// Re-fetch so the displayed metadata reflects the uploaded content.
	secret, err := c.GetSecret(ctx, vaultID)
	if err != nil {
		return err
	}

	resp := Secret(*secret)
	return outputManager.Output(ctx, &resp, output.WithTransposeTable())
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

// encryptToAgeRecipient encrypts plaintext to an Age X25519 recipient and returns the
// base64-encoded ciphertext. The API stores content in a protobuf bytes field, which is
// base64-encoded over JSON; the decoded bytes are the raw Age format beginning with
// "age-encryption.org/v1".
func encryptToAgeRecipient(recipientKey string, plaintext []byte) (string, error) {
	recipient, err := age.ParseX25519Recipient(recipientKey)
	if err != nil {
		return "", fmt.Errorf("invalid Age recipient: %w", err)
	}

	buf := &bytes.Buffer{}
	w, err := age.Encrypt(buf, recipient)
	if err != nil {
		return "", err
	}
	if _, err := w.Write(plaintext); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
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

// decryptFromAgeIdentity base64-decodes the Age ciphertext and decrypts it with the given identity.
func decryptFromAgeIdentity(identity *age.X25519Identity, encrypted string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("invalid base64 content: %w", err)
	}

	r, err := age.Decrypt(bytes.NewReader(raw), identity)
	if err != nil {
		return "", err
	}

	plaintext := &bytes.Buffer{}
	if _, err := io.Copy(plaintext, r); err != nil {
		return "", err
	}
	return plaintext.String(), nil
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
