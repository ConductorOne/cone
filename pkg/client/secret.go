package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/sdkerrors"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

// AgeSuiteMismatchError indicates that PaperSecret did not honor the exact Age
// suite requested by cone. Retrying is safe: no content has been encrypted or
// decrypted when this error is returned.
type AgeSuiteMismatchError struct {
	Operation string
	Returned  shared.PaperSecretServiceCreateResponseAgeSuite
}

func (e *AgeSuiteMismatchError) Error() string {
	returned := string(e.Returned)
	if e.Returned == "" || e.Returned == shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteUnspecified {
		returned = string(shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteUnspecified)
	}
	return fmt.Sprintf(
		"PaperSecret %s returned Age suite %s, required %s; retry the request",
		e.Operation,
		returned,
		shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519,
	)
}

func (e *AgeSuiteMismatchError) Temporary() bool { return true }

func requirePaperSecretAgeSuite(operation string, returned *shared.PaperSecretServiceCreateResponseAgeSuite) error {
	if returned == nil {
		return &AgeSuiteMismatchError{Operation: operation}
	}
	if *returned != shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519 {
		return &AgeSuiteMismatchError{Operation: operation, Returned: *returned}
	}
	return nil
}

func mapPaperSecretCreateError(err error) error {
	var sdkErr *sdkerrors.SDKError
	if errors.As(err, &sdkErr) && sdkErr.StatusCode >= http.StatusBadRequest {
		return &HTTPError{StatusCode: sdkErr.StatusCode, Body: sdkErr.Body}
	}
	return err
}

