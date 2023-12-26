package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) AuthIntrospect(ctx context.Context) (*shared.IntrospectResponse, error) {
	resp, err := c.sdk.Auth.Introspect(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.IntrospectResponse, nil
}
