package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Payment Authentication response
type GetPaymentAuthenticationResponse struct {
	Type         types.Optional[string]                  `json:"type"`
	ThreedSecure types.Optional[GetThreeDSecureResponse] `json:"threed_secure"`
}

func (g *GetPaymentAuthenticationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPaymentAuthenticationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.ThreedSecure.IsValueSet() {
		structMap["threed_secure"] = g.ThreedSecure.Value()
	}
	return structMap
}

func (g *GetPaymentAuthenticationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type         types.Optional[string]                  `json:"type"`
		ThreedSecure types.Optional[GetThreeDSecureResponse] `json:"threed_secure"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Type = temp.Type
	g.ThreedSecure = temp.ThreedSecure
	return nil
}
