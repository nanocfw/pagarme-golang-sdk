package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting an payable
type GetPayableResponse struct {
	Id                       types.Optional[int64]     `json:"id"`
	Status                   types.Optional[string]    `json:"status"`
	Amount                   types.Optional[int]       `json:"amount"`
	Fee                      types.Optional[int]       `json:"fee"`
	AnticipationFee          types.Optional[int]       `json:"anticipation_fee"`
	FraudCoverageFee         types.Optional[int]       `json:"fraud_coverage_fee"`
	Installment              types.Optional[int]       `json:"installment"`
	GatewayId                types.Optional[int]       `json:"gateway_id"`
	ChargeId                 types.Optional[string]    `json:"charge_id"`
	SplitId                  types.Optional[string]    `json:"split_id"`
	BulkAnticipationId       types.Optional[string]    `json:"bulk_anticipation_id"`
	AnticipationId           types.Optional[string]    `json:"anticipation_id"`
	RecipientId              types.Optional[string]    `json:"recipient_id"`
	OriginatorModel          types.Optional[string]    `json:"originator_model"`
	OriginatorModelId        types.Optional[string]    `json:"originator_model_id"`
	PaymentDate              types.Optional[time.Time] `json:"payment_date"`
	OriginalPaymentDate      types.Optional[time.Time] `json:"original_payment_date"`
	Type                     types.Optional[string]    `json:"type"`
	PaymentMethod            types.Optional[string]    `json:"payment_method"`
	AccrualAt                types.Optional[time.Time] `json:"accrual_at"`
	CreatedAt                types.Optional[time.Time] `json:"created_at"`
	LiquidationArrangementId types.Optional[string]    `json:"liquidation_arrangement_id"`
}

