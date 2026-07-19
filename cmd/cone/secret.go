package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"filippo.io/age"
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

const (
	displayNameFlag     = "display-name"
	labelFlag           = "label"
	contentFlag         = "content"
	contentFileFlag     = "content-file"
	fileFlag            = "file"
	inputFormatFlag     = "input-format"
	formatFlag          = "format"
	userFlag            = "user"
	allowedUserIDsFlag  = "allowed-user-ids"
	externalEmailsFlag  = "external-emails"
	externalEmailFlag   = "external-email"
	expiresInFlag       = "expires-in"
	maxViewsFlag        = "max-views"
	viewLimitFlag       = "view-limit"
	outputFileFlag      = "output"
	pageSizeFlag        = "page-size"
	secretStatusFlag    = "status"
	secretTypeFlag      = "type"
	secretSharingFlag   = "sharing-mode"
	defaultSecretExpiry = "1w"
	formatPlaintext     = "plaintext"
	formatJSON          = "json"
	formatYAML          = "yaml"
	formatKeyValue      = "key-value"
)

func secretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Create and retrieve secrets",
	}

	cmd.AddCommand(secretCreateCmd())
	cmd.AddCommand(secretListCmd())
	cmd.AddCommand(secretGetCmd())
	cmd.AddCommand(secretViewCmd())
	cmd.AddCommand(secretDownloadCmd())
	cmd.AddCommand(secretRevokeCmd())
	cmd.AddCommand(secretAuditCmd())

	return cmd
}

func secretCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a TEXT or FILE secret shared with team members or external recipients",
		Long: "Create a secret and encrypt its content client-side (Age) to the recipient the API returns.\n\n" +
			"Recipients (exactly one is required):\n" +
			"  --user               share with C1 team members who sign in with SSO; accepts user ID,\n" +
			"                       exact email, or a search query that resolves to one enabled user\n" +
			"  --allowed-user-ids   compatibility alias for internal recipient user IDs\n" +
			"  --external-emails    share with external recipients who verify via email\n\n" +
			"Content:\n" +
			"  --content        inline TEXT content; prefer --content-file or stdin for sensitive values\n" +
			"  --content-file   path to TEXT content, or '-' to read from stdin\n" +
			"  --file           path to a FILE to upload; its content type, size, and name are derived automatically\n\n" +
			"--content, --content-file, and --file are mutually exclusive, as are internal and external recipients.",
		RunE: secretCreateRun,
	}

	cmd.Flags().String(displayNameFlag, "", "A cleartext label for the secret, visible in the My Secrets view.")
	cmd.Flags().String(labelFlag, "", "Alias for --display-name.")
	cmd.Flags().String(contentFlag, "", "The plaintext TEXT content to encrypt and store. Prefer --content-file or stdin to avoid shell history.")
	cmd.Flags().String(contentFileFlag, "", "Path to plaintext TEXT content, or '-' to read from stdin.")
	cmd.Flags().String(fileFlag, "", "Path to a file to upload as a FILE secret (max 1GB). Mutually exclusive with --content.")
	cmd.Flags().String(inputFormatFlag, formatPlaintext, "TEXT format hint for the viewer UI: plaintext, json, yaml, or key-value/env. Ignored for files.")
	cmd.Flags().String(formatFlag, "", "Alias for --input-format.")
	cmd.Flags().StringSlice(userFlag, nil, "Internal team member recipient. Accepts user ID, exact email, or a search query that resolves to one enabled user.")
	cmd.Flags().StringSlice(allowedUserIDsFlag, nil, "C1 user IDs allowed to view this secret (1 to 128). Mutually exclusive with external recipients.")
	cmd.Flags().StringSlice(externalEmailsFlag, nil, "External recipient email addresses (1 to 64). Mutually exclusive with internal recipients.")
	cmd.Flags().StringSlice(externalEmailFlag, nil, "Alias for --external-emails.")
	cmd.Flags().String(expiresInFlag, defaultSecretExpiry, "Duration after which the secret content expires (1h to 30d, e.g. 1w, 3600s).")
	cmd.Flags().Int64(maxViewsFlag, 0, "Maximum number of views before the secret is burned (0 = unlimited).")
	cmd.Flags().Int64(viewLimitFlag, 0, "Alias for --max-views.")

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

	legacyUserIDs := cleanStringSlice(v.GetStringSlice(allowedUserIDsFlag))
	userRefs := cleanStringSlice(v.GetStringSlice(userFlag))
	emailFlags := append(v.GetStringSlice(externalEmailsFlag), v.GetStringSlice(externalEmailFlag)...)
	emails := cleanEmailSlice(emailFlags)
	filePath := v.GetString(fileFlag)
	contentFilePath := v.GetString(contentFileFlag)
	hasInlineContent := cmd.Flags().Changed(contentFlag)

	displayName, err := selectedStringFlag(cmd, v, displayNameFlag, labelFlag)
	if err != nil {
		return err
	}
	inputFormat, err := selectedStringFlag(cmd, v, inputFormatFlag, formatFlag)
	if err != nil {
		return err
	}
	inputFormat, err = normalizeSecretInputFormat(inputFormat)
	if err != nil {
		return err
	}
	expiresIn, err := normalizeSecretDuration(v.GetString(expiresInFlag))
	if err != nil {
		return err
	}
	maxViews, err := selectedInt64Flag(cmd, v, maxViewsFlag, viewLimitFlag)
	if err != nil {
		return err
	}

	input := secretCreateInput{
		userIDs:          legacyUserIDs,
		userRefs:         userRefs,
		emails:           emails,
		hasInlineContent: hasInlineContent,
		contentFilePath:  contentFilePath,
		filePath:         filePath,
		expiresIn:        expiresIn,
		maxViews:         maxViews,
	}
	if err := validateSecretCreateInput(input); err != nil {
		return err
	}

	userIDs, err := resolveSecretUserIDs(ctx, c, legacyUserIDs, userRefs)
	if err != nil {
		return err
	}

	content := ""
	params := createSecretParams{
		userIDs:     userIDs,
		emails:      emails,
		expiresIn:   expiresIn,
		displayName: displayName,
		inputFormat: inputFormat,
		isFile:      filePath != "",
	}
	if maxViews > 0 {
		params.maxViews = &maxViews
	}
	switch {
	case params.isFile:
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return fmt.Errorf("failed to stat --%s: %w", fileFlag, err)
		}
		if fileInfo.IsDir() {
			return fmt.Errorf("--%s must be a file, got directory %q", fileFlag, filePath)
		}
		if fileInfo.Size() == 0 {
			return fmt.Errorf("--%s must not be empty", fileFlag)
		}
		if fileInfo.Size() > maxFileSecretBytes {
			return fmt.Errorf("--%s is %d bytes; maximum is %d bytes", fileFlag, fileInfo.Size(), maxFileSecretBytes)
		}
		params.filename = filepath.Base(filePath)
		params.contentType = mime.TypeByExtension(filepath.Ext(filePath))
		if params.contentType == "" {
			params.contentType = "application/octet-stream"
		}
		params.fileSize = fileInfo.Size()
	case contentFilePath != "":
		content, err = readSecretContentFile(cmd, contentFilePath)
		if err != nil {
			return err
		}
	default:
		content = v.GetString(contentFlag)
	}

	if !params.isFile {
		if err := validateSecretTextContent(content, inputFormat); err != nil {
			return err
		}
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

	recipientKey := client.StringFromPtr(createResp.AgeRecipient)
	if recipientKey == "" {
		return fmt.Errorf("secret %s was created but the API returned no Age recipient for encryption", vaultID)
	}

	if params.isFile {
		uploadURL := client.StringFromPtr(createResp.UploadURL)
		if uploadURL == "" {
			return fmt.Errorf("file secret %s was created but the API returned no upload URL", vaultID)
		}
		encryptedFile, err := encryptFileToTemp(filePath, recipientKey)
		if err != nil {
			return fmt.Errorf("failed to encrypt file content: %w", err)
		}
		defer func() { _ = encryptedFile.Close() }()
		defer func() { _ = os.Remove(encryptedFile.Name()) }()
		encryptedInfo, err := encryptedFile.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat encrypted file: %w", err)
		}
		if err := c.UploadSecretFileReader(ctx, uploadURL, encryptedFile, encryptedInfo.Size()); err != nil {
			return fmt.Errorf("secret %s was created but uploading the file failed: %w", vaultID, err)
		}
	} else {
		encryptedBytes, err := encryptBytesToAgeRecipient(recipientKey, []byte(content))
		if err != nil {
			return fmt.Errorf("failed to encrypt secret content: %w", err)
		}
		if len(encryptedBytes) > maxTextSecretCiphertextBytes {
			return fmt.Errorf("encrypted text content is %d bytes; maximum is %d bytes", len(encryptedBytes), maxTextSecretCiphertextBytes)
		}
		encrypted := base64.StdEncoding.EncodeToString(encryptedBytes)
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

const (
	maxTextSecretPlaintextBytes  = 60 * 1024
	maxTextSecretCiphertextBytes = 64 * 1024
	maxFileSecretBytes           = int64(1 << 30)
	minSecretExpiry              = time.Hour
	maxSecretExpiry              = 30 * 24 * time.Hour
)

type secretCreateInput struct {
	userIDs          []string
	userRefs         []string
	emails           []string
	hasInlineContent bool
	contentFilePath  string
	filePath         string
	expiresIn        string
	maxViews         int64
}

func validateSecretCreateInput(input secretCreateInput) error {
	internalRecipients := len(input.userIDs) + len(input.userRefs)
	switch {
	case internalRecipients == 0 && len(input.emails) == 0:
		return fmt.Errorf("one of --%s, --%s, or --%s is required", userFlag, allowedUserIDsFlag, externalEmailsFlag)
	case internalRecipients > 0 && len(input.emails) > 0:
		return fmt.Errorf("internal recipients (--%s/--%s) and external recipients (--%s/--%s) are mutually exclusive", userFlag, allowedUserIDsFlag, externalEmailsFlag, externalEmailFlag)
	case internalRecipients > 128:
		return fmt.Errorf("internal recipient count is %d; maximum is 128", internalRecipients)
	case len(input.emails) > 64:
		return fmt.Errorf("external recipient count is %d; maximum is 64", len(input.emails))
	}
	if err := validateEmailSlice(input.emails); err != nil {
		return err
	}

	contentSources := 0
	if input.hasInlineContent {
		contentSources++
	}
	if input.contentFilePath != "" {
		contentSources++
	}
	if input.filePath != "" {
		contentSources++
	}
	switch contentSources {
	case 0:
		return fmt.Errorf("one of --%s, --%s, or --%s is required", contentFlag, contentFileFlag, fileFlag)
	case 1:
	default:
		return fmt.Errorf("--%s, --%s, and --%s are mutually exclusive", contentFlag, contentFileFlag, fileFlag)
	}

	if input.expiresIn == "" {
		return fmt.Errorf("--%s is required", expiresInFlag)
	}
	if input.maxViews < 0 || input.maxViews > 1000 {
		return fmt.Errorf("--%s must be between 0 and 1000", maxViewsFlag)
	}
	return nil
}

func validateSecretTextContent(content string, inputFormat string) error {
	if content == "" {
		return fmt.Errorf("text content must not be empty")
	}
	if len([]byte(content)) > maxTextSecretPlaintextBytes {
		return fmt.Errorf("text content is %d bytes; maximum is %d bytes", len([]byte(content)), maxTextSecretPlaintextBytes)
	}
	if inputFormat == formatJSON && !json.Valid([]byte(content)) {
		return fmt.Errorf("invalid JSON content")
	}
	return nil
}

func selectedStringFlag(cmd *cobra.Command, v interface{ GetString(string) string }, primary string, alias string) (string, error) {
	primaryValue := v.GetString(primary)
	aliasValue := v.GetString(alias)
	primaryChanged := cmd.Flags().Changed(primary)
	aliasChanged := cmd.Flags().Changed(alias)
	if primaryChanged && aliasChanged && primaryValue != aliasValue {
		return "", fmt.Errorf("--%s and --%s cannot both be set to different values", primary, alias)
	}
	if aliasChanged {
		return aliasValue, nil
	}
	return primaryValue, nil
}

func selectedInt64Flag(cmd *cobra.Command, v interface{ GetInt64(string) int64 }, primary string, alias string) (int64, error) {
	primaryValue := v.GetInt64(primary)
	aliasValue := v.GetInt64(alias)
	primaryChanged := cmd.Flags().Changed(primary)
	aliasChanged := cmd.Flags().Changed(alias)
	if primaryChanged && aliasChanged && primaryValue != aliasValue {
		return 0, fmt.Errorf("--%s and --%s cannot both be set to different values", primary, alias)
	}
	if aliasChanged {
		return aliasValue, nil
	}
	return primaryValue, nil
}

func normalizeSecretDuration(input string) (string, error) {
	d, err := strToDur(input)
	if err != nil {
		return "", fmt.Errorf("invalid --%s: %w", expiresInFlag, err)
	}
	if d == nil {
		return "", fmt.Errorf("--%s is required", expiresInFlag)
	}
	if *d < minSecretExpiry || *d > maxSecretExpiry {
		return "", fmt.Errorf("--%s must be between 1h and 30d", expiresInFlag)
	}
	return fmt.Sprintf("%ds", int64(d.Seconds())), nil
}

func normalizeSecretInputFormat(input string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "", formatPlaintext, "plain", "text":
		return formatPlaintext, nil
	case formatJSON:
		return formatJSON, nil
	case formatYAML, "yml":
		return formatYAML, nil
	case formatKeyValue, "key_value", "keyvalue", "env":
		return formatKeyValue, nil
	default:
		return "", fmt.Errorf("--%s must be one of plaintext, json, yaml, or env", inputFormatFlag)
	}
}

