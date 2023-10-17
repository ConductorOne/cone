<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "",
			Oauth:      "",
		}),
	)

	ctx := context.Background()
	res, err := s.AppEntitlementOwners.Add(ctx, operations.C1APIAppV1AppEntitlementOwnersAddRequest{
		AddAppEntitlementOwnerRequest: &shared.AddAppEntitlementOwnerRequest{},
		AppID:                         "program",
		EntitlementID:                 "North",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddAppEntitlementOwnerResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->