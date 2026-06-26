package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"filippo.io/age"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func TestEncryptDecryptRoundtrip(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
	}{
		{name: "simple", plaintext: "hello world"},
		{name: "empty", plaintext: ""},
		{name: "multiline", plaintext: "line1\nline2\nline3"},
		{name: "json", plaintext: `{"key":"value","n":1}`},
		{name: "unicode", plaintext: "héllo 🌍 wörld"},
		{name: "large", plaintext: strings.Repeat("a", 50000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity, err := age.GenerateX25519Identity()
			if err != nil {
				t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
			}

			encrypted, err := encryptToAgeRecipient(identity.Recipient().String(), []byte(tt.plaintext))
			if err != nil {
				t.Fatalf("encryptToAgeRecipient() unexpected error: %v", err)
			}

			// Encrypted output must be valid base64 of the raw Age format.
			raw, err := base64.StdEncoding.DecodeString(encrypted)
			if err != nil {
				t.Fatalf("encrypted content is not valid base64: %v", err)
			}
			if !strings.HasPrefix(string(raw), "age-encryption.org/v1") {
				t.Errorf("decoded content is not Age format, got prefix %q", string(raw[:min(len(raw), 22)]))
			}

			got, err := decryptFromAgeIdentity(identity, encrypted)
			if err != nil {
				t.Fatalf("decryptFromAgeIdentity() unexpected error: %v", err)
			}
			if got != tt.plaintext {
				t.Errorf("roundtrip = %q, want %q", got, tt.plaintext)
			}
		})
	}
}

func TestEncryptToAgeRecipientInvalid(t *testing.T) {
	if _, err := encryptToAgeRecipient("not-a-valid-age-recipient", []byte("data")); err == nil {
		t.Error("encryptToAgeRecipient() with invalid recipient expected error, got nil")
	}
}

func TestDecryptFromAgeIdentityErrors(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}

	tests := []struct {
		name      string
		encrypted string
	}{
		{name: "invalid base64", encrypted: "not valid base64!!!"},
		{name: "valid base64 but not age", encrypted: base64.StdEncoding.EncodeToString([]byte("plain bytes"))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decryptFromAgeIdentity(identity, tt.encrypted); err == nil {
				t.Errorf("decryptFromAgeIdentity(%q) expected error, got nil", tt.name)
			}
		})
	}
}

func TestDecryptWithWrongIdentity(t *testing.T) {
	sender, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}
	other, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}

	encrypted, err := encryptToAgeRecipient(sender.Recipient().String(), []byte("secret"))
	if err != nil {
		t.Fatalf("encryptToAgeRecipient() unexpected error: %v", err)
	}

	if _, err := decryptFromAgeIdentity(other, encrypted); err == nil {
		t.Error("decryptFromAgeIdentity() with wrong identity expected error, got nil")
	}
}

func TestCreateInputFormat(t *testing.T) {
	tests := []struct {
		name string
		want shared.PaperSecretServiceCreateInternalRequestInputFormat
	}{
		{name: "json", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatJSON},
		{name: "yaml", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatYaml},
		{name: "key-value", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatKeyValue},
		{name: "plaintext", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext},
		{name: "unknown", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext},
		{name: "", want: shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createInputFormat(tt.name)
			if got == nil {
				t.Fatal("createInputFormat() returned nil")
			}
			if *got != tt.want {
				t.Errorf("createInputFormat(%q) = %q, want %q", tt.name, *got, tt.want)
			}
		})
	}
}

func TestValidateSecretCreateInput(t *testing.T) {
	tests := []struct {
		name    string
		input   secretCreateInput
		wantErr string
	}{
		{name: "valid internal text", input: secretCreateInput{userIDs: []string{"u1"}, hasInlineContent: true, expiresIn: "3600s"}},
		{name: "valid external file", input: secretCreateInput{emails: []string{"a@b.com"}, filePath: "/tmp/f", expiresIn: "1h"}},
		{name: "valid content file", input: secretCreateInput{userRefs: []string{"alice@example.com"}, contentFilePath: "/tmp/content", expiresIn: "3600s"}},
		{name: "no recipient", input: secretCreateInput{hasInlineContent: true, expiresIn: "3600s"}, wantErr: "one of"},
		{name: "both recipients", input: secretCreateInput{userIDs: []string{"u1"}, emails: []string{"a@b.com"}, expiresIn: "3600s", hasInlineContent: true}, wantErr: "mutually exclusive"},
		{name: "invalid email", input: secretCreateInput{emails: []string{"not an email"}, filePath: "/tmp/f", expiresIn: "3600s"}, wantErr: "invalid external email"},
		{name: "missing content source", input: secretCreateInput{userIDs: []string{"u1"}, expiresIn: "3600s"}, wantErr: "one of"},
		{name: "missing expiry", input: secretCreateInput{userIDs: []string{"u1"}, hasInlineContent: true}, wantErr: "expires-in"},
		{name: "content and file", input: secretCreateInput{userIDs: []string{"u1"}, hasInlineContent: true, filePath: "/tmp/f", expiresIn: "3600s"}, wantErr: "mutually exclusive"},
		{name: "max views too high", input: secretCreateInput{userIDs: []string{"u1"}, hasInlineContent: true, expiresIn: "3600s", maxViews: 1001}, wantErr: "between 0 and 1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateSecretCreateInput(tt.input)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("validateSecretCreateInput() unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("validateSecretCreateInput() expected error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("validateSecretCreateInput() error = %q, want contains %q", err.Error(), tt.wantErr)
			}
		})
	}
}

