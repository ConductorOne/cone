package conductoronesdkgo

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/conductorone/conductorone-sdk-go/uhttp"
)

const c1TenantDomain = ".conductor.one"
const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

// Environment variable names for ConductorOne authentication.
// These are used across all ConductorOne client tools (sdk-go, cone, terraform-provider).
const (
	EnvAccessToken  = "CONDUCTORONE_ACCESS_TOKEN"
	EnvOIDCToken    = "CONDUCTORONE_OIDC_TOKEN"
	EnvClientID     = "CONDUCTORONE_CLIENT_ID"
	EnvClientSecret = "CONDUCTORONE_CLIENT_SECRET"
	EnvTenantDomain = "CONDUCTORONE_TENANT_DOMAIN"
	EnvServerURL    = "CONDUCTORONE_SERVER_URL"
)

func WithTenant(input string) (SDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	if resp.UseWithTenant() {
		return WithTenantDomain(resp.Tenant()), nil
	}

	if resp.UseWithServer() {
		return WithServerURL(resp.ServerURL()), nil
	}

	return func(api *ConductoroneAPI) {}, nil
}

func WithExtraUserAgent(userAgent string) SDKOption {
	return func(sdk *ConductoroneAPI) {
		sdk.sdkConfiguration.UserAgent = fmt.Sprintf("%s %s", userAgent, sdk.sdkConfiguration.UserAgent)
	}
}

type CustomSDKOption func(*CustomOptions)

func WithTenantCustom(input string) (CustomSDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	return func(sdk *CustomOptions) {
		sdk.ClientConfig = resp
	}, nil
}

func WithLog(logger *zap.Logger) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.logger = logger
	}
}

func WithUserAgent(userAgent string) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.userAgent = userAgent
	}
}
func WithTLSConfig(tlsConfig *tls.Config) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.tlsConfig = tlsConfig
	}
}

type ClientConfig struct {
	// These are mutually exclusive
	serverURL string
	tenant    string
}

func (c *ClientConfig) UseWithServer() bool {
	return c.ServerURL() != ""
}

func (c *ClientConfig) UseWithTenant() bool {
	return c.Tenant() != ""
}

func (c *ClientConfig) SetTenant(tenant string) error {
	if c == nil {
		return errors.New("client config is nil, cannot set tenant")
	}
	if c.UseWithServer() {
		return errors.New("cannot set tenant, tenant and serverURL are mutually exclusive")
	}
	c.tenant = tenant
	return nil
}

func (c *ClientConfig) SetServerURL(serverURL string) error {
	if c == nil {
		return errors.New("client config is nil, cannot set serverURL")
	}
	if c.UseWithTenant() {
		return errors.New("cannot set serverURL, tenant and serverURL are mutually exclusive")
	}
	c.serverURL = serverURL
	return nil
}

func (c *ClientConfig) Tenant() string {
	if c == nil {
		return ""
	}
	return c.tenant
}

// ServerURL returns the server URL.
func (c *ClientConfig) ServerURL() string {
	if c == nil {
		return ""
	}
	return c.serverURL
}

// GetServerURL returns the server URL. If serverURL is empty (""), it constructs the server URL using the tenant. However, if the tenant is also empty, then it will return an empty string.
func (c *ClientConfig) GetServerURL() string {
	if c.UseWithServer() {
		return c.serverURL
	}
	if c.UseWithTenant() {
		u := &url.URL{}
		tenant := strings.ToLower(c.Tenant())
		u.Host = tenant + c1TenantDomain
		u.Scheme = "https"
		return u.String()
	}
	return ""
}

type CustomOptions struct {
	*ClientConfig

	withClient *http.Client
	logger     *zap.Logger
	tlsConfig  *tls.Config

	userAgent string
}

// NewWithCredentials creates a ConductoroneAPI client using explicit client credentials
// (Ed25519 JWT bearer assertion). This function does NOT read environment variables;
// use [NewWithEnv] for environment-based auth.
func NewWithCredentials(ctx context.Context, cred *ClientCredentials, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.GetServerURL() == "" {
		resp, err := ParseClientID(cred.ClientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = resp
	}

	tokenSource, err := NewTokenSource(ctx, cred.ClientID, cred.ClientSecret, options.GetServerURL())
	if err != nil {
		return nil, err
	}

	return newWithTokenSource(ctx, tokenSource, cred.ClientID, options)
}

// NewWithOIDCToken creates a ConductoroneAPI client that exchanges an external OIDC token
// for a ConductorOne access token via RFC 8693 token exchange.
func NewWithOIDCToken(ctx context.Context, oidcToken, clientID string, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}
	return newWithOIDCToken(ctx, oidcToken, clientID, options)
}

// NewWithAccessToken creates a ConductoroneAPI client using a pre-exchanged bearer token.
func NewWithAccessToken(ctx context.Context, accessToken, clientID string, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}
	return newWithTokenSource(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken,
	}), clientID, options)
}

