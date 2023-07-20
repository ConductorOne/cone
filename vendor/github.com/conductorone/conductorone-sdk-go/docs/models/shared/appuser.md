# AppUser

The AppUser message.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AppUserStatus`                                                  | [*AppUserStatus](../../models/shared/appuserstatus.md)           | :heavy_minus_sign:                                               | The AppUserStatus message.                                       |
| `AppID`                                                          | **string*                                                        | :heavy_minus_sign:                                               | The appId field.                                                 |
| `AppUserType`                                                    | [*AppUserAppUserType](../../models/shared/appuserappusertype.md) | :heavy_minus_sign:                                               | The appUserType field.                                           |
| `CreatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `DeletedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `DisplayName`                                                    | **string*                                                        | :heavy_minus_sign:                                               | The displayName field.                                           |
| `Email`                                                          | **string*                                                        | :heavy_minus_sign:                                               | The email field.                                                 |
| `ID`                                                             | **string*                                                        | :heavy_minus_sign:                                               | The id field.                                                    |
| `IdentityUserID`                                                 | **string*                                                        | :heavy_minus_sign:                                               | The identityUserId field.                                        |
| `Profile`                                                        | map[string]*interface{}*                                         | :heavy_minus_sign:                                               | N/A                                                              |
| `UpdatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |