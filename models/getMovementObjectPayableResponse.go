package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

type GetMovementObjectPayableResponse struct {
	GetMovementObjectBaseResponse
	Fee                      types.Optional[string] `json:"fee"`
	AnticipationFee          string                 `json:"anticipation_fee"`
	FraudCoverageFee         string                 `json:"fraud_coverage_fee"`
	Installment              string                 `json:"installment"`
	SplitId                  string                 `json:"split_id"`
	BulkAnticipationId       string                 `json:"bulk_anticipation_id"`
	AnticipationId           string                 `json:"anticipation_id"`
	RecipientId              string                 `json:"recipient_id"`
	OriginatorModel          string                 `json:"originator_model"`
	OriginatorModelId        string                 `json:"originator_model_id"`
	PaymentDate              string                 `json:"payment_date"`
	OriginalPaymentDate      string                 `json:"original_payment_date"`
	PaymentMethod            string                 `json:"payment_method"`
	AccrualAt                string                 `json:"accrual_at"`
	LiquidationArrangementId string                 `json:"liquidation_arrangement_id"`
}

func (g *GetMovementObjectPayableResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetMovementObjectPayableResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "payable"
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	structMap["anticipation_fee"] = g.AnticipationFee
	structMap["fraud_coverage_fee"] = g.FraudCoverageFee
	structMap["installment"] = g.Installment
	structMap["split_id"] = g.SplitId
	structMap["bulk_anticipation_id"] = g.BulkAnticipationId
	structMap["anticipation_id"] = g.AnticipationId
	structMap["recipient_id"] = g.RecipientId
	structMap["originator_model"] = g.OriginatorModel
	structMap["originator_model_id"] = g.OriginatorModelId
	structMap["payment_date"] = g.PaymentDate
	structMap["original_payment_date"] = g.OriginalPaymentDate
	structMap["payment_method"] = g.PaymentMethod
	structMap["accrual_at"] = g.AccrualAt
	structMap["liquidation_arrangement_id"] = g.LiquidationArrangementId
	return structMap
}

func (g *GetMovementObjectPayableResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		Fee                      types.Optional[string] `json:"fee"`
		AnticipationFee          string                 `json:"anticipation_fee"`
		FraudCoverageFee         string                 `json:"fraud_coverage_fee"`
		Installment              string                 `json:"installment"`
		SplitId                  string                 `json:"split_id"`
		BulkAnticipationId       string                 `json:"bulk_anticipation_id"`
		AnticipationId           string                 `json:"anticipation_id"`
		RecipientId              string                 `json:"recipient_id"`
		OriginatorModel          string                 `json:"originator_model"`
		OriginatorModelId        string                 `json:"originator_model_id"`
		PaymentDate              string                 `json:"payment_date"`
		OriginalPaymentDate      string                 `json:"original_payment_date"`
		PaymentMethod            string                 `json:"payment_method"`
		AccrualAt                string                 `json:"accrual_at"`
		LiquidationArrangementId string                 `json:"liquidation_arrangement_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Fee = temp.Fee
	g.AnticipationFee = temp.AnticipationFee
	g.FraudCoverageFee = temp.FraudCoverageFee
	g.Installment = temp.Installment
	g.SplitId = temp.SplitId
	g.BulkAnticipationId = temp.BulkAnticipationId
	g.AnticipationId = temp.AnticipationId
	g.RecipientId = temp.RecipientId
	g.OriginatorModel = temp.OriginatorModel
	g.OriginatorModelId = temp.OriginatorModelId
	g.PaymentDate = temp.PaymentDate
	g.OriginalPaymentDate = temp.OriginalPaymentDate
	g.PaymentMethod = temp.PaymentMethod
	g.AccrualAt = temp.AccrualAt
	g.LiquidationArrangementId = temp.LiquidationArrangementId
	return nil
}

// Getter for fee.
func (g *GetMovementObjectPayableResponse) GetFee() types.Optional[string] {
	return g.Fee
}

// Getter for anticipation_fee.
func (g *GetMovementObjectPayableResponse) GetAnticipationFee() string {
	return g.AnticipationFee
}

// Getter for fraud_coverage_fee.
func (g *GetMovementObjectPayableResponse) GetFraudCoverageFee() string {
	return g.FraudCoverageFee
}

// Getter for installment.
func (g *GetMovementObjectPayableResponse) GetInstallment() string {
	return g.Installment
}

// Getter for split_id.
func (g *GetMovementObjectPayableResponse) GetSplitId() string {
	return g.SplitId
}

// Getter for bulk_anticipation_id.
func (g *GetMovementObjectPayableResponse) GetBulkAnticipationId() string {
	return g.BulkAnticipationId
}

// Getter for anticipation_id.
func (g *GetMovementObjectPayableResponse) GetAnticipationId() string {
	return g.AnticipationId
}

// Getter for recipient_id.
func (g *GetMovementObjectPayableResponse) GetRecipientId() string {
	return g.RecipientId
}

// Getter for originator_model.
func (g *GetMovementObjectPayableResponse) GetOriginatorModel() string {
	return g.OriginatorModel
}

// Getter for originator_model_id.
func (g *GetMovementObjectPayableResponse) GetOriginatorModelId() string {
	return g.OriginatorModelId
}

// Getter for payment_date.
func (g *GetMovementObjectPayableResponse) GetPaymentDate() string {
	return g.PaymentDate
}

// Getter for original_payment_date.
func (g *GetMovementObjectPayableResponse) GetOriginalPaymentDate() string {
	return g.OriginalPaymentDate
}

// Getter for payment_method.
func (g *GetMovementObjectPayableResponse) GetPaymentMethod() string {
	return g.PaymentMethod
}

// Getter for accrual_at.
func (g *GetMovementObjectPayableResponse) GetAccrualAt() string {
	return g.AccrualAt
}

// Getter for liquidation_arrangement_id.
func (g *GetMovementObjectPayableResponse) GetLiquidationArrangementId() string {
	return g.LiquidationArrangementId
}

func UnmarshalGetMovementObjectPayableResponse(input []byte) (
	GetMovementObjectPayableResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectPayableResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectPayableResponseInterface); ok {
		return getMovementObjectPayableResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectPayableResponse %v", getMovementObjectBaseResponse)
}

func ToGetMovementObjectPayableResponseArray(getMovementObjectPayableResponse []GetMovementObjectPayableResponseField) []GetMovementObjectPayableResponseInterface {
	var items []GetMovementObjectPayableResponseInterface
	for _, temp := range getMovementObjectPayableResponse {
		items = append(items, temp.GetMovementObjectPayableResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetMovementObjectPayableResponseField struct {
	GetMovementObjectPayableResponseInterface
}

func (g *GetMovementObjectPayableResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectPayableResponseInterface, err = UnmarshalGetMovementObjectPayableResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetMovementObjectPayableResponseSlice struct {
	Slice []GetMovementObjectPayableResponseInterface
}

func (ga *GetMovementObjectPayableResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectPayableResponseInterface{}
		for _, getMovementObjectPayableResponse := range temp {
			if getMovementObjectPayableResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectPayableResponse
			err := json.Unmarshal(getMovementObjectPayableResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
