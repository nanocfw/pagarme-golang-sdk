package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/nanocfw/pagarme-golang-sdk/models"
)

type ChargesController struct {
	baseController
}

func NewChargesController(baseController baseController) *ChargesController {
	chargesController := ChargesController{baseController: baseController}
	return &chargesController
}

// Updates the metadata from a charge
func (c *ChargesController) UpdateChargeMetadata(
	ctx context.Context,
	chargeId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "PATCH", fmt.Sprintf("/Charges/%s/metadata", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates a charge's payment method
func (c *ChargesController) UpdateChargePaymentMethod(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargePaymentMethodRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/charges/%s/payment-method", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) GetChargeTransactions(
	ctx context.Context,
	chargeId string,
	page *int,
	size *int) (
	https.ApiResponse[models.ListChargeTransactionsResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/charges/%s/transactions", chargeId),
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
		return https.ApiResponse[models.ListChargeTransactionsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListChargeTransactionsResponse]{Response: resp}, err
	}

	var result models.ListChargeTransactionsResponse
	result, err = utilities.DecodeResults[models.ListChargeTransactionsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListChargeTransactionsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the due date from a charge
func (c *ChargesController) UpdateChargeDueDate(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargeDueDateRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "PATCH", fmt.Sprintf("/Charges/%s/due-date", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Lists all charges
func (c *ChargesController) GetCharges(
	ctx context.Context,
	page *int,
	size *int,
	code *string,
	status *string,
	paymentMethod *string,
	customerId *string,
	orderId *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListChargesResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/charges")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if code != nil {
		req.QueryParam("code", *code)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if paymentMethod != nil {
		req.QueryParam("payment_method", *paymentMethod)
	}
	if customerId != nil {
		req.QueryParam("customer_id", *customerId)
	}
	if orderId != nil {
		req.QueryParam("order_id", *orderId)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListChargesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListChargesResponse]{Response: resp}, err
	}

	var result models.ListChargesResponse
	result, err = utilities.DecodeResults[models.ListChargesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListChargesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Captures a charge
func (c *ChargesController) CaptureCharge(
	ctx context.Context,
	chargeId string,
	request *models.CreateCaptureChargeRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", fmt.Sprintf("/charges/%s/capture", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the card from a charge
func (c *ChargesController) UpdateChargeCard(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargeCardRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "PATCH", fmt.Sprintf("/charges/%s/card", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get a charge from its id
func (c *ChargesController) GetCharge(ctx context.Context, chargeId string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/charges/%s", chargeId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) GetChargesSummary(
	ctx context.Context,
	status string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.GetChargesSummaryResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/charges/summary")
	req.Authenticate(true)
	req.QueryParam("status", status)
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargesSummaryResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargesSummaryResponse]{Response: resp}, err
	}

	var result models.GetChargesSummaryResponse
	result, err = utilities.DecodeResults[models.GetChargesSummaryResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargesSummaryResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Retries a charge
func (c *ChargesController) RetryCharge(
	ctx context.Context,
	chargeId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", fmt.Sprintf("/charges/%s/retry", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Cancel a charge
func (c *ChargesController) CancelCharge(
	ctx context.Context,
	chargeId string,
	request *models.CreateCancelChargeRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "DELETE", fmt.Sprintf("/charges/%s", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new charge
func (c *ChargesController) CreateCharge(
	ctx context.Context,
	request *models.CreateChargeRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", "/Charges")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) ConfirmPayment(
	ctx context.Context,
	chargeId string,
	request *models.CreateConfirmPaymentRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/charges/%s/confirm-payment", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	var result models.GetChargeResponse
	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetChargeResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
