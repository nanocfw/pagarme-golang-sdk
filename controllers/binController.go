package controllers

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/nanocfw/pagarme-golang-sdk/models"
)

type BinController struct {
	baseController
}

func NewBinController(baseController baseController) *BinController {
	binsController := BinController{baseController: baseController}
	return &binsController
}

// TODO: type endpoint description here
func (b *BinController) GetBin(ctx context.Context, bankIdentifier string) (
	https.ApiResponse[models.GetBinResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", fmt.Sprintf("/%s", bankIdentifier))
	req.BaseUrl("bin_v1")
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetBinResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetBinResponse]{Response: resp}, err
	}

	var result models.GetBinResponse
	result, err = utilities.DecodeResults[models.GetBinResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetBinResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
