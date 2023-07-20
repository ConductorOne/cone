# TaskRevokeSource

The TaskRevokeSource message.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `TaskRevokeSourceExpired`                                                    | [*TaskRevokeSourceExpired](../../models/shared/taskrevokesourceexpired.md)   | :heavy_minus_sign:                                                           | The TaskRevokeSourceExpired message.                                         |
| `TaskRevokeSourceNonUsage`                                                   | [*TaskRevokeSourceNonUsage](../../models/shared/taskrevokesourcenonusage.md) | :heavy_minus_sign:                                                           | The TaskRevokeSourceNonUsage message.                                        |
| `TaskRevokeSourceRequest`                                                    | [*TaskRevokeSourceRequest](../../models/shared/taskrevokesourcerequest.md)   | :heavy_minus_sign:                                                           | The TaskRevokeSourceRequest message.                                         |
| `TaskRevokeSourceReview`                                                     | [*TaskRevokeSourceReview](../../models/shared/taskrevokesourcereview.md)     | :heavy_minus_sign:                                                           | The TaskRevokeSourceReview message.                                          |