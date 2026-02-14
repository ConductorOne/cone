package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	sdk "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/uhttp"
)

const ConeClientID = "2RGdOS94VDferT9e80mdgntl36K"

// Environment variable names for ConductorOne authentication.
// These match the constants in conductorone-sdk-go and terraform-provider-conductorone.
const (
	EnvAccessToken  = "CONDUCTORONE_ACCESS_TOKEN"
	EnvOIDCToken    = "CONDUCTORONE_OIDC_TOKEN"
	EnvClientID     = "CONDUCTORONE_CLIENT_ID"
	EnvClientSecret = "CONDUCTORONE_CLIENT_SECRET"
	EnvServerURL    = "CONDUCTORONE_SERVER_URL"
)

// normalizeHost extracts the host (with port) from a value that may be a
// full URL ("https://host:port/"), a bare hostname ("host"), or host:port.
func normalizeHost(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}
	if strings.Contains(input, "://") {
		u, err := url.Parse(input)
		if err == nil && u.Host != "" {
			return u.Host
		}
	}
	return strings.TrimRight(input, "/")
}

// ResolveServerHost determines the API server host using a consistent priority:
//  1. --api-endpoint flag (via viper)
//  2. CONDUCTORONE_SERVER_URL env var
//  3. CONE_API_ENDPOINT env var
//  4. Parsed from clientID (e.g. "name@host/suffix" -> "host")
//
// Returns (clientName, host, error). clientName is empty if no clientID provided.
func ResolveServerHost(clientID string, v *viper.Viper) (string, string, error) {
	// Check explicit overrides first
	if h := normalizeHost(v.GetString("api-endpoint")); h != "" {
		clientName, _, err := parseClientIDName(clientID)
		if err != nil && clientID != "" {
			return "", "", err
		}
		return clientName, h, nil
	}
	if h := normalizeHost(os.Getenv(EnvServerURL)); h != "" {
		clientName, _, err := parseClientIDName(clientID)
		if err != nil && clientID != "" {
			return "", "", err
		}
		return clientName, h, nil
	}
	if h := normalizeHost(os.Getenv("CONE_API_ENDPOINT")); h != "" {
		clientName, _, err := parseClientIDName(clientID)
		if err != nil && clientID != "" {
			return "", "", err
		}
		return clientName, h, nil
	}

	// Fall back to parsing host from client ID
	if clientID != "" {
		clientName, host, err := parseClientIDName(clientID)
		if err != nil {
			return "", "", err
		}
		return clientName, host, nil
	}

	return "", "", nil
}

// parseClientIDName splits a client ID into (cutename, host, error).
// Client IDs have the format "cutename@host/suffix".
func parseClientIDName(input string) (string, string, error) {
	if input == "" {
		return "", "", nil
	}
	items := strings.SplitN(input, "@", 2)
	if len(items) != 2 {
		return "", "", ErrInvalidClientID
	}
	clientName := items[0]

	parts := strings.SplitN(items[1], "/", 2)
	if len(parts) != 2 {
		return "", "", ErrInvalidClientID
	}

	return clientName, parts[0], nil
}

type contextKey string

const VersionKey contextKey = "version"

type client struct {
	httpClient *http.Client
	clientName string
	endpoint   string
	baseURL    *url.URL
	sdk        *sdk.ConductoroneAPI
}

func StringFromPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func StringFromIntPtr(s *int64) string {
	if s == nil {
		return ""
	}

	return strconv.Itoa(int(*s))
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(i int) *int {
	return &i
}

type C1Client interface {
	BaseURL() string

	AuthIntrospect(ctx context.Context) (*shared.IntrospectResponse, error)
	GetUser(ctx context.Context, userID string) (*shared.User, error)
	GetEntitlement(ctx context.Context, appID string, entitlementID string) (*shared.AppEntitlement, error)
	SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error)
	GetResource(ctx context.Context, appID string, resourceID string, resourceTypeID string) (*shared.AppResource, error)
	GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*shared.AppResourceType, error)
	GetApp(ctx context.Context, appID string) (*shared.App, error)
	GetTask(ctx context.Context, taskId string) (*shared.TaskServiceGetResponse, error)
	CreateGrantTask(
		ctx context.Context,
		appId string,
		appEntitlementId string,
		identityUserId string,
		appUserId string,
		justification string,
		duration string,
		emergencyAccess bool,
		requestData map[string]any,
	) (*shared.TaskServiceCreateGrantResponse, error)
	CreateRevokeTask(
		ctx context.Context,
		appId string,
		appEntitlementId string,
		identityUserId string,
		justification string,
	) (*shared.TaskServiceCreateRevokeResponse, error)
	GetGrantsForIdentity(ctx context.Context, appID string, appEntitlementID string, identityID string) ([]shared.AppEntitlementUserBinding, error)
	SearchTasks(ctx context.Context, taskFilter shared.TaskSearchRequest) (*shared.TaskSearchResponse, error)
	CommentOnTask(ctx context.Context, taskID string, comment string) (*shared.TaskActionsServiceCommentResponse, error)
	ApproveTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceApproveResponse, error)
	DenyTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceDenyResponse, error)
	EscalateTask(ctx context.Context, taskId string) (*shared.TaskServiceActionResponse, error)
	UpdateTaskRequestData(ctx context.Context, taskID string, requestData map[string]any) (*shared.TaskServiceActionResponse, error)
	ListApps(ctx context.Context) ([]shared.App, error)
	ListAppUsers(ctx context.Context, appID string) ([]shared.AppUser, error)
	ListAppUsersForUser(ctx context.Context, appID string, userID string) ([]shared.AppUser, error)
	ListAppUserCredentials(ctx context.Context, appID string, appUserID string) ([]shared.AppUserCredential, error)
	ListPolicies(ctx context.Context) ([]shared.Policy, error)
	ListEntitlements(ctx context.Context, appId string) ([]shared.AppEntitlement, error)
}

