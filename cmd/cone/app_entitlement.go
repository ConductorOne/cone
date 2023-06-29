package main

import (
	"fmt"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
)

type Entitlement struct {
	entitlement *shared.AppEntitlement
	client      client.C1Client
}

func (r *Entitlement) Pretext() string {
	return fmt.Sprintf("Entitlement URL: %s/applications/%s/entitlements", r.client.BaseURL(), client.StringFromPtr(r.entitlement.AppID))
}

func (r *Entitlement) Header() []string {
	return []string{
		"Entitlement",
		"Description",
		"Slug",
	}
}

func (r *Entitlement) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.entitlement.DisplayName),
		client.StringFromPtr(r.entitlement.Description),
		client.StringFromPtr(r.entitlement.Slug),
	}}
}
