# FacetCategory

The FacetCategory message.

This message contains a oneof named item. Only a single field of the following list may be set at a time:
  - value
  - range



## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `FacetRangeItem`                                         | [*FacetRangeItem](../../models/shared/facetrangeitem.md) | :heavy_minus_sign:                                       | The FacetRangeItem message.                              |
| `FacetValueItem`                                         | [*FacetValueItem](../../models/shared/facetvalueitem.md) | :heavy_minus_sign:                                       | The FacetValueItem message.                              |
| `DisplayName`                                            | **string*                                                | :heavy_minus_sign:                                       | The displayName field.                                   |
| `IconURL`                                                | **string*                                                | :heavy_minus_sign:                                       | The iconUrl field.                                       |
| `Param`                                                  | **string*                                                | :heavy_minus_sign:                                       | The param field.                                         |