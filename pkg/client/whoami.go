package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) WhoAmI(ctx context.Context) (*c1api.C1ApiAuthV1IntrospectResponse, error) {
	introspectResp, resp, err := c.apiClient.DefaultAPI.C1ApiAuthV1AuthIntrospect(ctx).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return introspectResp, nil
}
