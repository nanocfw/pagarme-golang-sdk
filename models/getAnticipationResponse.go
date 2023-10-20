package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Anticipation
type GetAnticipationResponse struct {
	Id              types.Optional[string]               `json:"id"`
	RequestedAmount types.Optional[int]                  `json:"requested_amount"`
	ApprovedAmount  types.Optional[int]                  `json:"approved_amount"`
	Recipient       types.Optional[GetRecipientResponse] `json:"recipient"`
	Pgid            types.Optional[string]               `json:"pgid"`
	CreatedAt       types.Optional[time.Time]            `json:"created_at"`
	UpdatedAt       types.Optional[time.Time]            `json:"updated_at"`
	PaymentDate     types.Optional[time.Time]            `json:"payment_date"`
	Status          types.Optional[string]               `json:"status"`
	Timeframe       types.Optional[string]               `json:"timeframe"`
}

func (g *GetAnticipationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAnticipationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.RequestedAmount.IsValueSet() {
		structMap["requested_amount"] = g.RequestedAmount.Value()
	}
	if g.ApprovedAmount.IsValueSet() {
		structMap["approved_amount"] = g.ApprovedAmount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.Pgid.IsValueSet() {
		structMap["pgid"] = g.Pgid.Value()
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
	if g.PaymentDate.IsValueSet() {
		var PaymentDateVal *string = nil
		if g.PaymentDate.Value() != nil {
			val := g.PaymentDate.Value().Format(time.RFC3339)
			PaymentDateVal = &val
		}
		structMap["payment_date"] = PaymentDateVal
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Timeframe.IsValueSet() {
		structMap["timeframe"] = g.Timeframe.Value()
	}
	return structMap
}

func (g *GetAnticipationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id              types.Optional[string]               `json:"id"`
		RequestedAmount types.Optional[int]                  `json:"requested_amount"`
		ApprovedAmount  types.Optional[int]                  `json:"approved_amount"`
		Recipient       types.Optional[GetRecipientResponse] `json:"recipient"`
		Pgid            types.Optional[string]               `json:"pgid"`
		CreatedAt       types.Optional[string]               `json:"created_at"`
		UpdatedAt       types.Optional[string]               `json:"updated_at"`
		PaymentDate     types.Optional[string]               `json:"payment_date"`
		Status          types.Optional[string]               `json:"status"`
		Timeframe       types.Optional[string]               `json:"timeframe"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.RequestedAmount = temp.RequestedAmount
	g.ApprovedAmount = temp.ApprovedAmount
	g.Recipient = temp.Recipient
	g.Pgid = temp.Pgid
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
	g.PaymentDate.ShouldSetValue(temp.PaymentDate.IsValueSet())
	if temp.PaymentDate.Value() != nil {
		PaymentDateVal, err := time.Parse(time.RFC3339, (*temp.PaymentDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
		}
		g.PaymentDate.SetValue(&PaymentDateVal)
	}
	g.Status = temp.Status
	g.Timeframe = temp.Timeframe
	return nil
}
