
# Create Emv Data Decrypt Request

## Structure

`CreateEmvDataDecryptRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `cipher` | `string` | Required | Emv Decrypt cipher type |
| `dukpt` | [`*models.CreateEmvDataDukptDecryptRequest`](../../doc/models/create-emv-data-dukpt-decrypt-request.md) | Optional | Dukpt data request |
| `tags` | [`[]models.CreateEmvDataTlvDecryptRequest`](../../doc/models/create-emv-data-tlv-decrypt-request.md) | Required | Encrypted tags list |

## Example (as JSON)

```json
{
  "cipher": "cipher4",
  "tags": [
    {
      "tag": "tag9",
      "lenght": "lenght7",
      "value": "value7"
    }
  ],
  "dukpt": {
    "ksn": "ksn0"
  }
}
```

