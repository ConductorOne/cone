package client

import (
	"context"
	"net/http"
	"os"

	"github.com/conductorone/cone/pkg/uhttp"
)

type client struct {
	httpClient *http.Client
	clientName string
	tokenHost  string
}

type C1Client interface {
	WhoAmI(ctx context.Context) (interface{}, error)
	GetUser(ctx context.Context, userID string) (interface{}, error)
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

	return c, nil
}

func (c *client) apiHost() string {
	if envHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok {
		return envHost
	}
	return c.tokenHost
}
