package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a credit card
type GetCardResponse struct {
	Id             types.Optional[string]                    `json:"id"`
	LastFourDigits types.Optional[string]                    `json:"last_four_digits"`
	Brand          types.Optional[string]                    `json:"brand"`
	HolderName     types.Optional[string]                    `json:"holder_name"`
	ExpMonth       types.Optional[int]                       `json:"exp_month"`
	ExpYear        types.Optional[int]                       `json:"exp_year"`
	Status         types.Optional[string]                    `json:"status"`
	CreatedAt      types.Optional[time.Time]                 `json:"created_at"`
	UpdatedAt      types.Optional[time.Time]                 `json:"updated_at"`
	BillingAddress types.Optional[GetBillingAddressResponse] `json:"billing_address"`
	Customer       types.Optional[GetCustomerResponse]       `json:"customer"`
	Metadata       types.Optional[map[string]string]         `json:"metadata"`
	Type           types.Optional[string]                    `json:"type"`
	HolderDocument types.Optional[string]                    `json:"holder_document"`
	DeletedAt      types.Optional[time.Time]                 `json:"deleted_at"`
	FirstSixDigits types.Optional[string]                    `json:"first_six_digits"`
	Label          types.Optional[string]                    `json:"label"`
}

func (g *GetCardResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCardResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.LastFourDigits.IsValueSet() {
		structMap["last_four_digits"] = g.LastFourDigits.Value()
	}
	if g.Brand.IsValueSet() {
		structMap["brand"] = g.Brand.Value()
	}
	if g.HolderName.IsValueSet() {
		structMap["holder_name"] = g.HolderName.Value()
	}
	if g.ExpMonth.IsValueSet() {
		structMap["exp_month"] = g.ExpMonth.Value()
	}
	if g.ExpYear.IsValueSet() {
		structMap["exp_year"] = g.ExpYear.Value()
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
	if g.BillingAddress.IsValueSet() {
		structMap["billing_address"] = g.BillingAddress.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.HolderDocument.IsValueSet() {
		structMap["holder_document"] = g.HolderDocument.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.FirstSixDigits.IsValueSet() {
		structMap["first_six_digits"] = g.FirstSixDigits.Value()
	}
	if g.Label.IsValueSet() {
		structMap["label"] = g.Label.Value()
	}
	return structMap
}

func (g *GetCardResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id             types.Optional[string]                    `json:"id"`
		LastFourDigits types.Optional[string]                    `json:"last_four_digits"`
		Brand          types.Optional[string]                    `json:"brand"`
		HolderName     types.Optional[string]                    `json:"holder_name"`
		ExpMonth       types.Optional[int]                       `json:"exp_month"`
		ExpYear        types.Optional[int]                       `json:"exp_year"`
		Status         types.Optional[string]                    `json:"status"`
		CreatedAt      types.Optional[string]                    `json:"created_at"`
		UpdatedAt      types.Optional[string]                    `json:"updated_at"`
		BillingAddress types.Optional[GetBillingAddressResponse] `json:"billing_address"`
		Customer       types.Optional[GetCustomerResponse]       `json:"customer"`
		Metadata       types.Optional[map[string]string]         `json:"metadata"`
		Type           types.Optional[string]                    `json:"type"`
		HolderDocument types.Optional[string]                    `json:"holder_document"`
		DeletedAt      types.Optional[string]                    `json:"deleted_at"`
		FirstSixDigits types.Optional[string]                    `json:"first_six_digits"`
		Label          types.Optional[string]                    `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.LastFourDigits = temp.LastFourDigits
	g.Brand = temp.Brand
	g.HolderName = temp.HolderName
	g.ExpMonth = temp.ExpMonth
	g.ExpYear = temp.ExpYear
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
	g.BillingAddress = temp.BillingAddress
	g.Customer = temp.Customer
	g.Metadata = temp.Metadata
	g.Type = temp.Type
	g.HolderDocument = temp.HolderDocument
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.FirstSixDigits = temp.FirstSixDigits
	g.Label = temp.Label
	return nil
}