func (g *GetPayableResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPayableResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	if g.AnticipationFee.IsValueSet() {
		structMap["anticipation_fee"] = g.AnticipationFee.Value()
	}
	if g.FraudCoverageFee.IsValueSet() {
		structMap["fraud_coverage_fee"] = g.FraudCoverageFee.Value()
	}
	if g.Installment.IsValueSet() {
		structMap["installment"] = g.Installment.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.ChargeId.IsValueSet() {
		structMap["charge_id"] = g.ChargeId.Value()
	}
	if g.SplitId.IsValueSet() {
		structMap["split_id"] = g.SplitId.Value()
	}
	if g.BulkAnticipationId.IsValueSet() {
		structMap["bulk_anticipation_id"] = g.BulkAnticipationId.Value()
	}
	if g.AnticipationId.IsValueSet() {
		structMap["anticipation_id"] = g.AnticipationId.Value()
	}
	if g.RecipientId.IsValueSet() {
		structMap["recipient_id"] = g.RecipientId.Value()
	}
	if g.OriginatorModel.IsValueSet() {
		structMap["originator_model"] = g.OriginatorModel.Value()
	}
	if g.OriginatorModelId.IsValueSet() {
		structMap["originator_model_id"] = g.OriginatorModelId.Value()
	}
	if g.PaymentDate.IsValueSet() {
		var PaymentDateVal *string = nil
		if g.PaymentDate.Value() != nil {
			val := g.PaymentDate.Value().Format(time.RFC3339)
			PaymentDateVal = &val
		}
		structMap["payment_date"] = PaymentDateVal
	}
	if g.OriginalPaymentDate.IsValueSet() {
		var OriginalPaymentDateVal *string = nil
		if g.OriginalPaymentDate.Value() != nil {
			val := g.OriginalPaymentDate.Value().Format(time.RFC3339)
			OriginalPaymentDateVal = &val
		}
		structMap["original_payment_date"] = OriginalPaymentDateVal
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.PaymentMethod.IsValueSet() {
		structMap["payment_method"] = g.PaymentMethod.Value()
	}
	if g.AccrualAt.IsValueSet() {
		var AccrualAtVal *string = nil
		if g.AccrualAt.Value() != nil {
			val := g.AccrualAt.Value().Format(time.RFC3339)
			AccrualAtVal = &val
		}
		structMap["accrual_at"] = AccrualAtVal
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.LiquidationArrangementId.IsValueSet() {
		structMap["liquidation_arrangement_id"] = g.LiquidationArrangementId.Value()
	}
	return structMap
}

func (g *GetPayableResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                       types.Optional[int64]  `json:"id"`
		Status                   types.Optional[string] `json:"status"`
		Amount                   types.Optional[int]    `json:"amount"`
		Fee                      types.Optional[int]    `json:"fee"`
		AnticipationFee          types.Optional[int]    `json:"anticipation_fee"`
		FraudCoverageFee         types.Optional[int]    `json:"fraud_coverage_fee"`
		Installment              types.Optional[int]    `json:"installment"`
		GatewayId                types.Optional[int]    `json:"gateway_id"`
		ChargeId                 types.Optional[string] `json:"charge_id"`
		SplitId                  types.Optional[string] `json:"split_id"`
		BulkAnticipationId       types.Optional[string] `json:"bulk_anticipation_id"`
		AnticipationId           types.Optional[string] `json:"anticipation_id"`
		RecipientId              types.Optional[string] `json:"recipient_id"`
		OriginatorModel          types.Optional[string] `json:"originator_model"`
		OriginatorModelId        types.Optional[string] `json:"originator_model_id"`
		PaymentDate              types.Optional[string] `json:"payment_date"`
		OriginalPaymentDate      types.Optional[string] `json:"original_payment_date"`
		Type                     types.Optional[string] `json:"type"`
		PaymentMethod            types.Optional[string] `json:"payment_method"`
		AccrualAt                types.Optional[string] `json:"accrual_at"`
		CreatedAt                types.Optional[string] `json:"created_at"`
		LiquidationArrangementId types.Optional[string] `json:"liquidation_arrangement_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Status = temp.Status
	g.Amount = temp.Amount
	g.Fee = temp.Fee
	g.AnticipationFee = temp.AnticipationFee
	g.FraudCoverageFee = temp.FraudCoverageFee
	g.Installment = temp.Installment
	g.GatewayId = temp.GatewayId
	g.ChargeId = temp.ChargeId
	g.SplitId = temp.SplitId
	g.BulkAnticipationId = temp.BulkAnticipationId
	g.AnticipationId = temp.AnticipationId
	g.RecipientId = temp.RecipientId
	g.OriginatorModel = temp.OriginatorModel
	g.OriginatorModelId = temp.OriginatorModelId
	g.PaymentDate.ShouldSetValue(temp.PaymentDate.IsValueSet())
	if temp.PaymentDate.Value() != nil {
		PaymentDateVal, err := time.Parse(time.RFC3339, (*temp.PaymentDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
		}
		g.PaymentDate.SetValue(&PaymentDateVal)
	}
	g.OriginalPaymentDate.ShouldSetValue(temp.OriginalPaymentDate.IsValueSet())
	if temp.OriginalPaymentDate.Value() != nil {
		OriginalPaymentDateVal, err := time.Parse(time.RFC3339, (*temp.OriginalPaymentDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse original_payment_date as % s format.", time.RFC3339)
		}
		g.OriginalPaymentDate.SetValue(&OriginalPaymentDateVal)
	}
	g.Type = temp.Type
	g.PaymentMethod = temp.PaymentMethod
	g.AccrualAt.ShouldSetValue(temp.AccrualAt.IsValueSet())
	if temp.AccrualAt.Value() != nil {
		AccrualAtVal, err := time.Parse(time.RFC3339, (*temp.AccrualAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse accrual_at as % s format.", time.RFC3339)
		}
		g.AccrualAt.SetValue(&AccrualAtVal)
	}
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.LiquidationArrangementId = temp.LiquidationArrangementId
	return nil
}
