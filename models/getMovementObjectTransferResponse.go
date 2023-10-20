package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

type GetMovementObjectTransferResponse struct {
	GetMovementObjectBaseResponse
	SourceType           types.Optional[string] `json:"source_type"`
	SourceId             types.Optional[string] `json:"source_id"`
	TargetType           types.Optional[string] `json:"target_type"`
	TargetId             types.Optional[string] `json:"target_id"`
	Fee                  types.Optional[string] `json:"fee"`
	FundingDate          types.Optional[string] `json:"funding_date"`
	FundingEstimatedDate types.Optional[string] `json:"funding_estimated_date"`
	BankAccount          types.Optional[string] `json:"bank_account"`
}

func (g *GetMovementObjectTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetMovementObjectTransferResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "transfer"
	}
	if g.SourceType.IsValueSet() {
		structMap["source_type"] = g.SourceType.Value()
	}
	if g.SourceId.IsValueSet() {
		structMap["source_id"] = g.SourceId.Value()
	}
	if g.TargetType.IsValueSet() {
		structMap["target_type"] = g.TargetType.Value()
	}
	if g.TargetId.IsValueSet() {
		structMap["target_id"] = g.TargetId.Value()
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	if g.FundingDate.IsValueSet() {
		structMap["funding_date"] = g.FundingDate.Value()
	}
	if g.FundingEstimatedDate.IsValueSet() {
		structMap["funding_estimated_date"] = g.FundingEstimatedDate.Value()
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	return structMap
}

func (g *GetMovementObjectTransferResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		SourceType           types.Optional[string] `json:"source_type"`
		SourceId             types.Optional[string] `json:"source_id"`
		TargetType           types.Optional[string] `json:"target_type"`
		TargetId             types.Optional[string] `json:"target_id"`
		Fee                  types.Optional[string] `json:"fee"`
		FundingDate          types.Optional[string] `json:"funding_date"`
		FundingEstimatedDate types.Optional[string] `json:"funding_estimated_date"`
		BankAccount          types.Optional[string] `json:"bank_account"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SourceType = temp.SourceType
	g.SourceId = temp.SourceId
	g.TargetType = temp.TargetType
	g.TargetId = temp.TargetId
	g.Fee = temp.Fee
	g.FundingDate = temp.FundingDate
	g.FundingEstimatedDate = temp.FundingEstimatedDate
	g.BankAccount = temp.BankAccount
	return nil
}

// Getter for source_type.
func (g *GetMovementObjectTransferResponse) GetSourceType() types.Optional[string] {
	return g.SourceType
}

// Getter for source_id.
func (g *GetMovementObjectTransferResponse) GetSourceId() types.Optional[string] {
	return g.SourceId
}

// Getter for target_type.
func (g *GetMovementObjectTransferResponse) GetTargetType() types.Optional[string] {
	return g.TargetType
}

// Getter for target_id.
func (g *GetMovementObjectTransferResponse) GetTargetId() types.Optional[string] {
	return g.TargetId
}

// Getter for fee.
func (g *GetMovementObjectTransferResponse) GetFee() types.Optional[string] {
	return g.Fee
}

// Getter for funding_date.
func (g *GetMovementObjectTransferResponse) GetFundingDate() types.Optional[string] {
	return g.FundingDate
}

// Getter for funding_estimated_date.
func (g *GetMovementObjectTransferResponse) GetFundingEstimatedDate() types.Optional[string] {
	return g.FundingEstimatedDate
}

// Getter for bank_account.
func (g *GetMovementObjectTransferResponse) GetBankAccount() types.Optional[string] {
	return g.BankAccount
}

func UnmarshalGetMovementObjectTransferResponse(input []byte) (
	GetMovementObjectTransferResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectTransferResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectTransferResponseInterface); ok {
		return getMovementObjectTransferResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectTransferResponse %v", getMovementObjectBaseResponse)
}

func ToGetMovementObjectTransferResponseArray(getMovementObjectTransferResponse []GetMovementObjectTransferResponseField) []GetMovementObjectTransferResponseInterface {
	var items []GetMovementObjectTransferResponseInterface
	for _, temp := range getMovementObjectTransferResponse {
		items = append(items, temp.GetMovementObjectTransferResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetMovementObjectTransferResponseField struct {
	GetMovementObjectTransferResponseInterface
}

func (g *GetMovementObjectTransferResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectTransferResponseInterface, err = UnmarshalGetMovementObjectTransferResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetMovementObjectTransferResponseSlice struct {
	Slice []GetMovementObjectTransferResponseInterface
}

func (ga *GetMovementObjectTransferResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectTransferResponseInterface{}
		for _, getMovementObjectTransferResponse := range temp {
			if getMovementObjectTransferResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectTransferResponse
			err := json.Unmarshal(getMovementObjectTransferResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
