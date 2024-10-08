
# Get Checkout Payment Settings Response

Checkout Payment Settings Response

## Structure

`GetCheckoutPaymentSettingsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `successUrl` | `types.Optional[string]` | Optional | Success Url |
| `paymentUrl` | `types.Optional[string]` | Optional | Payment Url |
| `acceptedPaymentMethods` | `types.Optional[[]string]` | Optional | Accepted Payment Methods |
| `status` | `types.Optional[string]` | Optional | Status |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | Customer |
| `amount` | `types.Optional[int]` | Optional | Payment amount |
| `defaultPaymentMethod` | `types.Optional[string]` | Optional | Default Payment Method |
| `gatewayAffiliationId` | `types.Optional[string]` | Optional | Gateway Affiliation Id |

## Example (as JSON)

```json
{
  "success_url": "success_url2",
  "payment_url": "payment_url6",
  "accepted_payment_methods": [
    "accepted_payment_methods3",
    "accepted_payment_methods4",
    "accepted_payment_methods5"
  ],
  "status": "status8",
  "customer": {
    "id": "id0",
    "name": "name0",
    "email": "email6",
    "delinquent": false,
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