// CreateInternalSecret creates an internal secret and returns the create response.
// The response carries the vault ID (primary identifier) and the age recipient
// public key, which must be used to encrypt content before calling SetSecretTextContent.
func (c *client) CreateInternalSecret(
	ctx context.Context,
	req *shared.PaperSecretServiceCreateInternalRequest,
) (*shared.PaperSecretServiceCreateResponse, error) {
	request := shared.PaperSecretServiceCreateInternalRequest{}
	if req != nil {
		request = *req
	}
	request.RequiredAgeSuite = shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteMlkem768X25519.ToPointer()

	resp, err := c.sdk.PaperSecret.CreateInternal(ctx, &request)
	if err != nil {
		return nil, mapPaperSecretCreateError(err)
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	createResponse := resp.GetPaperSecretServiceCreateResponse()
	if err := requirePaperSecretAgeSuite("create", createResponse.GetAgeSuite()); err != nil {
		return nil, err
	}
	return createResponse, nil
}

// CreateExternalSecret creates a secret shared with external email recipients
// (who authenticate via email magic link or Google OAuth) and returns the create
// response. Like CreateInternalSecret, the response carries the vault ID and the
// age recipient public key used to encrypt content before upload.
func (c *client) CreateExternalSecret(
	ctx context.Context,
	req *shared.PaperSecretServiceCreateExternalRequest,
) (*shared.PaperSecretServiceCreateResponse, error) {
	request := shared.PaperSecretServiceCreateExternalRequest{}
	if req != nil {
		request = *req
	}
	request.RequiredAgeSuite = shared.RequiredAgeSuiteAgeSuiteMlkem768X25519.ToPointer()

	resp, err := c.sdk.PaperSecret.CreateExternal(ctx, &request)
	if err != nil {
		return nil, mapPaperSecretCreateError(err)
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	createResponse := resp.GetPaperSecretServiceCreateResponse()
	if err := requirePaperSecretAgeSuite("create", createResponse.GetAgeSuite()); err != nil {
		return nil, err
	}
	return createResponse, nil
}

// UploadSecretFile uploads the Age-encrypted bytes of a FILE secret to the capability
// upload URL returned by CreateInternalSecret/CreateExternalSecret. The encrypted bytes
// are sent verbatim as the PUT body (they must begin with the Age header
// "age-encryption.org/v1"). The upload URL is self-authorizing, so a bare HTTP client is
// used to avoid attaching the ConductorOne bearer token to the (foreign) storage host.
func (c *client) UploadSecretFile(ctx context.Context, uploadURL string, encrypted []byte) error {
	return c.UploadSecretFileReader(ctx, uploadURL, bytes.NewReader(encrypted), int64(len(encrypted)))
}

func (c *client) UploadSecretFileReader(ctx context.Context, uploadURL string, encrypted io.Reader, contentLength int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, encrypted)
	if err != nil {
		return err
	}
	req.ContentLength = contentLength
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("file upload failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

func (c *client) DownloadSecretFile(ctx context.Context, downloadURL string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		defer func() { _ = resp.Body.Close() }()
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("file download failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return resp.Body, nil
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
// readerRecipient (an Age "age1pq1..." hybrid recipient), so the caller must hold the matching
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
	if resp.PaperSecretServiceGetResponse == nil || resp.PaperSecretServiceGetResponse.Secret == nil {
		return nil, fmt.Errorf("get secret response was empty")
	}
	return resp.PaperSecretServiceGetResponse.Secret, nil
}

func (c *client) GetSecretByShareCode(ctx context.Context, shareCode string) (*shared.PaperSecret, error) {
	resp, err := c.sdk.PaperSecret.GetByShareCode(ctx, operations.C1APISecretsV1PaperSecretServiceGetByShareCodeRequest{
		ShareCode: shareCode,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	if resp.PaperSecretServiceGetResponse == nil || resp.PaperSecretServiceGetResponse.Secret == nil {
		return nil, fmt.Errorf("get secret by share code response was empty")
	}
	return resp.PaperSecretServiceGetResponse.Secret, nil
}

func (c *client) SearchMySecrets(ctx context.Context, req *shared.PaperSecretServiceSearchMySecretsRequest) ([]shared.PaperSecret, error) {
	if req == nil {
		req = &shared.PaperSecretServiceSearchMySecretsRequest{}
	}
	var out []shared.PaperSecret
	for {
		resp, err := c.sdk.PaperSecret.SearchMySecrets(ctx, req)
		if err != nil {
			return nil, err
		}
		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}
		if resp.PaperSecretServiceSearchResponse != nil {
			out = append(out, resp.PaperSecretServiceSearchResponse.List...)
			token := StringFromPtr(resp.PaperSecretServiceSearchResponse.NextPageToken)
			if token != "" {
				req.PageToken = &token
				continue
			}
		}
		return out, nil
	}
}

func (c *client) RevokeSecret(ctx context.Context, vaultID string) (*shared.PaperSecret, error) {
	resp, err := c.sdk.PaperSecret.Revoke(ctx, operations.C1APISecretsV1PaperSecretServiceRevokeRequest{
		VaultID:                         vaultID,
		PaperSecretServiceRevokeRequest: &shared.PaperSecretServiceRevokeRequest{},
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	if resp.PaperSecretServiceRevokeResponse == nil || resp.PaperSecretServiceRevokeResponse.Secret == nil {
		return nil, fmt.Errorf("revoke secret response was empty")
	}
	return resp.PaperSecretServiceRevokeResponse.Secret, nil
}

func (c *client) SearchSecretAuditEvents(ctx context.Context, vaultID string, pageSize int) ([]map[string]any, error) {
	req := &shared.PaperSecretServiceSearchAuditEventsRequest{
		VaultID:  &vaultID,
		PageSize: &pageSize,
	}
	var out []map[string]any
	for {
		resp, err := c.sdk.PaperSecret.SearchAuditEvents(ctx, req)
		if err != nil {
			return nil, err
		}
		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}
		if resp.PaperSecretServiceSearchAuditEventsResponse != nil {
			out = append(out, resp.PaperSecretServiceSearchAuditEventsResponse.List...)
			token := StringFromPtr(resp.PaperSecretServiceSearchAuditEventsResponse.NextPageToken)
			if token != "" {
				req.PageToken = &token
				continue
			}
		}
		return out, nil
	}
}
