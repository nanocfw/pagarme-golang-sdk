
# Get Phones Response

## Structure

`GetPhonesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `homePhone` | [`types.Optional[models.GetPhoneResponse]`](../../doc/models/get-phone-response.md) | Optional | - |
| `mobilePhone` | [`types.Optional[models.GetPhoneResponse]`](../../doc/models/get-phone-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "home_phone": {
    "country_code": "country_code0",
    "number": "number2",
    "area_code": "area_code0"
  },
  "mobile_phone": {
    "country_code": "country_code0",
    "number": "number8",
    "area_code": "area_code0"
  }
}
```