// fakeSecretCreator captures whichever create request createSecret builds so tests can
// assert the request fields without a live API.
type fakeSecretCreator struct {
	internalReq *shared.PaperSecretServiceCreateInternalRequest
	externalReq *shared.PaperSecretServiceCreateExternalRequest
}

func (f *fakeSecretCreator) CreateInternalSecret(_ context.Context, req *shared.PaperSecretServiceCreateInternalRequest) (*shared.PaperSecretServiceCreateResponse, error) {
	f.internalReq = req
	return &shared.PaperSecretServiceCreateResponse{}, nil
}

func (f *fakeSecretCreator) CreateExternalSecret(_ context.Context, req *shared.PaperSecretServiceCreateExternalRequest) (*shared.PaperSecretServiceCreateResponse, error) {
	f.externalReq = req
	return &shared.PaperSecretServiceCreateResponse{}, nil
}

func TestCreateSecretInternalText(t *testing.T) {
	maxViews := int64(5)
	f := &fakeSecretCreator{}
	if _, err := createSecret(context.Background(), f, createSecretParams{
		userIDs:     []string{"u1", "u2"},
		expiresIn:   "3600s",
		displayName: "label",
		maxViews:    &maxViews,
		inputFormat: "json",
	}); err != nil {
		t.Fatalf("createSecret() unexpected error: %v", err)
	}

	if f.externalReq != nil {
		t.Fatal("expected internal request, external was called")
	}
	r := f.internalReq
	if r == nil {
		t.Fatal("internal request was not built")
	}
	if r.SecretType == nil || *r.SecretType != shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeText {
		t.Errorf("SecretType = %v, want Text", r.SecretType)
	}
	if r.InputFormat == nil || *r.InputFormat != shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatJSON {
		t.Errorf("InputFormat = %v, want JSON", r.InputFormat)
	}
	if len(r.AllowedUserIds) != 2 {
		t.Errorf("AllowedUserIds = %v, want 2 ids", r.AllowedUserIds)
	}
	if r.DisplayName == nil || *r.DisplayName != "label" {
		t.Errorf("DisplayName = %v, want label", r.DisplayName)
	}
	if r.MaxViews == nil || *r.MaxViews != 5 {
		t.Errorf("MaxViews = %v, want 5", r.MaxViews)
	}
	if r.ContentType != nil || r.Filename != nil || r.FileSize != nil {
		t.Errorf("file metadata should be nil for TEXT, got contentType=%v filename=%v fileSize=%v", r.ContentType, r.Filename, r.FileSize)
	}
}

func TestCreateSecretInternalFile(t *testing.T) {
	f := &fakeSecretCreator{}
	if _, err := createSecret(context.Background(), f, createSecretParams{
		userIDs:     []string{"u1"},
		expiresIn:   "3600s",
		inputFormat: "json", // must be ignored for files
		isFile:      true,
		filename:    "data.bin",
		contentType: "application/octet-stream",
		fileSize:    1234,
	}); err != nil {
		t.Fatalf("createSecret() unexpected error: %v", err)
	}

	r := f.internalReq
	if r == nil {
		t.Fatal("internal request was not built")
	}
	if r.SecretType == nil || *r.SecretType != shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeFile {
		t.Errorf("SecretType = %v, want File", r.SecretType)
	}
	if r.InputFormat != nil {
		t.Errorf("InputFormat should be nil for FILE, got %v", *r.InputFormat)
	}
	if r.ContentType == nil || *r.ContentType != "application/octet-stream" {
		t.Errorf("ContentType = %v, want application/octet-stream", r.ContentType)
	}
	if r.Filename == nil || *r.Filename != "data.bin" {
		t.Errorf("Filename = %v, want data.bin", r.Filename)
	}
	if r.FileSize == nil || *r.FileSize != 1234 {
		t.Errorf("FileSize = %v, want 1234", r.FileSize)
	}
	if r.DisplayName != nil {
		t.Errorf("DisplayName should be nil when unset, got %v", *r.DisplayName)
	}
}

