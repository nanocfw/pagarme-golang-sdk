
# List Subscriptions Response

Response object for listing subscriptions

## Structure

`ListSubscriptionsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `data` | [`types.Optional[[]models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | The subscription objects |
| `paging` | [`types.Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "code": "code3",
      "start_at": "2016-03-13T12:52:32.123Z",
      "interval": "interval3",
      "interval_count": 249
    },
    {
      "id": "id6",
      "code": "code4",
      "start_at": "2016-03-13T12:52:32.123Z",
      "interval": "interval4",
      "interval_count": 250
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