func readSecretContentFile(cmd *cobra.Command, path string) (string, error) {
	var data []byte
	var err error
	if path == "-" {
		data, err = io.ReadAll(cmd.InOrStdin())
	} else {
		//nolint:gosec // path comes directly from the CLI flag.
		data, err = os.ReadFile(path)
	}
	if err != nil {
		return "", fmt.Errorf("failed to read --%s: %w", contentFileFlag, err)
	}
	return string(data), nil
}

func cleanStringSlice(values []string) []string {
	out := make([]string, 0, len(values))
	seen := map[string]struct{}{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}

func cleanEmailSlice(values []string) []string {
	out := make([]string, 0, len(values))
	seen := map[string]struct{}{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		addr, err := mail.ParseAddress(value)
		if err != nil || addr.Address == "" {
			out = append(out, value)
			continue
		}
		email := strings.ToLower(addr.Address)
		if _, ok := seen[email]; ok {
			continue
		}
		seen[email] = struct{}{}
		out = append(out, email)
	}
	return out
}

func validateEmailSlice(emails []string) error {
	for _, email := range emails {
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address == "" || strings.Contains(addr.Address, " ") {
			return fmt.Errorf("invalid external email address %q", email)
		}
	}
	return nil
}

func resolveSecretUserIDs(ctx context.Context, c client.C1Client, legacyUserIDs []string, userRefs []string) ([]string, error) {
	userIDs := cleanStringSlice(legacyUserIDs)
	seen := map[string]struct{}{}
	for _, userID := range userIDs {
		seen[userID] = struct{}{}
	}
	for _, ref := range userRefs {
		userID, err := resolveSecretUserID(ctx, c, ref)
		if err != nil {
			return nil, err
		}
		if _, ok := seen[userID]; ok {
			continue
		}
		seen[userID] = struct{}{}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}

func resolveSecretUserID(ctx context.Context, c client.C1Client, ref string) (string, error) {
	if looksLikeC1ID(ref) {
		return ref, nil
	}
	users, err := c.SearchUsers(ctx, secretUserSearchRequest(ref))
	if err != nil {
		return "", err
	}
	switch len(users) {
	case 0:
		return "", fmt.Errorf("no enabled user matched %q", ref)
	case 1:
	default:
		return "", fmt.Errorf("%q matched multiple enabled users; pass a C1 user ID or exact email", ref)
	}
	userID := client.StringFromPtr(users[0].ID)
	if userID == "" {
		return "", fmt.Errorf("matched user for %q had no id", ref)
	}
	return userID, nil
}

func secretUserSearchRequest(ref string) *shared.SearchUsersRequest {
	pageSize := 2
	status := shared.SearchUsersRequestUserStatusesEnabled
	req := &shared.SearchUsersRequest{
		PageSize:     &pageSize,
		UserStatuses: []shared.SearchUsersRequestUserStatuses{status},
	}
	if strings.Contains(ref, "@") {
		req.Email = &ref
		return req
	}
	req.Query = &ref
	return req
}

func looksLikeC1ID(ref string) bool {
	if len(ref) < 20 {
		return false
	}
	for _, r := range ref {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
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

func secretListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List secrets you created",
		RunE:  secretListRun,
	}
	cmd.Flags().String(queryFlag, "", "Fuzzy search by display name.")
	cmd.Flags().String(secretStatusFlag, "active", "Status filter: active, expired, burned, revoked, data-deleted, or all.")
	cmd.Flags().String(secretTypeFlag, "all", "Type filter: text, file, or all.")
	cmd.Flags().String(secretSharingFlag, "all", "Sharing mode filter: internal, external, or all.")
	cmd.Flags().Int(pageSizeFlag, 100, "Page size for API requests.")
	return cmd
}

func secretListRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(0, args, cmd); err != nil {
		return err
	}

	pageSize := v.GetInt(pageSizeFlag)
	if pageSize <= 0 || pageSize > 1000 {
		return fmt.Errorf("--%s must be between 1 and 1000", pageSizeFlag)
	}
	req := &shared.PaperSecretServiceSearchMySecretsRequest{
		PageSize: &pageSize,
	}
	sortBy := shared.PaperSecretServiceSearchMySecretsRequestSortBySearchSortByCreatedDesc
	req.SortBy = &sortBy
	if query := strings.TrimSpace(v.GetString(queryFlag)); query != "" {
		req.Query = &query
	}
	statuses, err := secretListStatuses(v.GetString(secretStatusFlag))
	if err != nil {
		return err
	}
	req.Statuses = statuses
	secretType, err := secretListType(v.GetString(secretTypeFlag))
	if err != nil {
		return err
	}
	req.SecretType = secretType
	sharingMode, err := secretListSharingMode(v.GetString(secretSharingFlag))
	if err != nil {
		return err
	}
	req.SharingMode = sharingMode

	secrets, err := c.SearchMySecrets(ctx, req)
	if err != nil {
		return err
	}

	resp := Secrets(secrets)
	return output.NewManager(ctx, v).Output(ctx, &resp)
}

func secretGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <vault-id|share-code|share-url>",
		Short: "Get a secret's metadata",
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

	secret, err := resolveSecret(ctx, c, args[0])
	if err != nil {
		return err
	}

	resp := Secret(*secret)
	return output.NewManager(ctx, v).Output(ctx, &resp, output.WithTransposeTable())
}