func TestCreateSecretExternalText(t *testing.T) {
	f := &fakeSecretCreator{}
	if _, err := createSecret(context.Background(), f, createSecretParams{
		emails:      []string{"a@b.com"},
		expiresIn:   "1h",
		inputFormat: "yaml",
	}); err != nil {
		t.Fatalf("createSecret() unexpected error: %v", err)
	}

	if f.internalReq != nil {
		t.Fatal("expected external request, internal was called")
	}
	r := f.externalReq
	if r == nil {
		t.Fatal("external request was not built")
	}
	if r.SecretType == nil || *r.SecretType != shared.PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeText {
		t.Errorf("SecretType = %v, want Text", r.SecretType)
	}
	if r.InputFormat == nil || *r.InputFormat != shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatYaml {
		t.Errorf("InputFormat = %v, want Yaml", r.InputFormat)
	}
	if len(r.AllowedEmails) != 1 || r.AllowedEmails[0] != "a@b.com" {
		t.Errorf("AllowedEmails = %v, want [a@b.com]", r.AllowedEmails)
	}
}

func TestCreateSecretExternalFile(t *testing.T) {
	f := &fakeSecretCreator{}
	if _, err := createSecret(context.Background(), f, createSecretParams{
		emails:      []string{"a@b.com"},
		expiresIn:   "1h",
		isFile:      true,
		filename:    "report.pdf",
		contentType: "application/pdf",
		fileSize:    99,
	}); err != nil {
		t.Fatalf("createSecret() unexpected error: %v", err)
	}

	r := f.externalReq
	if r == nil {
		t.Fatal("external request was not built")
	}
	if r.SecretType == nil || *r.SecretType != shared.PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeFile {
		t.Errorf("SecretType = %v, want File", r.SecretType)
	}
	if r.InputFormat != nil {
		t.Errorf("InputFormat should be nil for FILE, got %v", *r.InputFormat)
	}
	if r.ContentType == nil || *r.ContentType != "application/pdf" {
		t.Errorf("ContentType = %v, want application/pdf", r.ContentType)
	}
	if r.Filename == nil || *r.Filename != "report.pdf" {
		t.Errorf("Filename = %v, want report.pdf", r.Filename)
	}
	if r.FileSize == nil || *r.FileSize != 99 {
		t.Errorf("FileSize = %v, want 99", r.FileSize)
	}
}

func TestCreateExternalInputFormat(t *testing.T) {
	tests := []struct {
		name string
		want shared.PaperSecretServiceCreateExternalRequestInputFormat
	}{
		{name: "json", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatJSON},
		{name: "yaml", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatYaml},
		{name: "key-value", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatKeyValue},
		{name: "plaintext", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext},
		{name: "unknown", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext},
		{name: "", want: shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createExternalInputFormat(tt.name)
			if got == nil {
				t.Fatal("createExternalInputFormat() returned nil")
			}
			if *got != tt.want {
				t.Errorf("createExternalInputFormat(%q) = %q, want %q", tt.name, *got, tt.want)
			}
		})
	}
}

// TestEncryptBytesToAgeRecipient verifies the FILE path produces raw (non-base64) Age bytes
// that decrypt back to the original, since file content is PUT verbatim rather than base64-encoded.
func TestEncryptBytesToAgeRecipient(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}

	plaintext := []byte("binary\x00file\xff content")
	raw, err := encryptBytesToAgeRecipient(identity.Recipient().String(), plaintext)
	if err != nil {
		t.Fatalf("encryptBytesToAgeRecipient() unexpected error: %v", err)
	}
	if !strings.HasPrefix(string(raw), "age-encryption.org/v1") {
		t.Errorf("raw output is not Age format, got prefix %q", string(raw[:min(len(raw), 22)]))
	}

	// Decrypt via the base64-wrapping helper to confirm the bytes roundtrip.
	got, err := decryptFromAgeIdentity(identity, base64.StdEncoding.EncodeToString(raw))
	if err != nil {
		t.Fatalf("decryptFromAgeIdentity() unexpected error: %v", err)
	}
	if got != string(plaintext) {
		t.Errorf("roundtrip = %q, want %q", got, plaintext)
	}
}

func TestEncryptBytesToAgeRecipientInvalid(t *testing.T) {
	if _, err := encryptBytesToAgeRecipient("not-a-valid-age-recipient", []byte("data")); err == nil {
		t.Error("encryptBytesToAgeRecipient() with invalid recipient expected error, got nil")
	}
}

func TestSetContentInputFormat(t *testing.T) {
	tests := []struct {
		name string
		want shared.PaperSecretServiceSetTextContentRequestInputFormat
	}{
		{name: "json", want: shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatJSON},
		{name: "yaml", want: shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatYaml},
		{name: "key-value", want: shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatKeyValue},
		{name: "plaintext", want: shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatPlaintext},
		{name: "unknown", want: shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatPlaintext},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setContentInputFormat(tt.name)
			if got == nil {
				t.Fatal("setContentInputFormat() returned nil")
			}
			if *got != tt.want {
				t.Errorf("setContentInputFormat(%q) = %q, want %q", tt.name, *got, tt.want)
			}
		})
	}
}

