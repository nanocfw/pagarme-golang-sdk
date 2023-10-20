package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

type GetSubscriptionItemResponse struct {
	Id            types.Optional[string]                   `json:"id"`
	Description   types.Optional[string]                   `json:"description"`
	Status        types.Optional[string]                   `json:"status"`
	CreatedAt     types.Optional[time.Time]                `json:"created_at"`
	UpdatedAt     types.Optional[time.Time]                `json:"updated_at"`
	PricingScheme types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
	Discounts     types.Optional[[]GetDiscountResponse]    `json:"discounts"`
	Increments    types.Optional[[]GetIncrementResponse]   `json:"increments"`
	Subscription  types.Optional[GetSubscriptionResponse]  `json:"subscription"`
	Name          types.Optional[string]                   `json:"name"`
	Quantity      types.Optional[int]                      `json:"quantity"`
	Cycles        types.Optional[int]                      `json:"cycles"`
	DeletedAt     types.Optional[time.Time]                `json:"deleted_at"`
}

func (g *GetSubscriptionItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSubscriptionItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
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
	if g.PricingScheme.IsValueSet() {
		structMap["pricing_scheme"] = g.PricingScheme.Value()
	}
	if g.Discounts.IsValueSet() {
		structMap["discounts"] = g.Discounts.Value()
	}
	if g.Increments.IsValueSet() {
		structMap["increments"] = g.Increments.Value()
	}
	if g.Subscription.IsValueSet() {
		structMap["subscription"] = g.Subscription.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
	}
	if g.Cycles.IsValueSet() {
		structMap["cycles"] = g.Cycles.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	return structMap
}

func (g *GetSubscriptionItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            types.Optional[string]                   `json:"id"`
		Description   types.Optional[string]                   `json:"description"`
		Status        types.Optional[string]                   `json:"status"`
		CreatedAt     types.Optional[string]                   `json:"created_at"`
		UpdatedAt     types.Optional[string]                   `json:"updated_at"`
		PricingScheme types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
		Discounts     types.Optional[[]GetDiscountResponse]    `json:"discounts"`
		Increments    types.Optional[[]GetIncrementResponse]   `json:"increments"`
		Subscription  types.Optional[GetSubscriptionResponse]  `json:"subscription"`
		Name          types.Optional[string]                   `json:"name"`
		Quantity      types.Optional[int]                      `json:"quantity"`
		Cycles        types.Optional[int]                      `json:"cycles"`
		DeletedAt     types.Optional[string]                   `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Description = temp.Description
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
	g.PricingScheme = temp.PricingScheme
	g.Discounts = temp.Discounts
	g.Increments = temp.Increments
	g.Subscription = temp.Subscription
	g.Name = temp.Name
	g.Quantity = temp.Quantity
	g.Cycles = temp.Cycles
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	return nil
}
