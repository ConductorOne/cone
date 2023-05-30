package client

import (
	"context"
	"net/http"
	"os"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/uhttp"
)

type client struct {
	httpClient *http.Client
	clientName string
	tokenHost  string
	apiClient  *c1api.APIClient
}

type C1Client interface {
	WhoAmI(ctx context.Context) (*whoamiResponse, error)
	GetUser(ctx context.Context, userID string) (*UserResponse, error)
}

func New(ctx context.Context, clientId string, clientSecret string) (C1Client, error) {
	tokenSrc, clientName, tokenHost, err := NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return nil, err
	}
	uclient, err := uhttp.NewClient(ctx, uhttp.WithTokenSource(tokenSrc))
	if err != nil {
		return nil, err
	}
	c := &client{
		tokenHost:  tokenHost,
		clientName: clientName,
		httpClient: uclient,
	}

	apiCfg := c1api.NewConfiguration()
	apiCfg.HTTPClient = uclient

	// If the API host is set in the environment, use that instead of the default
	// HACK(jirwin): Instead of using the generated client's server address, use the hostname from the token.
	if apiHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok {
		apiCfg.Servers[0].URL = apiHost
	} else {
		apiCfg.Servers[0].URL = c.tokenHost
	}

	c.apiClient = c1api.NewAPIClient(apiCfg)

	return c, nil
}

func (c *client) apiHost() string {
	if envHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok {
		return envHost
	}
	return c.tokenHost
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