func TestNormalizeSecretDuration(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr string
	}{
		{name: "one week", input: "1w", want: "604800s"},
		{name: "seconds", input: "3600s", want: "3600s"},
		{name: "too short", input: "59m", wantErr: "between 1h and 30d"},
		{name: "too long", input: "31d", wantErr: "between 1h and 30d"},
		{name: "empty", wantErr: "required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeSecretDuration(tt.input)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("normalizeSecretDuration() unexpected error: %v", err)
				}
				if got != tt.want {
					t.Fatalf("normalizeSecretDuration() = %q, want %q", got, tt.want)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("normalizeSecretDuration() error = %v, want contains %q", err, tt.wantErr)
			}
		})
	}
}

func TestNormalizeSecretInputFormat(t *testing.T) {
	tests := map[string]string{
		"":          "plaintext",
		"text":      "plaintext",
		"json":      "json",
		"yml":       "yaml",
		"env":       "key-value",
		"key_value": "key-value",
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			got, err := normalizeSecretInputFormat(input)
			if err != nil {
				t.Fatalf("normalizeSecretInputFormat() unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("normalizeSecretInputFormat() = %q, want %q", got, want)
			}
		})
	}
	if _, err := normalizeSecretInputFormat("toml"); err == nil {
		t.Fatal("normalizeSecretInputFormat() expected error")
	}
}

func TestValidateSecretTextContent(t *testing.T) {
	if err := validateSecretTextContent(`{"ok":true}`, "json"); err != nil {
		t.Fatalf("validateSecretTextContent() valid json unexpected error: %v", err)
	}
	if err := validateSecretTextContent(`{"ok":`, "json"); err == nil {
		t.Fatal("validateSecretTextContent() invalid json expected error")
	}
	if err := validateSecretTextContent("", "plaintext"); err == nil {
		t.Fatal("validateSecretTextContent() empty content expected error")
	}
	if err := validateSecretTextContent(strings.Repeat("a", maxTextSecretPlaintextBytes+1), "plaintext"); err == nil {
		t.Fatal("validateSecretTextContent() oversized content expected error")
	}
}

func TestShareCodeFromSecretRef(t *testing.T) {
	tests := map[string]string{
		"ABCD-EFGH-IJKL": "ABCD-EFGH-IJKL",
		"abcd-efgh-ijkl": "ABCD-EFGH-IJKL",
		"https://tenant.example.com/secrets/view/abcd-efgh-ijkl":               "ABCD-EFGH-IJKL",
		"https://tenant.example.com/ext/secrets/share/abcd-efgh-ijkl?x=ignore": "ABCD-EFGH-IJKL",
		"not-a-share-code": "",
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			if got := shareCodeFromSecretRef(input); got != want {
				t.Fatalf("shareCodeFromSecretRef() = %q, want %q", got, want)
			}
		})
	}
}

func TestWriteDecryptedFileAtomically(t *testing.T) {
	dir := t.TempDir()
	dst := filepath.Join(dir, "secret.txt")
	if err := writeDecryptedFileAtomically(dst, bytes.NewReader([]byte("secret"))); err != nil {
		t.Fatalf("writeDecryptedFileAtomically() unexpected error: %v", err)
	}
	//nolint:gosec // test reads a file path created under t.TempDir.
	got, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("ReadFile() unexpected error: %v", err)
	}
	if string(got) != "secret" {
		t.Fatalf("written content = %q, want secret", got)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir() unexpected error: %v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "secret.txt" {
		t.Fatalf("temp file was not cleaned up, entries=%v", entries)
	}
	if err := writeDecryptedFileAtomically(dst, bytes.NewReader([]byte("again"))); err == nil {
		t.Fatal("writeDecryptedFileAtomically() expected existing destination error")
	}
}

func TestEncryptFileToTemp(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}
	src := filepath.Join(t.TempDir(), "source.bin")
	want := []byte("file secret")
	if err := os.WriteFile(src, want, 0o600); err != nil {
		t.Fatalf("WriteFile() unexpected error: %v", err)
	}
	tmp, err := encryptFileToTemp(src, identity.Recipient().String())
	if err != nil {
		t.Fatalf("encryptFileToTemp() unexpected error: %v", err)
	}
	defer func() { _ = tmp.Close() }()
	defer func() { _ = os.Remove(tmp.Name()) }()

	r, err := age.Decrypt(tmp, identity)
	if err != nil {
		t.Fatalf("Decrypt() unexpected error: %v", err)
	}
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll() unexpected error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("decrypt = %q, want %q", got, want)
	}
}
