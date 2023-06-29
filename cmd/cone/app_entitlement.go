package main

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
)

type Entitlement shared.AppEntitlement

func (r *Entitlement) Header() []string {
	return []string{
		"Entitlement",
		"Alias",
		"Description",
	}
}

func (r *Entitlement) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.DisplayName),
		client.StringFromPtr(r.Description),
		client.StringFromPtr(r.Alias),
	}}
}
