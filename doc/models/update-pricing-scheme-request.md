
# Update Pricing Scheme Request

Request for updating a pricing scheme

## Structure

`UpdatePricingSchemeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `schemeType` | `string` | Required | Scheme type |
| `priceBrackets` | [`[]models.UpdatePriceBracketRequest`](../../doc/models/update-price-bracket-request.md) | Required | Price brackets |
| `price` | `*int` | Optional | Price |
| `minimumPrice` | `*int` | Optional | Minimum price |
| `percentage` | `*float64` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "scheme_type": "scheme_type0",
  "price_brackets": [
    {
      "start_quantity": 193,
      "price": 125,
      "end_quantity": 201,
      "overage_price": 215
    }
  ],
  "price": 16,
  "minimum_price": 176,
  "percentage": 4.18
}
```

