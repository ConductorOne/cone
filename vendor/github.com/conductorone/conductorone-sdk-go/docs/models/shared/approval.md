# Approval

The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppGroupApproval`                                                           | [*AppGroupApproval](../../models/shared/appgroupapproval.md)                 | :heavy_minus_sign:                                                           | The AppGroupApproval message.                                                |
| `AppOwnerApproval`                                                           | [*AppOwnerApproval](../../models/shared/appownerapproval.md)                 | :heavy_minus_sign:                                                           | The AppOwnerApproval message.                                                |
| `EntitlementOwnerApproval`                                                   | [*EntitlementOwnerApproval](../../models/shared/entitlementownerapproval.md) | :heavy_minus_sign:                                                           | The EntitlementOwnerApproval message.                                        |
| `ManagerApproval`                                                            | [*ManagerApproval](../../models/shared/managerapproval.md)                   | :heavy_minus_sign:                                                           | The ManagerApproval message.                                                 |
| `SelfApproval`                                                               | [*SelfApproval](../../models/shared/selfapproval.md)                         | :heavy_minus_sign:                                                           | The SelfApproval message.                                                    |
| `UserApproval`                                                               | [*UserApproval](../../models/shared/userapproval.md)                         | :heavy_minus_sign:                                                           | The UserApproval message.                                                    |
| `AllowReassignment`                                                          | **bool*                                                                      | :heavy_minus_sign:                                                           | The allowReassignment field.                                                 |
| `Assigned`                                                                   | **bool*                                                                      | :heavy_minus_sign:                                                           | The assigned field.                                                          |
| `RequireApprovalReason`                                                      | **bool*                                                                      | :heavy_minus_sign:                                                           | The requireApprovalReason field.                                             |
| `RequireReassignmentReason`                                                  | **bool*                                                                      | :heavy_minus_sign:                                                           | The requireReassignmentReason field.                                         |