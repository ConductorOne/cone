# PolicyInput

The Policy message.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Description`                                                   | **string*                                                       | :heavy_minus_sign:                                              | The description field.                                          |
| `DisplayName`                                                   | **string*                                                       | :heavy_minus_sign:                                              | The displayName field.                                          |
| `ID`                                                            | **string*                                                       | :heavy_minus_sign:                                              | The id field.                                                   |
| `PolicySteps`                                                   | map[string][PolicySteps](../../models/shared/policysteps.md)    | :heavy_minus_sign:                                              | The policySteps field.                                          |
| `PolicyType`                                                    | [*PolicyPolicyType](../../models/shared/policypolicytype.md)    | :heavy_minus_sign:                                              | The policyType field.                                           |
| `PostActions`                                                   | [][PolicyPostActions](../../models/shared/policypostactions.md) | :heavy_minus_sign:                                              | The postActions field.                                          |
| `ReassignTasksToDelegates`                                      | **bool*                                                         | :heavy_minus_sign:                                              | The reassignTasksToDelegates field.                             |
| `SystemBuiltin`                                                 | **bool*                                                         | :heavy_minus_sign:                                              | The systemBuiltin field.                                        |