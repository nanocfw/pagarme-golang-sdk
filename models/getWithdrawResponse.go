package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

type GetWithdrawResponse struct {
	Id                   types.Optional[string]                    `json:"id"`
	GatewayId            types.Optional[string]                    `json:"gateway_id"`
	Amount               types.Optional[int]                       `json:"amount"`
	Status               types.Optional[string]                    `json:"status"`
	CreatedAt            types.Optional[time.Time]                 `json:"created_at"`
	UpdatedAt            types.Optional[time.Time]                 `json:"updated_at"`
	Metadata             types.Optional[[]string]                  `json:"metadata"`
	Fee                  types.Optional[int]                       `json:"fee"`
	FundingDate          types.Optional[time.Time]                 `json:"funding_date"`
	FundingEstimatedDate types.Optional[time.Time]                 `json:"funding_estimated_date"`
	Type                 types.Optional[string]                    `json:"type"`
	Source               types.Optional[GetWithdrawSourceResponse] `json:"source"`
	Target               types.Optional[GetWithdrawTargetResponse] `json:"target"`
}

func (g *GetWithdrawResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetWithdrawResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
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
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	if g.FundingDate.IsValueSet() {
		var FundingDateVal *string = nil
		if g.FundingDate.Value() != nil {
			val := g.FundingDate.Value().Format(time.RFC3339)
			FundingDateVal = &val
		}
		structMap["funding_date"] = FundingDateVal
	}
	if g.FundingEstimatedDate.IsValueSet() {
		var FundingEstimatedDateVal *string = nil
		if g.FundingEstimatedDate.Value() != nil {
			val := g.FundingEstimatedDate.Value().Format(time.RFC3339)
			FundingEstimatedDateVal = &val
		}
		structMap["funding_estimated_date"] = FundingEstimatedDateVal
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Source.IsValueSet() {
		structMap["source"] = g.Source.Value()
	}
	if g.Target.IsValueSet() {
		structMap["target"] = g.Target.Value()
	}
	return structMap
}

func (g *GetWithdrawResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                   types.Optional[string]                    `json:"id"`
		GatewayId            types.Optional[string]                    `json:"gateway_id"`
		Amount               types.Optional[int]                       `json:"amount"`
		Status               types.Optional[string]                    `json:"status"`
		CreatedAt            types.Optional[string]                    `json:"created_at"`
		UpdatedAt            types.Optional[string]                    `json:"updated_at"`
		Metadata             types.Optional[[]string]                  `json:"metadata"`
		Fee                  types.Optional[int]                       `json:"fee"`
		FundingDate          types.Optional[string]                    `json:"funding_date"`
		FundingEstimatedDate types.Optional[string]                    `json:"funding_estimated_date"`
		Type                 types.Optional[string]                    `json:"type"`
		Source               types.Optional[GetWithdrawSourceResponse] `json:"source"`
		Target               types.Optional[GetWithdrawTargetResponse] `json:"target"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.GatewayId = temp.GatewayId
	g.Amount = temp.Amount
	g.Status = temp.Status
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
	g.Metadata = temp.Metadata
	g.Fee = temp.Fee
	g.FundingDate.ShouldSetValue(temp.FundingDate.IsValueSet())
	if temp.FundingDate.Value() != nil {
		FundingDateVal, err := time.Parse(time.RFC3339, (*temp.FundingDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse funding_date as % s format.", time.RFC3339)
		}
		g.FundingDate.SetValue(&FundingDateVal)
	}
	g.FundingEstimatedDate.ShouldSetValue(temp.FundingEstimatedDate.IsValueSet())
	if temp.FundingEstimatedDate.Value() != nil {
		FundingEstimatedDateVal, err := time.Parse(time.RFC3339, (*temp.FundingEstimatedDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse funding_estimated_date as % s format.", time.RFC3339)
		}
		g.FundingEstimatedDate.SetValue(&FundingEstimatedDateVal)
	}
	g.Type = temp.Type
	g.Source = temp.Source
	g.Target = temp.Target
	return nil
}
