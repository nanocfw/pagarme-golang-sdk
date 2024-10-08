package controllers

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/nanocfw/pagarme-golang-sdk/models"
)

type TransfersController struct {
	baseController
}

func NewTransfersController(baseController baseController) *TransfersController {
	transfersController := TransfersController{baseController: baseController}
	return &transfersController
}

// TODO: type endpoint description here
func (t *TransfersController) GetTransferById(ctx context.Context, transferId string) (
	https.ApiResponse[models.GetTransfer],
	error) {
	req := t.prepareRequest(ctx, "GET", fmt.Sprintf("/transfers/%s", transferId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}

	var result models.GetTransfer
	result, err = utilities.DecodeResults[models.GetTransfer](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (t *TransfersController) CreateTransfer(ctx context.Context, request *models.CreateTransfer) (
	https.ApiResponse[models.GetTransfer],
	error) {
	req := t.prepareRequest(ctx, "POST", "/transfers/recipients")
	req.Authenticate(true)
	req.Json(request)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}

	var result models.GetTransfer
	result, err = utilities.DecodeResults[models.GetTransfer](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTransfer]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets all transfers
func (t *TransfersController) GetTransfers(ctx context.Context) (
	https.ApiResponse[models.ListTransfers],
	error) {
	req := t.prepareRequest(ctx, "GET", "/transfers")
	req.Authenticate(true)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListTransfers]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListTransfers]{Response: resp}, err
	}

	var result models.ListTransfers
	result, err = utilities.DecodeResults[models.ListTransfers](decoder)
	if err != nil {
		return https.ApiResponse[models.ListTransfers]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
