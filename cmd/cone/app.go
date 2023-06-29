package main

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
)

type App shared.App

func (r *App) Header() []string {
	return []string{
		"App Name",
		"App ID",
		"Description",
	}
}

func (r *App) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.DisplayName),
		client.StringFromPtr(r.ID),
		client.StringFromPtr(r.Description),
	}}
}
