package models

import (
	"encoding/json"
)

// Request for updating a Subscription Affiliation Id
type UpdateSubscriptionAffiliationIdRequest struct {
	GatewayAffiliationId string `json:"gateway_affiliation_id"`
}

func (u *UpdateSubscriptionAffiliationIdRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionAffiliationIdRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["gateway_affiliation_id"] = u.GatewayAffiliationId
	return structMap
}

func (u *UpdateSubscriptionAffiliationIdRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		GatewayAffiliationId string `json:"gateway_affiliation_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.GatewayAffiliationId = temp.GatewayAffiliationId
	return nil
}
