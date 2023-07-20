# ConnectorStatus

The ConnectorStatus message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CompletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `LastError`                                                            | **string*                                                              | :heavy_minus_sign:                                                     | The lastError field.                                                   |
| `StartedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Status`                                                               | [*ConnectorStatusStatus](../../models/shared/connectorstatusstatus.md) | :heavy_minus_sign:                                                     | The status field.                                                      |
| `UpdatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |