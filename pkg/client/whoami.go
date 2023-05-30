package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type whoamiResponse struct {
	AccessTokenId string   `json:"access_token_id"`
	TenantId      string   `json:"tenant_id"`
	PrincipleId   string   `json:"principle_id"`
	Roles         []string `json:"roles"`
	Permissions   []string `json:"permissions"`
	UserId        string   `json:"user_id"`
	Features      []string `json:"features"`
}

func (c *client) WhoAmI(ctx context.Context) (*whoamiResponse, error) {
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

	whoamiResp := &whoamiResponse{}
	err = json.Unmarshal(body, whoamiResp)
	if err != nil {
		return nil, err
	}

	return whoamiResp, nil
}
