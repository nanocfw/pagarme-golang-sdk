
# Get Subscription Boleto Response

Response object for getting a boleto

## Structure

`GetSubscriptionBoletoResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `interest` | [`types.Optional[models.GetInterestResponse]`](../../doc/models/get-interest-response.md) | Optional | Interest |
| `fine` | [`types.Optional[models.GetFineResponse]`](../../doc/models/get-fine-response.md) | Optional | Fine |
| `maxDaysToPayPastDue` | `types.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "interest": {
    "days": 2,
    "type": "percentage",
    "amount": 20
  },
  "fine": {
    "days": 2,
    "type": "flat",
    "amount": 10
  },
  "max_days_to_pay_past_due": 2
}
```

