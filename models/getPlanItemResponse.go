package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a plan item
type GetPlanItemResponse struct {
	Id            types.Optional[string]                   `json:"id"`
	Name          types.Optional[string]                   `json:"name"`
	Status        types.Optional[string]                   `json:"status"`
	CreatedAt     types.Optional[time.Time]                `json:"created_at"`
	UpdatedAt     types.Optional[time.Time]                `json:"updated_at"`
	PricingScheme types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
	Description   types.Optional[string]                   `json:"description"`
	Plan          types.Optional[GetPlanResponse]          `json:"plan"`
	Quantity      types.Optional[int]                      `json:"quantity"`
	Cycles        types.Optional[int]                      `json:"cycles"`
	DeletedAt     types.Optional[time.Time]                `json:"deleted_at"`
}

func (g *GetPlanItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPlanItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
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
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Plan.IsValueSet() {
		structMap["plan"] = g.Plan.Value()
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

func (g *GetPlanItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            types.Optional[string]                   `json:"id"`
		Name          types.Optional[string]                   `json:"name"`
		Status        types.Optional[string]                   `json:"status"`
		CreatedAt     types.Optional[string]                   `json:"created_at"`
		UpdatedAt     types.Optional[string]                   `json:"updated_at"`
		PricingScheme types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
		Description   types.Optional[string]                   `json:"description"`
		Plan          types.Optional[GetPlanResponse]          `json:"plan"`
		Quantity      types.Optional[int]                      `json:"quantity"`
		Cycles        types.Optional[int]                      `json:"cycles"`
		DeletedAt     types.Optional[string]                   `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Name = temp.Name
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
	g.Description = temp.Description
	g.Plan = temp.Plan
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