// encryptBytesToAgeRecipient encrypts plaintext to an Age hybrid recipient and returns the
// raw ciphertext bytes, which begin with the Age header "age-encryption.org/v1". FILE secrets
// PUT these bytes verbatim to the upload URL.
func encryptBytesToAgeRecipient(recipientKey string, plaintext []byte) ([]byte, error) {
	recipient, err := age.ParseHybridRecipient(recipientKey)
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

func encryptFileToTemp(filePath string, recipientKey string) (*os.File, error) {
	recipient, err := age.ParseHybridRecipient(recipientKey)
	if err != nil {
		return nil, fmt.Errorf("invalid Age recipient: %w", err)
	}

	//nolint:gosec // path comes directly from the CLI flag.
	src, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = src.Close() }()

	tmp, err := os.CreateTemp("", "cone-secret-upload-*.age")
	if err != nil {
		return nil, err
	}
	ok := false
	defer func() {
		if !ok {
			_ = tmp.Close()
			_ = os.Remove(tmp.Name())
		}
	}()

	w, err := age.Encrypt(tmp, recipient)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(w, src); err != nil {
		_ = w.Close()
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	if _, err := tmp.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	ok = true
	return tmp, nil
}

// encryptToAgeRecipient encrypts plaintext to an Age hybrid recipient and returns the
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
		Use:   "view <vault-id|share-code|share-url>",
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
	identity, err := age.GenerateHybridIdentity()
	if err != nil {
		return fmt.Errorf("failed to generate Age identity: %w", err)
	}

	secret, err := resolveSecret(ctx, c, args[0])
	if err != nil {
		return err
	}
	if isExternalSecret(secret) {
		return fmt.Errorf("external secret content must be opened with the external share URL; cone supports metadata for external secrets but not the opener reveal flow")
	}
	vaultID, err := secretVaultID(secret)
	if err != nil {
		return err
	}
	content, err := c.GetSecretContent(ctx, vaultID, identity.Recipient().String())
	if err != nil {
		return err
	}
	if content == nil {
		return fmt.Errorf("secret content response was empty")
	}

	if content.DownloadURL != nil && *content.DownloadURL != "" {
		return fmt.Errorf("secret is a FILE secret; use cone secret download")
	}

	encrypted := client.StringFromPtr(content.EncryptedContent)
	if encrypted == "" {
		return fmt.Errorf("secret has no text content to view")
	}

	plaintext, err := decryptFromAgeIdentity(identity, encrypted)
	if err != nil {
		return fmt.Errorf("failed to decrypt secret content: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), plaintext); err != nil {
		return err
	}
	return nil
}

func secretDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download <vault-id|share-code|share-url>",
		Short: "Fetch and decrypt a FILE secret to disk",
		Long: "Fetch a FILE secret and decrypt it locally. The plaintext is written to a temporary " +
			"file in the destination directory first, then moved into place after decryption succeeds.",
		RunE: secretDownloadRun,
	}
	cmd.Flags().StringP(outputFileFlag, "o", "", "Path for the decrypted file. Defaults to the secret filename when available.")
	return cmd
}

func secretDownloadRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}
	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	identity, err := age.GenerateHybridIdentity()
	if err != nil {
		return fmt.Errorf("failed to generate Age identity: %w", err)
	}

	secret, err := resolveSecret(ctx, c, args[0])
	if err != nil {
		return err
	}
	if isExternalSecret(secret) {
		return fmt.Errorf("external secret content must be opened with the external share URL; cone supports metadata for external secrets but not the opener reveal flow")
	}
	vaultID, err := secretVaultID(secret)
	if err != nil {
		return err
	}
	content, err := c.GetSecretContent(ctx, vaultID, identity.Recipient().String())
	if err != nil {
		return err
	}
	if content == nil {
		return fmt.Errorf("secret content response was empty")
	}
	downloadURL := client.StringFromPtr(content.DownloadURL)
	if downloadURL == "" {
		return fmt.Errorf("secret is not a FILE secret; use cone secret view")
	}

	outputPath, err := cmd.Flags().GetString(outputFileFlag)
	if err != nil {
		return err
	}
	outputPath = strings.TrimSpace(outputPath)
	if outputPath == "" {
		outputPath = filepath.Base(client.StringFromPtr(content.Filename))
	}
	if outputPath == "" || outputPath == "." || outputPath == string(filepath.Separator) {
		return fmt.Errorf("--%s is required when the secret has no filename", outputFileFlag)
	}

	encryptedBody, err := c.DownloadSecretFile(ctx, downloadURL)
	if err != nil {
		return err
	}
	defer func() { _ = encryptedBody.Close() }()

	plaintext, err := age.Decrypt(encryptedBody, identity)
	if err != nil {
		return fmt.Errorf("failed to decrypt file content: %w", err)
	}
	if err := writeDecryptedFileAtomically(outputPath, plaintext); err != nil {
		return err
	}
	_, err = fmt.Fprintf(cmd.OutOrStdout(), "Downloaded %s\n", outputPath)
	return err
}

func secretRevokeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "revoke <vault-id|share-code|share-url>",
		Short: "Revoke a secret you created",
		RunE:  secretRevokeRun,
	}
}

func secretRevokeRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}
	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}
	vaultID, err := resolveSecretVaultID(ctx, c, args[0])
	if err != nil {
		return err
	}
	secret, err := c.RevokeSecret(ctx, vaultID)
	if err != nil {
		return err
	}
	resp := Secret(*secret)
	return output.NewManager(ctx, v).Output(ctx, &resp, output.WithTransposeTable())
}

func secretAuditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audit <vault-id|share-code|share-url>",
		Short: "List audit events for a secret you created",
		RunE:  secretAuditRun,
	}
	cmd.Flags().Int(pageSizeFlag, 100, "Page size for API requests.")
	return cmd
}

func secretAuditRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}
	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}
	pageSize := v.GetInt(pageSizeFlag)
	if pageSize <= 0 || pageSize > 1000 {
		return fmt.Errorf("--%s must be between 1 and 1000", pageSizeFlag)
	}
	vaultID, err := resolveSecretVaultID(ctx, c, args[0])
	if err != nil {
		return err
	}
	events, err := c.SearchSecretAuditEvents(ctx, vaultID, pageSize)
	if err != nil {
		return err
	}
	resp := SecretAuditEvents{Events: events}
	return output.NewManager(ctx, v).Output(ctx, &resp)
}

