
# Get Checkout Boleto Payment Response

## Structure

`GetCheckoutBoletoPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `dueAt` | `types.Optional[time.Time]` | Optional | Data de vencimento do boleto |
| `instructions` | `types.Optional[string]` | Optional | Instruções do boleto |

## Example (as JSON)

```json
{
  "due_at": "2016-03-13T12:52:32.123Z",
  "instructions": "instructions2"
}
```

