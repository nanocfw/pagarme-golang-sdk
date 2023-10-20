package controllers

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/nanocfw/pagarme-golang-sdk/models"
)

type CustomersController struct {
	baseController
}

func NewCustomersController(baseController baseController) *CustomersController {
	customersController := CustomersController{baseController: baseController}
	return &customersController
}

// Updates a card
func (c *CustomersController) UpdateCard(
	ctx context.Context,
	customerId string,
	cardId string,
	request *models.UpdateCardRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	var result models.GetCardResponse
	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates an address
func (c *CustomersController) UpdateAddress(
	ctx context.Context,
	customerId string,
	addressId string,
	request *models.UpdateAddressRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	var result models.GetAddressResponse
	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Delete a customer's access token
func (c *CustomersController) DeleteAccessToken(
	ctx context.Context,
	customerId string,
	tokenId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%s/access-tokens/%s", customerId, tokenId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	var result models.GetAccessTokenResponse
	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new customer
func (c *CustomersController) CreateCustomer(
	ctx context.Context,
	request *models.CreateCustomerRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", "/customers")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	var result models.GetCustomerResponse
	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new address for a customer
func (c *CustomersController) CreateAddress(
	ctx context.Context,
	customerId string,
	request *models.CreateAddressRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%s/addresses", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	var result models.GetAddressResponse
	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Delete a Customer's access tokens
func (c *CustomersController) DeleteAccessTokens(ctx context.Context, customerId string) (
	https.ApiResponse[models.ListAccessTokensResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/access-tokens/", customerId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}

	var result models.ListAccessTokensResponse
	result, err = utilities.DecodeResults[models.ListAccessTokensResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get a customer's address
func (c *CustomersController) GetAddress(
	ctx context.Context,
	customerId string,
	addressId string) (
	https.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	var result models.GetAddressResponse
	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Delete a Customer's address
func (c *CustomersController) DeleteAddress(
	ctx context.Context,
	customerId string,
	addressId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	var result models.GetAddressResponse
	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAddressResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new card for a customer
func (c *CustomersController) CreateCard(
	ctx context.Context,
	customerId string,
	request *models.CreateCardRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", fmt.Sprintf("/customers/%s/cards", customerId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	var result models.GetCardResponse
	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get all Customers
func (c *CustomersController) GetCustomers(
	ctx context.Context,
	name *string,
	document *string,
	page *int,
	size *int,
	email *string,
	code *string) (
	https.ApiResponse[models.ListCustomersResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/customers")
	req.Authenticate(true)
	if name != nil {
		req.QueryParam("name", *name)
	}
	if document != nil {
		req.QueryParam("document", *document)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if email != nil {
		req.QueryParam("email", *email)
	}
	if code != nil {
		req.QueryParam("Code", *code)
	}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListCustomersResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListCustomersResponse]{Response: resp}, err
	}

	var result models.ListCustomersResponse
	result, err = utilities.DecodeResults[models.ListCustomersResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListCustomersResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates a customer
func (c *CustomersController) UpdateCustomer(
	ctx context.Context,
	customerId string,
	request *models.UpdateCustomerRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "PUT", fmt.Sprintf("/customers/%s", customerId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	var result models.GetCustomerResponse
	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a access token for a customer
func (c *CustomersController) CreateAccessToken(
	ctx context.Context,
	customerId string,
	request *models.CreateAccessTokenRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%s/access-tokens", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	var result models.GetAccessTokenResponse
	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get all access tokens from a customer
func (c *CustomersController) GetAccessTokens(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	https.ApiResponse[models.ListAccessTokensResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/access-tokens", customerId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}

	var result models.ListAccessTokensResponse
	result, err = utilities.DecodeResults[models.ListAccessTokensResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListAccessTokensResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get all cards from a customer
func (c *CustomersController) GetCards(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	https.ApiResponse[models.ListCardsResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/customers/%s/cards", customerId))
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListCardsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListCardsResponse]{Response: resp}, err
	}

	var result models.ListCardsResponse
	result, err = utilities.DecodeResults[models.ListCardsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListCardsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Renew a card
func (c *CustomersController) RenewCard(
	ctx context.Context,
	customerId string,
	cardId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%s/cards/%s/renew", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	var result models.GetCardResponse
	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get a Customer's access token
func (c *CustomersController) GetAccessToken(
	ctx context.Context,
	customerId string,
	tokenId string) (
	https.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/access-tokens/%s", customerId, tokenId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	var result models.GetAccessTokenResponse
	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAccessTokenResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the metadata a customer
func (c *CustomersController) UpdateCustomerMetadata(
	ctx context.Context,
	customerId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Customers/%s/metadata", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	var result models.GetCustomerResponse
	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Delete a customer's card
func (c *CustomersController) DeleteCard(
	ctx context.Context,
	customerId string,
	cardId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	var result models.GetCardResponse
	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets all adressess from a customer
func (c *CustomersController) GetAddresses(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	https.ApiResponse[models.ListAddressesResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/addresses", customerId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListAddressesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListAddressesResponse]{Response: resp}, err
	}

	var result models.ListAddressesResponse
	result, err = utilities.DecodeResults[models.ListAddressesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListAddressesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get a customer
func (c *CustomersController) GetCustomer(ctx context.Context, customerId string) (
	https.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/customers/%s", customerId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	var result models.GetCustomerResponse
	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCustomerResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get a customer's card
func (c *CustomersController) GetCard(
	ctx context.Context,
	customerId string,
	cardId string) (
	https.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	var result models.GetCardResponse
	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetCardResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
