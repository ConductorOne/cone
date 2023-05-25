package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (c *client) WhoAmI(ctx context.Context) (interface{}, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.apiHost(),
		Path:   "/api/v1/auth/introspect",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	return nil, nil
}

func (c *client) GetUser(ctx context.Context, userID string) (interface{}, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.apiHost(),
		Path:   fmt.Sprintf("/api/v1/user/get/%s", userID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	return nil, nil
}

func NewC1Client(ctx context.Context, clientId string, clientSecret string) (C1Client, error) {
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
	if envHost, ok := os.LookupEnv("CONE_C1_API_HOST"); ok {
		return envHost
	}
	return c.tokenHost
}
