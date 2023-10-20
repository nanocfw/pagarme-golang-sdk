package models

import (
	"encoding/json"
)

// Request for updating a plan item
type UpdatePlanItemRequest struct {
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	Status        string                     `json:"status"`
	PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
	Quantity      *int                       `json:"quantity,omitempty"`
	Cycles        *int                       `json:"cycles,omitempty"`
}

func (u *UpdatePlanItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdatePlanItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = u.Name
	structMap["description"] = u.Description
	structMap["status"] = u.Status
	structMap["pricing_scheme"] = u.PricingScheme
	if u.Quantity != nil {
		structMap["quantity"] = u.Quantity
	}
	if u.Cycles != nil {
		structMap["cycles"] = u.Cycles
	}
	return structMap
}

func (u *UpdatePlanItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name          string                     `json:"name"`
		Description   string                     `json:"description"`
		Status        string                     `json:"status"`
		PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
		Quantity      *int                       `json:"quantity,omitempty"`
		Cycles        *int                       `json:"cycles,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Description = temp.Description
	u.Status = temp.Status
	u.PricingScheme = temp.PricingScheme
	u.Quantity = temp.Quantity
	u.Cycles = temp.Cycles
	return nil
}
