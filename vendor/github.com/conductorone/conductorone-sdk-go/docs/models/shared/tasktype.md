# TaskType

The TaskType message.

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify



## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `TaskTypeCertify`                                          | [*TaskTypeCertify](../../models/shared/tasktypecertify.md) | :heavy_minus_sign:                                         | The TaskTypeCertify message.                               |
| `TaskTypeGrant`                                            | [*TaskTypeGrant](../../models/shared/tasktypegrant.md)     | :heavy_minus_sign:                                         | The TaskTypeGrant message.                                 |
| `TaskTypeRevoke`                                           | [*TaskTypeRevoke](../../models/shared/tasktyperevoke.md)   | :heavy_minus_sign:                                         | The TaskTypeRevoke message.                                |