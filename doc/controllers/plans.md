# Plans

```go
plansController := client.PlansController
```

## Class Name

`PlansController`

## Methods

* [Get Plan](../../doc/controllers/plans.md#get-plan)
* [Delete Plan](../../doc/controllers/plans.md#delete-plan)
* [Update Plan Metadata](../../doc/controllers/plans.md#update-plan-metadata)
* [Update Plan Item](../../doc/controllers/plans.md#update-plan-item)
* [Create Plan Item](../../doc/controllers/plans.md#create-plan-item)
* [Get Plan Item](../../doc/controllers/plans.md#get-plan-item)
* [Create Plan](../../doc/controllers/plans.md#create-plan)
* [Delete Plan Item](../../doc/controllers/plans.md#delete-plan-item)
* [Get Plans](../../doc/controllers/plans.md#get-plans)
* [Update Plan](../../doc/controllers/plans.md#update-plan)


# Get Plan

Gets a plan

```go
GetPlan(
    planId string) (
    https.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
planId := "plan_id8"

apiResponse, err := plansController.GetPlan(planId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Plan

Deletes a plan

```go
DeletePlan(
    planId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
planId := "plan_id8"

apiResponse, err := plansController.DeletePlan(planId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan Metadata

Updates the metadata from a plan

```go
UpdatePlanMetadata(
    planId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | The plan id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the plan metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
planId := "plan_id8"

request := models.UpdateMetadataRequest{
    Metadata: map[string]*string{
"key0" : "metadata3",
},
}

apiResponse, err := plansController.UpdatePlanMetadata(planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan Item

Updates a plan item

```go
UpdatePlanItem(
    planId string,
    planItemId string,
    body models.UpdatePlanItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |
| `body` | [`models.UpdatePlanItemRequest`](../../doc/models/update-plan-item-request.md) | Body, Required | Request for updating the plan item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
planId := "plan_id8"
planItemId := "plan_item_id0"

bodyPricingSchemePriceBrackets0 := models.UpdatePriceBracketRequest{
    StartQuantity: 31,
    Price:         225,
}

bodyPricingSchemePriceBrackets := []models.UpdatePriceBracketRequest{bodyPricingSchemePriceBrackets0}
bodyPricingScheme := models.UpdatePricingSchemeRequest{
    SchemeType:    "scheme_type2",
    PriceBrackets: bodyPricingSchemePriceBrackets,
}

body := models.UpdatePlanItemRequest{
    Name:          "name6",
    Description:   "description4",
    Status:        "status2",
    PricingScheme: bodyPricingScheme,
}

apiResponse, err := plansController.UpdatePlanItem(planId, planItemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Plan Item

Adds a new item to a plan

```go
CreatePlanItem(
    planId string,
    request models.CreatePlanItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `request` | [`models.CreatePlanItemRequest`](../../doc/models/create-plan-item-request.md) | Body, Required | Request for creating a plan item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
planId := "plan_id8"

requestPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type2",
}

request := models.CreatePlanItemRequest{
    Name:          "name6",
    Id:            "id6",
    Description:   "description6",
    PricingScheme: requestPricingScheme,
}

apiResponse, err := plansController.CreatePlanItem(planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Plan Item

Gets a plan item

```go
GetPlanItem(
    planId string,
    planItemId string) (
    https.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
planId := "plan_id8"
planItemId := "plan_item_id0"

apiResponse, err := plansController.GetPlanItem(planId, planItemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Plan

Creates a new plan

```go
CreatePlan(
    body models.CreatePlanRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreatePlanRequest`](../../doc/models/create-plan-request.md) | Body, Required | Request for creating a plan |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
bodyItems0PricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type5",
}

bodyItems0 := models.CreatePlanItemRequest{
    Name:          "name3",
    Id:            "id3",
    Description:   "description3",
    PricingScheme: bodyItems0PricingScheme,
}
bodyItems := []models.CreatePlanItemRequest{bodyItems0}
bodyPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type2",
}

body := models.CreatePlanRequest{
    Name:                "name6",
    Description:         "description4",
    StatementDescriptor: "statement_descriptor6",
    Shippable:           false,
    PaymentMethods:      []string{"payment_methods9"},
    Installments:        []int{207},
    Currency:            "currency6",
    Interval:            "interval6",
    IntervalCount:       170,
    BillingDays:         []int{201, 200},
    BillingType:         "billing_type0",
    Metadata:            map[string]*string{
"key0" : "metadata7",
"key1" : "metadata8",
},
    Items:               bodyItems,
    PricingScheme:       bodyPricingScheme,
}

apiResponse, err := plansController.CreatePlan(&body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Plan Item

Removes an item from a plan

```go
DeletePlanItem(
    planId string,
    planItemId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
planId := "plan_id8"
planItemId := "plan_item_id0"

apiResponse, err := plansController.DeletePlanItem(planId, planItemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Plans

Gets all plans

```go
GetPlans(
    page *int,
    size *int,
    name *string,
    status *string,
    billingType *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    https.ApiResponse[models.ListPlansResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `name` | `*string` | Query, Optional | Filter for Plan's name |
| `status` | `*string` | Query, Optional | Filter for Plan's status |
| `billingType` | `*string` | Query, Optional | Filter for plan's billing type |
| `createdSince` | `*time.Time` | Query, Optional | Filter for plan's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for plan's creation date end range |

## Response Type

[`models.ListPlansResponse`](../../doc/models/list-plans-response.md)

## Example Usage

```go
apiResponse, err := plansController.GetPlans(nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan

Updates a plan

```go
UpdatePlan(
    planId string,
    request models.UpdatePlanRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `request` | [`models.UpdatePlanRequest`](../../doc/models/update-plan-request.md) | Body, Required | Request for updating a plan |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
planId := "plan_id8"

request := models.UpdatePlanRequest{
    Name:                "name6",
    Description:         "description6",
    Installments:        []int{151, 152},
    StatementDescriptor: "statement_descriptor6",
    Currency:            "currency6",
    Interval:            "interval4",
    IntervalCount:       114,
    PaymentMethods:      []string{"payment_methods1", "payment_methods0", "payment_methods9"},
    BillingType:         "billing_type0",
    Status:              "status8",
    Shippable:           false,
    BillingDays:         []int{115},
    Metadata:            map[string]*string{
"key0" : "metadata3",
},
}

apiResponse, err := plansController.UpdatePlan(planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

