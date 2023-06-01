package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/output"
)

func (c *client) GetApp(ctx context.Context, appID string) (*C1ApiAppV1App, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppsGet(ctx, appID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	app := C1ApiAppV1App(*resp)
	return &app, nil
}

func AppCacheKey(appID *string) string {
	return output.FromPtr(appID)
}

type C1ApiAppV1App c1api.C1ApiAppV1App

func (a *C1ApiAppV1App) CacheKey() string {
	return output.FromPtr(a.Id)
}
