package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Generic response object for getting a MovementObjectRefund.
type GetMovementObjectRefundResponse struct {
	GetMovementObjectBaseResponse
	FraudCoverageFee     types.Optional[string] `json:"fraud_coverage_fee"`
	ChargeFeeRecipientId types.Optional[string] `json:"charge_fee_recipient_id"`
	BankAccountId        types.Optional[string] `json:"bank_account_id"`
	LocalTransactionId   types.Optional[string] `json:"local_transaction_id"`
	UpdatedAt            types.Optional[string] `json:"updated_at"`
}

func (g *GetMovementObjectRefundResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetMovementObjectRefundResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "refund"
	}
	if g.FraudCoverageFee.IsValueSet() {
		structMap["fraud_coverage_fee"] = g.FraudCoverageFee.Value()
	}
	if g.ChargeFeeRecipientId.IsValueSet() {
		structMap["charge_fee_recipient_id"] = g.ChargeFeeRecipientId.Value()
	}
	if g.BankAccountId.IsValueSet() {
		structMap["bank_account_id"] = g.BankAccountId.Value()
	}
	if g.LocalTransactionId.IsValueSet() {
		structMap["local_transaction_id"] = g.LocalTransactionId.Value()
	}
	if g.UpdatedAt.IsValueSet() {
		structMap["updated_at"] = g.UpdatedAt.Value()
	}
	return structMap
}

func (g *GetMovementObjectRefundResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		FraudCoverageFee     types.Optional[string] `json:"fraud_coverage_fee"`
		ChargeFeeRecipientId types.Optional[string] `json:"charge_fee_recipient_id"`
		BankAccountId        types.Optional[string] `json:"bank_account_id"`
		LocalTransactionId   types.Optional[string] `json:"local_transaction_id"`
		UpdatedAt            types.Optional[string] `json:"updated_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.FraudCoverageFee = temp.FraudCoverageFee
	g.ChargeFeeRecipientId = temp.ChargeFeeRecipientId
	g.BankAccountId = temp.BankAccountId
	g.LocalTransactionId = temp.LocalTransactionId
	g.UpdatedAt = temp.UpdatedAt
	return nil
}

// Getter for fraud_coverage_fee.
func (g *GetMovementObjectRefundResponse) GetFraudCoverageFee() types.Optional[string] {
	return g.FraudCoverageFee
}

// Getter for charge_fee_recipient_id.
func (g *GetMovementObjectRefundResponse) GetChargeFeeRecipientId() types.Optional[string] {
	return g.ChargeFeeRecipientId
}

// Getter for bank_account_id.
func (g *GetMovementObjectRefundResponse) GetBankAccountId() types.Optional[string] {
	return g.BankAccountId
}

// Getter for local_transaction_id.
func (g *GetMovementObjectRefundResponse) GetLocalTransactionId() types.Optional[string] {
	return g.LocalTransactionId
}

// Getter for updated_at.
func (g *GetMovementObjectRefundResponse) GetUpdatedAt() types.Optional[string] {
	return g.UpdatedAt
}

func UnmarshalGetMovementObjectRefundResponse(input []byte) (
	GetMovementObjectRefundResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectRefundResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectRefundResponseInterface); ok {
		return getMovementObjectRefundResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectRefundResponse %v", getMovementObjectBaseResponse)
}

func ToGetMovementObjectRefundResponseArray(getMovementObjectRefundResponse []GetMovementObjectRefundResponseField) []GetMovementObjectRefundResponseInterface {
	var items []GetMovementObjectRefundResponseInterface
	for _, temp := range getMovementObjectRefundResponse {
		items = append(items, temp.GetMovementObjectRefundResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetMovementObjectRefundResponseField struct {
	GetMovementObjectRefundResponseInterface
}

func (g *GetMovementObjectRefundResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectRefundResponseInterface, err = UnmarshalGetMovementObjectRefundResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetMovementObjectRefundResponseSlice struct {
	Slice []GetMovementObjectRefundResponseInterface
}

func (ga *GetMovementObjectRefundResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectRefundResponseInterface{}
		for _, getMovementObjectRefundResponse := range temp {
			if getMovementObjectRefundResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectRefundResponse
			err := json.Unmarshal(getMovementObjectRefundResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
