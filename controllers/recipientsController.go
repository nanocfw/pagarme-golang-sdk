package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/nanocfw/pagarme-golang-sdk/models"
)

type RecipientsController struct {
	baseController
}

func NewRecipientsController(baseController baseController) *RecipientsController {
	recipientsController := RecipientsController{baseController: baseController}
	return &recipientsController
}

// Updates a recipient
func (r *RecipientsController) UpdateRecipient(
	ctx context.Context,
	recipientId string,
	request *models.UpdateRecipientRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "PUT", fmt.Sprintf("/recipients/%s", recipientId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates an anticipation
func (r *RecipientsController) CreateAnticipation(
	ctx context.Context,
	recipientId string,
	request *models.CreateAnticipationRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%s/anticipations", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}

	var result models.GetAnticipationResponse
	result, err = utilities.DecodeResults[models.GetAnticipationResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets the anticipation limits for a recipient
func (r *RecipientsController) GetAnticipationLimits(
	ctx context.Context,
	recipientId string,
	timeframe string,
	paymentDate time.Time) (
	https.ApiResponse[models.GetAnticipationLimitResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/anticipation_limits", recipientId),
	)
	req.Authenticate(true)
	req.QueryParam("timeframe", timeframe)
	req.QueryParam("payment_date", paymentDate.Format(time.RFC3339))

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAnticipationLimitResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationLimitResponse]{Response: resp}, err
	}

	var result models.GetAnticipationLimitResponse
	result, err = utilities.DecodeResults[models.GetAnticipationLimitResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationLimitResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Retrieves paginated recipients information
func (r *RecipientsController) GetRecipients(
	ctx context.Context,
	page *int,
	size *int) (
	https.ApiResponse[models.ListRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/recipients")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListRecipientResponse]{Response: resp}, err
	}

	var result models.ListRecipientResponse
	result, err = utilities.DecodeResults[models.ListRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) GetWithdrawById(
	ctx context.Context,
	recipientId string,
	withdrawalId string) (
	https.ApiResponse[models.GetWithdrawResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/withdrawals/%s", recipientId, withdrawalId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}

	var result models.GetWithdrawResponse
	result, err = utilities.DecodeResults[models.GetWithdrawResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the default bank account from a recipient
func (r *RecipientsController) UpdateRecipientDefaultBankAccount(
	ctx context.Context,
	recipientId string,
	request *models.UpdateRecipientBankAccountRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%s/default-bank-account", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates recipient metadata
func (r *RecipientsController) UpdateRecipientMetadata(
	ctx context.Context,
	recipientId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%s/metadata", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetTransfers(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/transfers", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListTransferResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListTransferResponse]{Response: resp}, err
	}

	var result models.ListTransferResponse
	result, err = utilities.DecodeResults[models.ListTransferResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListTransferResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a transfer
func (r *RecipientsController) GetTransfer(
	ctx context.Context,
	recipientId string,
	transferId string) (
	https.ApiResponse[models.GetTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/transfers/%s", recipientId, transferId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}

	var result models.GetTransferResponse
	result, err = utilities.DecodeResults[models.GetTransferResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) CreateWithdraw(
	ctx context.Context,
	recipientId string,
	request *models.CreateWithdrawRequest) (
	https.ApiResponse[models.GetWithdrawResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%s/withdrawals", recipientId),
	)
	req.Authenticate(true)
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}

	var result models.GetWithdrawResponse
	result, err = utilities.DecodeResults[models.GetWithdrawResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetWithdrawResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates recipient metadata
func (r *RecipientsController) UpdateAutomaticAnticipationSettings(
	ctx context.Context,
	recipientId string,
	request *models.UpdateAutomaticAnticipationSettingsRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%s/automatic-anticipation-settings", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets an anticipation
func (r *RecipientsController) GetAnticipation(
	ctx context.Context,
	recipientId string,
	anticipationId string) (
	https.ApiResponse[models.GetAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/anticipations/%s", recipientId, anticipationId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}

	var result models.GetAnticipationResponse
	result, err = utilities.DecodeResults[models.GetAnticipationResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetAnticipationResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) UpdateRecipientTransferSettings(
	ctx context.Context,
	recipientId string,
	request *models.UpdateTransferSettingsRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%s/transfer-settings", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Retrieves a paginated list of anticipations from a recipient
func (r *RecipientsController) GetAnticipations(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	timeframe *string,
	paymentDateSince *time.Time,
	paymentDateUntil *time.Time,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/anticipations", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if timeframe != nil {
		req.QueryParam("timeframe", *timeframe)
	}
	if paymentDateSince != nil {
		req.QueryParam("payment_date_since", paymentDateSince.Format(time.RFC3339))
	}
	if paymentDateUntil != nil {
		req.QueryParam("payment_date_until", paymentDateUntil.Format(time.RFC3339))
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListAnticipationResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListAnticipationResponse]{Response: resp}, err
	}

	var result models.ListAnticipationResponse
	result, err = utilities.DecodeResults[models.ListAnticipationResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListAnticipationResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Retrieves recipient information
func (r *RecipientsController) GetRecipient(ctx context.Context, recipientId string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", fmt.Sprintf("/recipients/%s", recipientId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get balance information for a recipient
func (r *RecipientsController) GetBalance(ctx context.Context, recipientId string) (
	https.ApiResponse[models.GetBalanceResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/balance", recipientId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetBalanceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetBalanceResponse]{Response: resp}, err
	}

	var result models.GetBalanceResponse
	result, err = utilities.DecodeResults[models.GetBalanceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetBalanceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetWithdrawals(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListWithdrawals],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%s/withdrawals", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListWithdrawals]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListWithdrawals]{Response: resp}, err
	}

	var result models.ListWithdrawals
	result, err = utilities.DecodeResults[models.ListWithdrawals](decoder)
	if err != nil {
		return https.ApiResponse[models.ListWithdrawals]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a transfer for a recipient
func (r *RecipientsController) CreateTransfer(
	ctx context.Context,
	recipientId string,
	request *models.CreateTransferRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%s/transfers", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}

	var result models.GetTransferResponse
	result, err = utilities.DecodeResults[models.GetTransferResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTransferResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new recipient
func (r *RecipientsController) CreateRecipient(
	ctx context.Context,
	request *models.CreateRecipientRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/recipients")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Retrieves recipient information
func (r *RecipientsController) GetRecipientByCode(ctx context.Context, code string) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", fmt.Sprintf("/recipients/%s", code))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) GetDefaultRecipient(ctx context.Context) (
	https.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/recipients/default")
	req.Authenticate(true)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	var result models.GetRecipientResponse
	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetRecipientResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
