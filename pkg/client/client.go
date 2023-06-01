package client

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/uhttp"
)

type client struct {
	httpClient *http.Client
	clientName string
	tokenHost  string
	apiClient  *c1api.APIClient
	config     clientConfig
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

func float32Ptr(i int) *float32 {
	f := float32(i)
	return &f
}

type C1Client interface {
	WhoAmI(ctx context.Context) (*c1api.C1ApiAuthV1IntrospectResponse, error)
	GetUser(ctx context.Context, userID string) (*c1api.C1ApiUserV1UserServiceGetResponse, error)
	SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) (*c1api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse, error)
	GetResource(ctx context.Context, appID string, resourceID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceServiceGetResponse, error)
	GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceTypeServiceGetResponse, error)
	GetApp(ctx context.Context, appID string) (*c1api.C1ApiAppV1App, error)
	CreateGrantTask(ctx context.Context, appId string, appEntitlementId string, identityUserId string) (*c1api.C1ApiTaskV1TaskServiceGetResponse, error)
}

func New(
	ctx context.Context,
	clientId string,
	clientSecret string,
	optionFuncs ...ClientOptionFunc,
) (C1Client, error) {
	tokenSrc, clientName, tokenHost, err := NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	opt := applyOpts(optionFuncs)
	uclient, err := uhttp.NewClient(ctx, uhttp.WithTokenSource(tokenSrc), uhttp.WithDebug(opt.Debug))
	if err != nil {
		return nil, err
	}
	c := &client{
		tokenHost:  tokenHost,
		clientName: clientName,
		httpClient: uclient,
		config:     opt,
	}

	apiCfg := c1api.NewConfiguration()
	apiCfg.HTTPClient = uclient

	var apiHostname string
	// If the API host is set in the environment, use that instead of the default
	// HACK(jirwin): Instead of using the generated client's server address, use the hostname from the token.
	if apiHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok {
		apiHostname = apiHost
	} else {
		apiHostname = c.tokenHost
	}
	apiURL := url.URL{
		Scheme: "https",
		Host:   apiHostname,
	}
	apiCfg.Servers[0].URL = apiURL.String()
	c.apiClient = c1api.NewAPIClient(apiCfg)

	return c, nil
}

// The c1api client uses the context to set various configuration options. Do that here.
func (c *client) GetContext(ctx context.Context) context.Context {
	// If the API host is set in the environment, we don't need to populate any server variables
	if _, ok := os.LookupEnv("CONE_API_ENDPOINT"); !ok {
		return ctx
	}

	// TODO(jirwin): if we choose to use this, we will need to parse the tenant name out of the token, and set it as `tenantHost` here.
	// serverVars := map[string]string{
	// 	"tenantHost": c.tokenHost,
	// }
	// return context.WithValue(ctx, c1api.ContextServerVariables, c.clientName)
	return ctx
}
