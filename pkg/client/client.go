package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"

	sdk "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/uhttp"
)

const ConeClientID = "2RGdOS94VDferT9e80mdgntl36K"

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
		justification string,
		duration string,
		emergencyAccess bool,
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
	ListApps(ctx context.Context) ([]shared.App, error)
	ListAppUsers(ctx context.Context, appID string) ([]shared.AppUser, error)
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
	tokenSrc, clientName, tokenHost, err := NewC1TokenSource(ctx,
		clientId, clientSecret,
		v.GetString("api-endpoint"),
		v.GetBool("debug"),
	)
	if err != nil {
		return nil, err
	}

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
