package client

import (
	"context"

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
