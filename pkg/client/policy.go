package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) ListPolicies(ctx context.Context) ([]shared.Policy, error) {
	policies := make([]shared.Policy, 0)
	pageSize := float64(100)
	pageToken := ""
	for {
		resp, err := c.sdk.Policies.List(ctx, operations.C1APIPolicyV1PoliciesListRequest{
			PageToken: &pageToken,
			PageSize:  &pageSize,
		})
		if err != nil {
			return nil, err
		}
		if err := handleBadStatus(resp.RawResponse); err != nil {
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