// NewWithEnv creates a ConductoroneAPI client using environment variables.
// Auth priority:
//  1. CONDUCTORONE_ACCESS_TOKEN -- static bearer token
//  2. CONDUCTORONE_OIDC_TOKEN + CONDUCTORONE_CLIENT_ID -- RFC 8693 token exchange
//  3. CONDUCTORONE_CLIENT_ID + CONDUCTORONE_CLIENT_SECRET -- Ed25519 JWT assertion
func NewWithEnv(ctx context.Context, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}

	clientID := os.Getenv(EnvClientID)

	// Priority 1: CONDUCTORONE_ACCESS_TOKEN
	if accessToken := os.Getenv(EnvAccessToken); accessToken != "" {
		return newWithTokenSource(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: accessToken,
		}), clientID, options)
	}

	// Priority 2: CONDUCTORONE_OIDC_TOKEN
	if oidcToken := os.Getenv(EnvOIDCToken); oidcToken != "" {
		if clientID == "" {
			return nil, fmt.Errorf("%s requires %s", EnvOIDCToken, EnvClientID)
		}
		return newWithOIDCToken(ctx, oidcToken, clientID, options)
	}

	// Priority 3: CONDUCTORONE_CLIENT_ID + CONDUCTORONE_CLIENT_SECRET
	clientSecret := os.Getenv(EnvClientSecret)
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("one of %s, %s, or %s+%s must be set", EnvAccessToken, EnvOIDCToken, EnvClientID, EnvClientSecret)
	}

	if options.GetServerURL() == "" {
		resp, err := ParseClientID(clientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = resp
	}

	tokenSource, err := NewTokenSource(ctx, clientID, clientSecret, options.GetServerURL())
	if err != nil {
		return nil, err
	}

	return newWithTokenSource(ctx, tokenSource, clientID, options)
}

func newWithOIDCToken(ctx context.Context, oidcToken, clientID string, options *CustomOptions) (*ConductoroneAPI, error) {
	if options.GetServerURL() == "" {
		resp, err := ParseClientID(clientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = resp
	}

	tokenSource, err := NewTokenExchangeSource(ctx, oidcToken, clientID, options.GetServerURL())
	if err != nil {
		return nil, err
	}

	return newWithTokenSource(ctx, tokenSource, clientID, options)
}

func newWithTokenSource(ctx context.Context, tokenSource oauth2.TokenSource, clientID string, options *CustomOptions) (*ConductoroneAPI, error) {
	if options.GetServerURL() == "" && clientID != "" {
		resp, err := ParseClientID(clientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = resp
	}

	if options.userAgent == "" {
		options.userAgent = "conductorone-sdk-go"
	}

	sdkOpts := []SDKOption{}
	if options.withClient == nil {
		uclient, err := uhttp.NewClient(
			ctx,
			uhttp.WithTokenSource(tokenSource),
			uhttp.WithLogger(options.logger != nil, options.logger),
			uhttp.WithUserAgent(options.userAgent),
			uhttp.WithTLSClientConfig(options.tlsConfig),
		)
		if err != nil {
			return nil, err
		}
		sdkOpts = append(sdkOpts, WithClient(uclient))
	} else {
		sdkOpts = append(sdkOpts, WithClient(options.withClient))
	}

	if options.UseWithServer() {
		sdkOpts = append(sdkOpts, WithServerURL(options.ServerURL()))
	}
	if options.UseWithTenant() {
		sdkOpts = append(sdkOpts, WithTenantDomain(options.Tenant()))
	}

	return New(sdkOpts...), nil
}

func NormalizeTenant(input string) (*ClientConfig, error) {
	input = strings.ToLower(input)

	var err error
	u := &url.URL{}
	if !strings.Contains(input, "://") {
		if !strings.Contains(input, ".") {
			input += c1TenantDomain
		}
		u.Host = input
	} else {
		u, err = url.Parse(input)
		if err != nil {
			return nil, err
		}
	}

	normalize := &ClientConfig{}

	parts := strings.Split(u.Host, ".")
	if len(parts) == 3 && parts[1] == "conductor" && parts[2] == "one" {
		err := normalize.SetTenant(parts[0])
		if err != nil {
			return nil, err
		}
		return normalize, nil
	}

	u.Scheme = "https"
	err = normalize.SetServerURL(u.String())
	if err != nil {
		return nil, err
	}
	return normalize, nil
}

func ParseClientID(input string) (*ClientConfig, error) {
	// split the input into 2 parts by @
	items := strings.SplitN(input, "@", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	// split the right part into 2 parts by /
	items = strings.SplitN(items[1], "/", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	resp, err := NormalizeTenant(items[0])
	if err != nil {
		return nil, err
	}

	if resp.GetServerURL() == "" {
		return nil, ErrInvalidClientID
	}

	return resp, nil
}
