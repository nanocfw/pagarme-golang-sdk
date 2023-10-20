package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object when getting a pix transaction
type GetPixTransactionResponse struct {
	GetTransactionResponse
	QrCode                types.Optional[string]                     `json:"qr_code"`
	QrCodeUrl             types.Optional[string]                     `json:"qr_code_url"`
	ExpiresAt             types.Optional[time.Time]                  `json:"expires_at"`
	AdditionalInformation types.Optional[[]PixAdditionalInformation] `json:"additional_information"`
	EndToEndId            types.Optional[string]                     `json:"end_to_end_id"`
	Payer                 types.Optional[GetPixPayerResponse]        `json:"payer"`
	PixProviderTid        types.Optional[string]                     `json:"pix_provider_tid"`
}

func (g *GetPixTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPixTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "pix"
	}
	if g.QrCode.IsValueSet() {
		structMap["qr_code"] = g.QrCode.Value()
	}
	if g.QrCodeUrl.IsValueSet() {
		structMap["qr_code_url"] = g.QrCodeUrl.Value()
	}
	if g.ExpiresAt.IsValueSet() {
		var ExpiresAtVal *string = nil
		if g.ExpiresAt.Value() != nil {
			val := g.ExpiresAt.Value().Format(time.RFC3339)
			ExpiresAtVal = &val
		}
		structMap["expires_at"] = ExpiresAtVal
	}
	if g.AdditionalInformation.IsValueSet() {
		structMap["additional_information"] = g.AdditionalInformation.Value()
	}
	if g.EndToEndId.IsValueSet() {
		structMap["end_to_end_id"] = g.EndToEndId.Value()
	}
	if g.Payer.IsValueSet() {
		structMap["payer"] = g.Payer.Value()
	}
	if g.PixProviderTid.IsValueSet() {
		structMap["pix_provider_tid"] = g.PixProviderTid.Value()
	}
	return structMap
}

func (g *GetPixTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		QrCode                types.Optional[string]                     `json:"qr_code"`
		QrCodeUrl             types.Optional[string]                     `json:"qr_code_url"`
		ExpiresAt             types.Optional[string]                     `json:"expires_at"`
		AdditionalInformation types.Optional[[]PixAdditionalInformation] `json:"additional_information"`
		EndToEndId            types.Optional[string]                     `json:"end_to_end_id"`
		Payer                 types.Optional[GetPixPayerResponse]        `json:"payer"`
		PixProviderTid        types.Optional[string]                     `json:"pix_provider_tid"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.QrCode = temp.QrCode
	g.QrCodeUrl = temp.QrCodeUrl
	g.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
	if temp.ExpiresAt.Value() != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		g.ExpiresAt.SetValue(&ExpiresAtVal)
	}
	g.AdditionalInformation = temp.AdditionalInformation
	g.EndToEndId = temp.EndToEndId
	g.Payer = temp.Payer
	g.PixProviderTid = temp.PixProviderTid
	return nil
}

// Getter for qr_code.
func (g *GetPixTransactionResponse) GetQrCode() types.Optional[string] {
	return g.QrCode
}

// Getter for qr_code_url.
func (g *GetPixTransactionResponse) GetQrCodeUrl() types.Optional[string] {
	return g.QrCodeUrl
}

// Getter for expires_at.
func (g *GetPixTransactionResponse) GetExpiresAt() types.Optional[time.Time] {
	return g.ExpiresAt
}

// Getter for additional_information.
func (g *GetPixTransactionResponse) GetAdditionalInformation() types.Optional[[]PixAdditionalInformation] {
	return g.AdditionalInformation
}

// Getter for end_to_end_id.
func (g *GetPixTransactionResponse) GetEndToEndId() types.Optional[string] {
	return g.EndToEndId
}

// Getter for payer.
func (g *GetPixTransactionResponse) GetPayer() types.Optional[GetPixPayerResponse] {
	return g.Payer
}

// Getter for pix_provider_tid.
func (g *GetPixTransactionResponse) GetPixProviderTid() types.Optional[string] {
	return g.PixProviderTid
}

func UnmarshalGetPixTransactionResponse(input []byte) (
	GetPixTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getPixTransactionResponse, ok := getTransactionResponse.(GetPixTransactionResponseInterface); ok {
		return getPixTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getPixTransactionResponse %v", getTransactionResponse)
}

func ToGetPixTransactionResponseArray(getPixTransactionResponse []GetPixTransactionResponseField) []GetPixTransactionResponseInterface {
	var items []GetPixTransactionResponseInterface
	for _, temp := range getPixTransactionResponse {
		items = append(items, temp.GetPixTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetPixTransactionResponseField struct {
	GetPixTransactionResponseInterface
}

func (g *GetPixTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetPixTransactionResponseInterface, err = UnmarshalGetPixTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetPixTransactionResponseSlice struct {
	Slice []GetPixTransactionResponseInterface
}

func (ga *GetPixTransactionResponseSlice) UnmarshalJSON(input []byte) error {
	tempStruct := struct {
		Slice []json.RawMessage
	}{}
	temp := []json.RawMessage{}
	err := json.Unmarshal(input, &tempStruct)
	if err != nil {
		err := json.Unmarshal(input, &temp)
		if err != nil {
			return err
		}
	} else {
		temp = tempStruct.Slice
	}

	ga.Slice = nil
	if temp != nil {
		ga.Slice = []GetPixTransactionResponseInterface{}
		for _, getPixTransactionResponse := range temp {
			if getPixTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetPixTransactionResponse
			err := json.Unmarshal(getPixTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
