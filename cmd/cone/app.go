package main

import (
	"fmt"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
)

type App struct {
	app    *shared.App
	client client.C1Client
}

func (r *App) Pretext() string {
	return fmt.Sprintf("App URL: %s/applications/%s", r.client.BaseURL(), client.StringFromPtr(r.app.ID))
}

func (r *App) Header() []string {
	return []string{
		"App Name",
		"Description",
	}
}

func (r *App) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.app.DisplayName),
		client.StringFromPtr(r.app.Description),
	}}
}
