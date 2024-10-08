
# List Charge Transactions Response

Response object for listing charge transactions

## Structure

`ListChargeTransactionsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `data` | [`types.Optional[[]models.GetTransactionResponseInterface]`](../../doc/models/get-transaction-response.md) | Optional | The charge transactions objects |
| `paging` | [`types.Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "gateway_id": "gateway_id5",
      "amount": 121,
      "status": "status7",
      "success": true,
      "created_at": "2016-03-13T12:52:32.123Z",
      "transaction_type": "transaction"
    },
    {
      "gateway_id": "gateway_id6",
      "amount": 122,
      "status": "status8",
      "success": false,
      "created_at": "2016-03-13T12:52:32.123Z",
      "transaction_type": "transaction"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

