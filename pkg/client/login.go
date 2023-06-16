package client

import (
	"context"

	conductoroneapi "github.com/conductorone/conductorone-sdk-go"
)

func (c *client) Login(ctx context.Context, tenantID string) (string, string, error) {
	creds, err := conductoroneapi.LoginFlow(ctx, tenantID, ConeClientID)
	if err != nil {
		return "", "", err
	}

	return creds.ClientID, creds.ClientSecret, nil
}
