package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

// CreateInternalSecret creates an internal secret and returns the create response.
// The response carries the vault ID (primary identifier) and the age recipient
// public key, which must be used to encrypt content before calling SetSecretTextContent.
func (c *client) CreateInternalSecret(
	ctx context.Context,
	req *shared.PaperSecretServiceCreateInternalRequest,
) (*shared.PaperSecretServiceCreateResponse, error) {
	resp, err := c.sdk.PaperSecret.CreateInternal(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.PaperSecretServiceCreateResponse, nil
}

// CreateExternalSecret creates a secret shared with external email recipients
// (who authenticate via email magic link or Google OAuth) and returns the create
// response. Like CreateInternalSecret, the response carries the vault ID and the
// age recipient public key used to encrypt content before upload.
func (c *client) CreateExternalSecret(
	ctx context.Context,
	req *shared.PaperSecretServiceCreateExternalRequest,
) (*shared.PaperSecretServiceCreateResponse, error) {
	resp, err := c.sdk.PaperSecret.CreateExternal(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.PaperSecretServiceCreateResponse, nil
}

// UploadSecretFile uploads the Age-encrypted bytes of a FILE secret to the capability
// upload URL returned by CreateInternalSecret/CreateExternalSecret. The encrypted bytes
// are sent verbatim as the PUT body (they must begin with the Age header
// "age-encryption.org/v1"). The upload URL is self-authorizing, so a bare HTTP client is
// used to avoid attaching the ConductorOne bearer token to the (foreign) storage host.
func (c *client) UploadSecretFile(ctx context.Context, uploadURL string, encrypted []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, bytes.NewReader(encrypted))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("file upload failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

// DownloadSecretFile fetches the bytes of a FILE secret from the capability download URL
// returned by GetSecretContent. The bytes are Age-encrypted to the reader recipient supplied
// to GetSecretContent, so the caller decrypts them with the matching identity. A bare HTTP
// client is used since the URL is self-authorizing.
func (c *client) DownloadSecretFile(ctx context.Context, downloadURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("file download failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return io.ReadAll(resp.Body)
}

// SetSecretTextContent uploads the encrypted content for a TEXT secret. The
// encryptedContent must already be Age-encrypted to the recipient returned by
// CreateInternalSecret; the SDK base64-encodes the bytes for transport.
func (c *client) SetSecretTextContent(
	ctx context.Context,
	vaultID string,
	encryptedContent string,
	inputFormat *shared.PaperSecretServiceSetTextContentRequestInputFormat,
) error {
	resp, err := c.sdk.PaperSecret.SetTextContent(ctx, operations.C1APISecretsV1PaperSecretServiceSetTextContentRequest{
		VaultID: vaultID,
		PaperSecretServiceSetTextContentRequest: &shared.PaperSecretServiceSetTextContentRequest{
			EncryptedContent: &encryptedContent,
			InputFormat:      inputFormat,
		},
	})
	if err != nil {
		return err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return err
	}
	return nil
}

// GetSecretContent fetches a secret's content. The server re-encrypts the content to
// readerRecipient (an Age "age1..." recipient), so the caller must hold the matching
// Age identity to decrypt the returned bytes. Reading may count against the secret's
// view limit.
func (c *client) GetSecretContent(
	ctx context.Context,
	vaultID string,
	readerRecipient string,
) (*shared.PaperSecretServiceGetContentResponse, error) {
	resp, err := c.sdk.PaperSecret.GetContent(ctx, operations.C1APISecretsV1PaperSecretServiceGetContentRequest{
		VaultID: vaultID,
		PaperSecretServiceGetContentRequest: &shared.PaperSecretServiceGetContentRequest{
			ReaderRecipient: &readerRecipient,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.PaperSecretServiceGetContentResponse, nil
}

// GetSecret returns the metadata for a secret by its vault ID.
func (c *client) GetSecret(ctx context.Context, vaultID string) (*shared.PaperSecret, error) {
	resp, err := c.sdk.PaperSecret.Get(ctx, operations.C1APISecretsV1PaperSecretServiceGetRequest{
		VaultID: vaultID,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.PaperSecretServiceGetResponse.PaperSecret, nil
}
