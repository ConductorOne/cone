package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) ListPolicies(ctx context.Context) ([]shared.Policy, error) {
	policies := make([]shared.Policy, 0)
	pageSize := 100
	pageToken := ""
	for {
		resp, err := c.sdk.PolicySearch.Search(ctx, &shared.SearchPoliciesRequest{
			PageSize:  &pageSize,
			PageToken: &pageToken,
		})
		if err != nil {
			return nil, err
		}
		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}

		policies = append(policies, resp.SearchPoliciesResponse.List...)

		if resp.SearchPoliciesResponse.NextPageToken == nil || *resp.SearchPoliciesResponse.NextPageToken == "" {
			break
		}
		pageToken = *resp.SearchPoliciesResponse.NextPageToken
	}

	return policies, nil
}