// decryptFromAgeIdentity base64-decodes the Age ciphertext and decrypts it with the given identity.
func decryptFromAgeIdentity(identity age.Identity, encrypted string) (string, error) {
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

func writeDecryptedFileAtomically(outputPath string, plaintext io.Reader) error {
	outputPath = filepath.Clean(outputPath)
	if outputPath == "." {
		return fmt.Errorf("--%s must name a file", outputFileFlag)
	}
	if _, err := os.Stat(outputPath); err == nil {
		return fmt.Errorf("%s already exists", outputPath)
	} else if !os.IsNotExist(err) {
		return err
	}

	dir := filepath.Dir(outputPath)
	base := filepath.Base(outputPath)
	tmp, err := os.CreateTemp(dir, "."+base+".tmp-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	ok := false
	defer func() {
		if !ok {
			_ = os.Remove(tmpPath)
		}
	}()

	if _, err := io.Copy(tmp, plaintext); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Link(tmpPath, outputPath); err != nil {
		return err
	}
	ok = true
	return os.Remove(tmpPath)
}

func resolveSecretVaultID(ctx context.Context, c client.C1Client, ref string) (string, error) {
	secret, err := resolveSecret(ctx, c, ref)
	if err != nil {
		return "", err
	}
	return secretVaultID(secret)
}

func secretVaultID(secret *shared.PaperSecret) (string, error) {
	vaultID := client.StringFromPtr(secret.VaultID)
	if vaultID == "" {
		return "", fmt.Errorf("secret has no vault id")
	}
	return vaultID, nil
}

func isExternalSecret(secret *shared.PaperSecret) bool {
	return secret != nil &&
		secret.SharingMode != nil &&
		*secret.SharingMode == shared.SharingModePaperVaultSharingModeExternal
}

func resolveSecret(ctx context.Context, c client.C1Client, ref string) (*shared.PaperSecret, error) {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return nil, fmt.Errorf("secret identifier must not be empty")
	}
	if vaultID := vaultIDFromSecretURL(ref); vaultID != "" {
		return c.GetSecret(ctx, vaultID)
	}
	if shareCode := shareCodeFromSecretRef(ref); shareCode != "" {
		return c.GetSecretByShareCode(ctx, shareCode)
	}
	return c.GetSecret(ctx, ref)
}

func vaultIDFromSecretURL(ref string) string {
	u, err := url.Parse(ref)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	return strings.TrimSpace(u.Query().Get("vaultId"))
}

func shareCodeFromSecretRef(ref string) string {
	if shareCodeLooksValid(ref) {
		return strings.ToUpper(ref)
	}
	u, err := url.Parse(ref)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	segments := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(segments) == 0 {
		return ""
	}
	last := segments[len(segments)-1]
	if shareCodeLooksValid(last) {
		return strings.ToUpper(last)
	}
	return ""
}

func shareCodeLooksValid(value string) bool {
	parts := strings.Split(strings.TrimSpace(value), "-")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		if len(part) != 4 {
			return false
		}
		for _, r := range part {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
				continue
			}
			return false
		}
	}
	return true
}

