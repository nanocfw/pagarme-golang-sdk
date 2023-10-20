package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Generic response object for getting a transaction.
type GetTransactionResponse struct {
	GatewayId           types.Optional[string]                     `json:"gateway_id"`
	Amount              types.Optional[int]                        `json:"amount"`
	Status              types.Optional[string]                     `json:"status"`
	Success             types.Optional[bool]                       `json:"success"`
	CreatedAt           types.Optional[time.Time]                  `json:"created_at"`
	UpdatedAt           types.Optional[time.Time]                  `json:"updated_at"`
	AttemptCount        types.Optional[int]                        `json:"attempt_count"`
	MaxAttempts         types.Optional[int]                        `json:"max_attempts"`
	Splits              types.Optional[[]GetSplitResponse]         `json:"splits"`
	NextAttempt         types.Optional[time.Time]                  `json:"next_attempt"`
	TransactionType     *string                                    `json:"transaction_type,omitempty"`
	Id                  types.Optional[string]                     `json:"id"`
	GatewayResponse     types.Optional[GetGatewayResponseResponse] `json:"gateway_response"`
	AntifraudResponse   types.Optional[GetAntifraudResponse]       `json:"antifraud_response"`
	Metadata            types.Optional[map[string]string]          `json:"metadata"`
	Split               types.Optional[[]GetSplitResponse]         `json:"split"`
	Interest            types.Optional[GetInterestResponse]        `json:"interest"`
	Fine                types.Optional[GetFineResponse]            `json:"fine"`
	MaxDaysToPayPastDue types.Optional[int]                        `json:"max_days_to_pay_past_due"`
}

