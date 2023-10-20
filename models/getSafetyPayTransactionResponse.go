package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponse struct {
	GetTransactionResponse
	Url        types.Optional[string]    `json:"url"`
	BankTid    types.Optional[string]    `json:"bank_tid"`
	PaidAt     types.Optional[time.Time] `json:"paid_at"`
	PaidAmount types.Optional[int]       `json:"paid_amount"`
}

func (g *GetSafetyPayTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSafetyPayTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "safetypay"
	}
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.BankTid.IsValueSet() {
		structMap["bank_tid"] = g.BankTid.Value()
	}
	if g.PaidAt.IsValueSet() {
		var PaidAtVal *string = nil
		if g.PaidAt.Value() != nil {
			val := g.PaidAt.Value().Format(time.RFC3339)
			PaidAtVal = &val
		}
		structMap["paid_at"] = PaidAtVal
	}
	if g.PaidAmount.IsValueSet() {
		structMap["paid_amount"] = g.PaidAmount.Value()
	}
	return structMap
}

func (g *GetSafetyPayTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Url        types.Optional[string] `json:"url"`
		BankTid    types.Optional[string] `json:"bank_tid"`
		PaidAt     types.Optional[string] `json:"paid_at"`
		PaidAmount types.Optional[int]    `json:"paid_amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.BankTid = temp.BankTid
	g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
	if temp.PaidAt.Value() != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt.SetValue(&PaidAtVal)
	}
	g.PaidAmount = temp.PaidAmount
	return nil
}

// Getter for url.
func (g *GetSafetyPayTransactionResponse) GetUrl() types.Optional[string] {
	return g.Url
}

// Getter for bank_tid.
func (g *GetSafetyPayTransactionResponse) GetBankTid() types.Optional[string] {
	return g.BankTid
}

// Getter for paid_at.
func (g *GetSafetyPayTransactionResponse) GetPaidAt() types.Optional[time.Time] {
	return g.PaidAt
}

// Getter for paid_amount.
func (g *GetSafetyPayTransactionResponse) GetPaidAmount() types.Optional[int] {
	return g.PaidAmount
}

func UnmarshalGetSafetyPayTransactionResponse(input []byte) (
	GetSafetyPayTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getSafetyPayTransactionResponse, ok := getTransactionResponse.(GetSafetyPayTransactionResponseInterface); ok {
		return getSafetyPayTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getSafetyPayTransactionResponse %v", getTransactionResponse)
}

func ToGetSafetyPayTransactionResponseArray(getSafetyPayTransactionResponse []GetSafetyPayTransactionResponseField) []GetSafetyPayTransactionResponseInterface {
	var items []GetSafetyPayTransactionResponseInterface
	for _, temp := range getSafetyPayTransactionResponse {
		items = append(items, temp.GetSafetyPayTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetSafetyPayTransactionResponseField struct {
	GetSafetyPayTransactionResponseInterface
}

func (g *GetSafetyPayTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetSafetyPayTransactionResponseInterface, err = UnmarshalGetSafetyPayTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetSafetyPayTransactionResponseSlice struct {
	Slice []GetSafetyPayTransactionResponseInterface
}

func (ga *GetSafetyPayTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetSafetyPayTransactionResponseInterface{}
		for _, getSafetyPayTransactionResponse := range temp {
			if getSafetyPayTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetSafetyPayTransactionResponse
			err := json.Unmarshal(getSafetyPayTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
