package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) ListPolicies(ctx context.Context) ([]shared.Policy, error) {
	policies := make([]shared.Policy, 0)
	pageSize := float64(100)
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

		policies = append(policies, resp.ListPolicyResponse.List...)

		if resp.ListPolicyResponse.NextPageToken == nil || *resp.ListPolicyResponse.NextPageToken == "" {
			break
		}
		pageToken = *resp.ListPolicyResponse.NextPageToken
	}

	return policies, nil
}
