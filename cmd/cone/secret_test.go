package main

import (
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