func (c *client) BaseURL() string {
	return c.baseURL.String()
}

func New(
	ctx context.Context,
	clientId string,
	clientSecret string,
	v *viper.Viper,
	cmdName string,
) (C1Client, error) {
	clientName, tokenHost, err := ResolveServerHost(clientId, v)
	if err != nil {
		return nil, err
	}

	tokenSrc, err := NewC1TokenSource(ctx, clientId, clientSecret, tokenHost, v.GetBool("debug"))
	if err != nil {
		return nil, err
	}

	return newClientWithTokenSource(ctx, tokenSrc, clientName, tokenHost, v, cmdName)
}

// NewWithAccessToken creates a client using a pre-exchanged bearer token.
func NewWithAccessToken(
	ctx context.Context,
	accessToken string,
	clientID string,
	v *viper.Viper,
	cmdName string,
) (C1Client, error) {
	clientName, tokenHost, err := ResolveServerHost(clientID, v)
	if err != nil {
		return nil, err
	}
	if tokenHost == "" {
		return nil, fmt.Errorf("%s requires --client-id, %s, or --api-endpoint to determine the server", EnvAccessToken, EnvClientID)
	}

	tokenSrc := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken,
	})

	return newClientWithTokenSource(ctx, tokenSrc, clientName, tokenHost, v, cmdName)
}

// NewWithOIDCToken creates a client that exchanges an OIDC token for a C1 access token.
func NewWithOIDCToken(
	ctx context.Context,
	oidcToken string,
	clientID string,
	v *viper.Viper,
	cmdName string,
) (C1Client, error) {
	if oidcToken == "" {
		return nil, fmt.Errorf("NewWithOIDCToken: oidcToken must be non-empty; set %s or pass --oidc-token", EnvOIDCToken)
	}
	if clientID == "" {
		return nil, fmt.Errorf("NewWithOIDCToken: clientID must be non-empty; set %s or pass --client-id", EnvClientID)
	}

	clientName, tokenHost, err := ResolveServerHost(clientID, v)
	if err != nil {
		return nil, err
	}
	if tokenHost == "" {
		return nil, fmt.Errorf("NewWithOIDCToken: could not determine server host from clientID or --api-endpoint; parseClientID requires a clientID in the form \"name@host/suffix\"")
	}

	tokenSrc, err := NewTokenExchangeSource(ctx, oidcToken, clientID, tokenHost, v.GetBool("debug"))
	if err != nil {
		return nil, err
	}

	return newClientWithTokenSource(ctx, tokenSrc, clientName, tokenHost, v, cmdName)
}

func newClientWithTokenSource(
	ctx context.Context,
	tokenSrc oauth2.TokenSource,
	clientName string,
	tokenHost string,
	v *viper.Viper,
	cmdName string,
) (C1Client, error) {
	uclient, err := uhttp.NewClient(
		ctx,
		uhttp.WithTokenSource(tokenSrc),
		uhttp.WithDebug(v.GetBool("debug")),
		uhttp.WithRequestSource(cmdName),
	)
	if err != nil {
		return nil, err
	}
	c := &client{
		endpoint:   tokenHost,
		clientName: clientName,
		httpClient: uclient,
	}

	apiURL := url.URL{
		Scheme: "https",
		Host:   c.endpoint,
	}
	c.baseURL = &apiURL

	var version = "dev"

	if v := ctx.Value(VersionKey).(string); v != "" {
		version = v
	}

	c.sdk = sdk.New(
		sdk.WithClient(uclient),
		sdk.WithServerURL(apiURL.String()),
		sdk.WithExtraUserAgent(fmt.Sprintf("cone/%s", version)),
	)

	return c, nil
}
