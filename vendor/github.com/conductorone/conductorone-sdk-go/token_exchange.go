package conductoronesdkgo

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/square/go-jose.v2/json"

	"github.com/conductorone/conductorone-sdk-go/uhttp"
)

const (
	grantTypeTokenExchange = "urn:ietf:params:oauth:grant-type:token-exchange"
	subjectTokenTypeJWT    = "urn:ietf:params:oauth:token-type:jwt"
)

// tokenExchangeSource implements oauth2.TokenSource by exchanging an external
// OIDC JWT for a ConductorOne access token via RFC 8693 token exchange.
type tokenExchangeSource struct {
	oidcToken  string
	clientID   string
	tokenHost  string
	httpClient *http.Client
}

func (t *tokenExchangeSource) Token() (*oauth2.Token, error) {
	body := url.Values{
		"grant_type":         []string{grantTypeTokenExchange},
		"subject_token":      []string{t.oidcToken},
		"subject_token_type": []string{subjectTokenTypeJWT},
		"client_id":          []string{t.clientID},
	}

	tokenURL := url.URL{
		Scheme: "https",
		Host:   t.tokenHost,
		Path:   "auth/v1/token",
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, tokenURL.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Unauthenticated, "token exchange failed: %s", resp.Status)
	}

	c1t := &c1Token{}
	err = json.NewDecoder(resp.Body).Decode(c1t)
	if err != nil {
		return nil, err
	}

	if c1t.AccessToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "token exchange failed: empty access token")
	}

	return &oauth2.Token{
		AccessToken: c1t.AccessToken,
		TokenType:   c1t.TokenType,
		Expiry:      time.Now().Add(time.Duration(c1t.Expiry) * time.Second),
	}, nil
}

// NewTokenExchangeSource creates an oauth2.TokenSource that exchanges an external
// OIDC token for a ConductorOne access token via RFC 8693 token exchange.
//
// Parameters:
//   - oidcToken: the raw JWT from the external OIDC provider (e.g. GitHub Actions, HCP Terraform)
//   - clientID: the workload federation trust client ID (e.g. "clever-fox@acme.conductorone.com/wfe")
//   - serverURL: the ConductorOne server URL (e.g. "https://acme.conductor.one")
func NewTokenExchangeSource(ctx context.Context, oidcToken, clientID, serverURL string) (oauth2.TokenSource, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, ctxzap.Extract(ctx)), uhttp.WithUserAgent("conductorone-sdk-go-wfe"))
	if err != nil {
		return nil, err
	}

	tokenHost := strings.TrimPrefix(serverURL, "https://")

	return oauth2.ReuseTokenSource(nil, &tokenExchangeSource{
		oidcToken:  oidcToken,
		clientID:   clientID,
		tokenHost:  tokenHost,
		httpClient: httpClient,
	}), nil
}
