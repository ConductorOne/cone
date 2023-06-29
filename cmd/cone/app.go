package main

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
)

type App shared.App

func (r *App) Header() []string {
	return []string{
		"App Name",
		"Description",
	}
}

func (r *App) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.DisplayName),
		client.StringFromPtr(r.Description),
	}}
}
