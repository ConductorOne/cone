package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"

	sdk "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/uhttp"
)

const ConeClientID = "2RGdOS94VDferT9e80mdgntl36K"

type client struct {
	httpClient *http.Client
	clientName string
	tokenHost  string
	baseURL    *url.URL
	sdk        *sdk.ConductoroneAPI
}

func StringFromPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func float64Ptr(i int) *float64 {
	f := float64(i)
	return &f
}

type C1Client interface {
	BaseURL() string

	AuthIntrospect(ctx context.Context) (*shared.IntrospectResponse, error)
	GetUser(ctx context.Context, userID string) (*shared.User, error)
	GetEntitlement(ctx context.Context, appID string, entitlementID string) (*shared.AppEntitlement, error)
	SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error)
	ExpandEntitlements(ctx context.Context, in []*EntitlementWithBindings) (*Expander, error)
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
	SearchTasks(ctx context.Context, taskFilter shared.TaskSearchRequestInput) (*shared.TaskSearchResponse, error)
	CommentOnTask(ctx context.Context, taskID string, comment string) (*shared.TaskActionsServiceCommentResponse, error)
	ApproveTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceApproveResponse, error)
	DenyTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceDenyResponse, error)
	EscalateTask(ctx context.Context, taskId string) (*shared.TaskServiceActionResponse, error)
}

func (c *client) BaseURL() string {
	return c.baseURL.String()
}

func New(
	ctx context.Context,
	clientId string,
	clientSecret string,
	v *viper.Viper,
) (C1Client, error) {
	tokenSrc, clientName, tokenHost, err := NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	uclient, err := uhttp.NewClient(
		ctx,
		uhttp.WithTokenSource(tokenSrc),
		uhttp.WithDebug(v.GetBool("debug")),
	)
	if err != nil {
		return nil, err
	}
	c := &client{
		tokenHost:  tokenHost,
		clientName: clientName,
		httpClient: uclient,
	}

	var apiHostname string
	// If the API host is set in the environment, use that instead of the default
	// HACK(jirwin): Instead of using the generated client's server address, use the hostname from the token.
	if apiHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok && apiHost != "" {
		apiHostname = apiHost
	} else {
		apiHostname = c.tokenHost
	}
	apiURL := url.URL{
		Scheme: "https",
		Host:   apiHostname,
	}
	c.baseURL = &apiURL

	c.sdk = sdk.New(
		sdk.WithClient(uclient),
		sdk.WithServerURL(apiURL.String()),
	)

	return c, nil
}

func handleBadStatus(resp *http.Response) error {
	// This is added temporarily to ensure we return an error if we get a non-success status code.
	// Eventually (ideally), we'll be generating this error handling as part of the SDK
	if resp.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
