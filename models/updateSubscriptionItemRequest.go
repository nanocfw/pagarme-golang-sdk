package models

import (
	"encoding/json"
)

// Request for updating a subscription item
type UpdateSubscriptionItemRequest struct {
	Description   string                     `json:"description"`
	Status        string                     `json:"status"`
	PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
	Name          string                     `json:"name"`
	Cycles        *int                       `json:"cycles,omitempty"`
	Quantity      *int                       `json:"quantity,omitempty"`
	MinimumPrice  *int                       `json:"minimum_price,omitempty"`
}

func (u *UpdateSubscriptionItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = u.Description
	structMap["status"] = u.Status
	structMap["pricing_scheme"] = u.PricingScheme
	structMap["name"] = u.Name
	if u.Cycles != nil {
		structMap["cycles"] = u.Cycles
	}
	if u.Quantity != nil {
		structMap["quantity"] = u.Quantity
	}
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	return structMap
}

func (u *UpdateSubscriptionItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description   string                     `json:"description"`
		Status        string                     `json:"status"`
		PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
		Name          string                     `json:"name"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
		MinimumPrice  *int                       `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Description = temp.Description
	u.Status = temp.Status
	u.PricingScheme = temp.PricingScheme
	u.Name = temp.Name
	u.Cycles = temp.Cycles
	u.Quantity = temp.Quantity
	u.MinimumPrice = temp.MinimumPrice
	return nil
}
