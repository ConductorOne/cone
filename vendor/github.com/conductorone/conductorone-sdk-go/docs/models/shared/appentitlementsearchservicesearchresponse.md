# AppEntitlementSearchServiceSearchResponse

The AppEntitlementSearchServiceSearchResponse message.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Facets`                                                          | [*Facets](../../models/shared/facets.md)                          | :heavy_minus_sign:                                                | The Facets message.                                               |
| `Expanded`                                                        | []map[string]*interface{}*                                        | :heavy_minus_sign:                                                | The expanded field.                                               |
| `List`                                                            | [][AppEntitlementView](../../models/shared/appentitlementview.md) | :heavy_minus_sign:                                                | The list field.                                                   |
| `NextPageToken`                                                   | **string*                                                         | :heavy_minus_sign:                                                | The nextPageToken field.                                          |