func (g *GetTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTransactionResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "transaction"
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Success.IsValueSet() {
		structMap["success"] = g.Success.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.AttemptCount.IsValueSet() {
		structMap["attempt_count"] = g.AttemptCount.Value()
	}
	if g.MaxAttempts.IsValueSet() {
		structMap["max_attempts"] = g.MaxAttempts.Value()
	}
	if g.Splits.IsValueSet() {
		structMap["splits"] = g.Splits.Value()
	}
	if g.NextAttempt.IsValueSet() {
		var NextAttemptVal *string = nil
		if g.NextAttempt.Value() != nil {
			val := g.NextAttempt.Value().Format(time.RFC3339)
			NextAttemptVal = &val
		}
		structMap["next_attempt"] = NextAttemptVal
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.GatewayResponse.IsValueSet() {
		structMap["gateway_response"] = g.GatewayResponse.Value()
	}
	if g.AntifraudResponse.IsValueSet() {
		structMap["antifraud_response"] = g.AntifraudResponse.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Split.IsValueSet() {
		structMap["split"] = g.Split.Value()
	}
	if g.Interest.IsValueSet() {
		structMap["interest"] = g.Interest.Value()
	}
	if g.Fine.IsValueSet() {
		structMap["fine"] = g.Fine.Value()
	}
	if g.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = g.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

func (g *GetTransactionResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		GatewayId           types.Optional[string]                     `json:"gateway_id"`
		Amount              types.Optional[int]                        `json:"amount"`
		Status              types.Optional[string]                     `json:"status"`
		Success             types.Optional[bool]                       `json:"success"`
		CreatedAt           types.Optional[string]                     `json:"created_at"`
		UpdatedAt           types.Optional[string]                     `json:"updated_at"`
		AttemptCount        types.Optional[int]                        `json:"attempt_count"`
		MaxAttempts         types.Optional[int]                        `json:"max_attempts"`
		Splits              types.Optional[[]GetSplitResponse]         `json:"splits"`
		NextAttempt         types.Optional[string]                     `json:"next_attempt"`
		TransactionType     *string                                    `json:"transaction_type,omitempty"`
		Id                  types.Optional[string]                     `json:"id"`
		GatewayResponse     types.Optional[GetGatewayResponseResponse] `json:"gateway_response"`
		AntifraudResponse   types.Optional[GetAntifraudResponse]       `json:"antifraud_response"`
		Metadata            types.Optional[map[string]string]          `json:"metadata"`
		Split               types.Optional[[]GetSplitResponse]         `json:"split"`
		Interest            types.Optional[GetInterestResponse]        `json:"interest"`
		Fine                types.Optional[GetFineResponse]            `json:"fine"`
		MaxDaysToPayPastDue types.Optional[int]                        `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.GatewayId = temp.GatewayId
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.Success = temp.Success
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.AttemptCount = temp.AttemptCount
	g.MaxAttempts = temp.MaxAttempts
	g.Splits = temp.Splits
	g.NextAttempt.ShouldSetValue(temp.NextAttempt.IsValueSet())
	if temp.NextAttempt.Value() != nil {
		NextAttemptVal, err := time.Parse(time.RFC3339, (*temp.NextAttempt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse next_attempt as % s format.", time.RFC3339)
		}
		g.NextAttempt.SetValue(&NextAttemptVal)
	}
	g.TransactionType = temp.TransactionType
	g.Id = temp.Id
	g.GatewayResponse = temp.GatewayResponse
	g.AntifraudResponse = temp.AntifraudResponse
	g.Metadata = temp.Metadata
	g.Split = temp.Split
	g.Interest = temp.Interest
	g.Fine = temp.Fine
	g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}

// Getter for gateway_id.
func (g *GetTransactionResponse) GetGatewayId() types.Optional[string] {
	return g.GatewayId
}

// Getter for amount.
func (g *GetTransactionResponse) GetAmount() types.Optional[int] {
	return g.Amount
}

// Getter for status.
func (g *GetTransactionResponse) GetStatus() types.Optional[string] {
	return g.Status
}

// Getter for success.
func (g *GetTransactionResponse) GetSuccess() types.Optional[bool] {
	return g.Success
}

// Getter for created_at.
func (g *GetTransactionResponse) GetCreatedAt() types.Optional[time.Time] {
	return g.CreatedAt
}

// Getter for updated_at.
func (g *GetTransactionResponse) GetUpdatedAt() types.Optional[time.Time] {
	return g.UpdatedAt
}

// Getter for attempt_count.
func (g *GetTransactionResponse) GetAttemptCount() types.Optional[int] {
	return g.AttemptCount
}

// Getter for max_attempts.
func (g *GetTransactionResponse) GetMaxAttempts() types.Optional[int] {
	return g.MaxAttempts
}

// Getter for splits.
func (g *GetTransactionResponse) GetSplits() types.Optional[[]GetSplitResponse] {
	return g.Splits
}

// Getter for next_attempt.
func (g *GetTransactionResponse) GetNextAttempt() types.Optional[time.Time] {
	return g.NextAttempt
}

// Getter for transaction_type.
func (g *GetTransactionResponse) GetTransactionType() *string {
	return g.TransactionType
}

// Getter for id.
func (g *GetTransactionResponse) GetId() types.Optional[string] {
	return g.Id
}

// Getter for gateway_response.
func (g *GetTransactionResponse) GetGatewayResponse() types.Optional[GetGatewayResponseResponse] {
	return g.GatewayResponse
}

// Getter for antifraud_response.
func (g *GetTransactionResponse) GetAntifraudResponse() types.Optional[GetAntifraudResponse] {
	return g.AntifraudResponse
}

// Getter for metadata.
func (g *GetTransactionResponse) GetMetadata() types.Optional[map[string]string] {
	return g.Metadata
}

// Getter for split.
func (g *GetTransactionResponse) GetSplit() types.Optional[[]GetSplitResponse] {
	return g.Split
}

// Getter for interest.
func (g *GetTransactionResponse) GetInterest() types.Optional[GetInterestResponse] {
	return g.Interest
}

// Getter for fine.
func (g *GetTransactionResponse) GetFine() types.Optional[GetFineResponse] {
	return g.Fine
}

// Getter for max_days_to_pay_past_due.
func (g *GetTransactionResponse) GetMaxDaysToPayPastDue() types.Optional[int] {
	return g.MaxDaysToPayPastDue
}

func UnmarshalGetTransactionResponse(input []byte) (
	GetTransactionResponseInterface,
	error) {
	discrim := &struct {
		TransactionType *string
	}{}

	err := json.Unmarshal(input, &discrim)
	if err != nil {
		return nil, err
	}
	if discrim.TransactionType == nil {
		transactionType := "transaction"
		discrim.TransactionType = &transactionType
	}

	var res GetTransactionResponseInterface
	switch *discrim.TransactionType {
	case "bank_transfer":
		res = &GetBankTransferTransactionResponse{}
	case "safetypay":
		res = &GetSafetyPayTransactionResponse{}
	case "voucher":
		res = &GetVoucherTransactionResponse{}
	case "boleto":
		res = &GetBoletoTransactionResponse{}
	case "debit_card":
		res = &GetDebitCardTransactionResponse{}
	case "private_label":
		res = &GetPrivateLabelTransactionResponse{}
	case "cash":
		res = &GetCashTransactionResponse{}
	case "credit_card":
		res = &GetCreditCardTransactionResponse{}
	case "pix":
		res = &GetPixTransactionResponse{}
	default:
		res = &GetTransactionResponse{}
	}
	json.Unmarshal(input, res)
	return res, nil
}

func ToGetTransactionResponseArray(getTransactionResponse []GetTransactionResponseField) []GetTransactionResponseInterface {
	var items []GetTransactionResponseInterface
	for _, temp := range getTransactionResponse {
		items = append(items, temp.GetTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetTransactionResponseField struct {
	GetTransactionResponseInterface
}

func (g *GetTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetTransactionResponseInterface, err = UnmarshalGetTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetTransactionResponseSlice struct {
	Slice []GetTransactionResponseInterface
}

func (ga *GetTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetTransactionResponseInterface{}
		for _, getTransactionResponse := range temp {
			if getTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetTransactionResponse
			err := json.Unmarshal(getTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