func secretListStatuses(input string) ([]shared.PaperSecretServiceSearchMySecretsRequestStatuses, error) {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "", "active":
		return []shared.PaperSecretServiceSearchMySecretsRequestStatuses{
			shared.PaperSecretServiceSearchMySecretsRequestStatusesSecretStatusActive,
		}, nil
	case "all":
		return nil, nil
	case "expired":
		return []shared.PaperSecretServiceSearchMySecretsRequestStatuses{
			shared.PaperSecretServiceSearchMySecretsRequestStatusesSecretStatusExpired,
		}, nil
	case "burned":
		return []shared.PaperSecretServiceSearchMySecretsRequestStatuses{
			shared.PaperSecretServiceSearchMySecretsRequestStatusesSecretStatusBurned,
		}, nil
	case "revoked":
		return []shared.PaperSecretServiceSearchMySecretsRequestStatuses{
			shared.PaperSecretServiceSearchMySecretsRequestStatusesSecretStatusRevoked,
		}, nil
	case "data-deleted":
		return []shared.PaperSecretServiceSearchMySecretsRequestStatuses{
			shared.PaperSecretServiceSearchMySecretsRequestStatusesSecretStatusDataDeleted,
		}, nil
	default:
		return nil, fmt.Errorf("--%s must be active, expired, burned, revoked, data-deleted, or all", secretStatusFlag)
	}
}

func secretListType(input string) (*shared.PaperSecretServiceSearchMySecretsRequestSecretType, error) {
	var secretType shared.PaperSecretServiceSearchMySecretsRequestSecretType
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "", "all":
		return nil, nil
	case "text":
		secretType = shared.PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeText
	case "file":
		secretType = shared.PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeFile
	default:
		return nil, fmt.Errorf("--%s must be text, file, or all", secretTypeFlag)
	}
	return &secretType, nil
}

func secretListSharingMode(input string) (*shared.PaperSecretServiceSearchMySecretsRequestSharingMode, error) {
	var sharingMode shared.PaperSecretServiceSearchMySecretsRequestSharingMode
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "", "all":
		return nil, nil
	case "internal", "team":
		sharingMode = shared.PaperSecretServiceSearchMySecretsRequestSharingModePaperVaultSharingModeInternal
	case "external":
		sharingMode = shared.PaperSecretServiceSearchMySecretsRequestSharingModePaperVaultSharingModeExternal
	default:
		return nil, fmt.Errorf("--%s must be internal, external, or all", secretSharingFlag)
	}
	return &sharingMode, nil
}

func createInputFormat(name string) *shared.PaperSecretServiceCreateInternalRequestInputFormat {
	var format shared.PaperSecretServiceCreateInternalRequestInputFormat
	switch name {
	case formatJSON:
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatJSON
	case formatYAML:
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatYaml
	case formatKeyValue:
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

func createExternalInputFormat(name string) *shared.PaperSecretServiceCreateExternalRequestInputFormat {
	var format shared.PaperSecretServiceCreateExternalRequestInputFormat
	switch name {
	case formatJSON:
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatJSON
	case formatYAML:
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatYaml
	case formatKeyValue:
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

func setContentInputFormat(name string) *shared.PaperSecretServiceSetTextContentRequestInputFormat {
	var format shared.PaperSecretServiceSetTextContentRequestInputFormat
	switch name {
	case formatJSON:
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatJSON
	case formatYAML:
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatYaml
	case formatKeyValue:
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatKeyValue
	default:
		format = shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatPlaintext
	}
	return &format
}

type Secrets []shared.PaperSecret

func (r *Secrets) Header() []string {
	return (&Secret{}).WideHeader()
}

func (r *Secrets) Rows() [][]string {
	rows := make([][]string, 0, len(*r))
	for _, secret := range *r {
		row := Secret(secret)
		rows = append(rows, row.WideRows()[0])
	}
	return rows
}

type SecretAuditEvents struct {
	Events []map[string]any `json:"events"`
}

func (r *SecretAuditEvents) Header() []string {
	return []string{"Event"}
}

func (r *SecretAuditEvents) Rows() [][]string {
	rows := make([][]string, 0, len(r.Events))
	for _, event := range r.Events {
		data, err := json.Marshal(event)
		if err != nil {
			rows = append(rows, []string{fmt.Sprintf("%v", event)})
			continue
		}
		rows = append(rows, []string{string(data)})
	}
	return rows
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
