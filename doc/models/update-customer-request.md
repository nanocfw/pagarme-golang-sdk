
# Update Customer Request

Request for updating a customer

## Structure

`UpdateCustomerRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `*string` | Optional | Name |
| `email` | `*string` | Optional | Email |
| `document` | `*string` | Optional | Document number |
| `mType` | `*string` | Optional | Person type |
| `address` | [`*models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Optional | Address |
| `metadata` | `map[string]*string` | Optional | Metadata |
| `phones` | [`*models.CreatePhonesRequest`](../../doc/models/create-phones-request.md) | Optional | - |
| `code` | `*string` | Optional | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `gender` | `*string` | Optional | Gênero do cliente |
| `documentType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name0",
  "email": "email6",
  "document": "document6",
  "type": "type0",
  "address": {
    "street": "street6",
    "number": "number4",
    "zip_code": "zip_code0",
    "neighborhood": "neighborhood2",
    "city": "city6",
    "state": "state2",
    "country": "country0",
    "complement": "complement2",
    "metadata": {
      "key0": "metadata3",
      "key1": "metadata2",
      "key2": "metadata1"
    },
    "line_1": "line_10",
    "line_2": "line_24"
  }
}
```

