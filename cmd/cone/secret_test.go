package main

import (
	"context"
	"encoding/base64"
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
		name      string
		userIDs   []string
		emails    []string
		content   string
		filePath  string
		expiresIn string
		wantErr   string
	}{
		{name: "valid internal text", userIDs: []string{"u1"}, content: "hi", expiresIn: "3600s"},
		{name: "valid external file", emails: []string{"a@b.com"}, filePath: "/tmp/f", expiresIn: "1h"},
		{name: "valid empty content", userIDs: []string{"u1"}, expiresIn: "3600s"},
		{name: "no recipient", expiresIn: "3600s", wantErr: "one of"},
		{name: "both recipients", userIDs: []string{"u1"}, emails: []string{"a@b.com"}, expiresIn: "3600s", wantErr: "mutually exclusive"},
		{name: "missing expiry", userIDs: []string{"u1"}, content: "hi", wantErr: "expires-in"},
		{name: "content and file", userIDs: []string{"u1"}, content: "hi", filePath: "/tmp/f", expiresIn: "3600s", wantErr: "mutually exclusive"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateSecretCreateInput(tt.userIDs, tt.emails, tt.content, tt.filePath, tt.expiresIn)
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

func TestNormalizeExpiresIn(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr string
	}{
		{name: "seconds at min", in: "3600s", want: "3600s"},
		{name: "hours normalized to seconds", in: "1h", want: "3600s"},
		{name: "compound", in: "1h30m", want: "5400s"},
		{name: "max", in: "720h", want: "2592000s"},
		{name: "below min", in: "30m", wantErr: "must be between"},
		{name: "above max", in: "721h", wantErr: "must be between"},
		{name: "unparseable", in: "soon", wantErr: "invalid"},
		{name: "empty", in: "", wantErr: "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeExpiresIn(tt.in)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("normalizeExpiresIn(%q) expected error containing %q, got nil", tt.in, tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("normalizeExpiresIn(%q) error = %q, want contains %q", tt.in, err.Error(), tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("normalizeExpiresIn(%q) unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("normalizeExpiresIn(%q) = %q, want %q", tt.in, got, tt.want)
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

// TestFileBytesRoundtrip mirrors the FILE upload/download path: raw Age encrypt then raw
// decrypt, with no base64 wrapping (file bytes are PUT/GET verbatim).
func TestFileBytesRoundtrip(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}

	plaintext := []byte("binary\x00file\xff bytes")
	raw, err := encryptBytesToAgeRecipient(identity.Recipient().String(), plaintext)
	if err != nil {
		t.Fatalf("encryptBytesToAgeRecipient() unexpected error: %v", err)
	}

	got, err := decryptBytesFromAgeIdentity(identity, raw)
	if err != nil {
		t.Fatalf("decryptBytesFromAgeIdentity() unexpected error: %v", err)
	}
	if string(got) != string(plaintext) {
		t.Errorf("roundtrip = %q, want %q", got, plaintext)
	}
}

func TestDecryptBytesFromAgeIdentityInvalid(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatalf("GenerateX25519Identity() unexpected error: %v", err)
	}
	if _, err := decryptBytesFromAgeIdentity(identity, []byte("not age ciphertext")); err == nil {
		t.Error("decryptBytesFromAgeIdentity() with non-Age bytes expected error, got nil")
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
