package client

import (
	"context"
)

func (c *client) AuthIntrospect(ctx context.Context) (*c1api.C1ApiAuthV1IntrospectResponse, error) {
	introspectResp, resp, err := c.apiClient.DefaultAPI.C1ApiAuthV1AuthIntrospect(ctx).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return introspectResp, nil
}
