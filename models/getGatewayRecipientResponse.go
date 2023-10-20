package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Information about the recipient on the gateway
type GetGatewayRecipientResponse struct {
	Gateway   types.Optional[string] `json:"gateway"`
	Status    types.Optional[string] `json:"status"`
	Pgid      types.Optional[string] `json:"pgid"`
	CreatedAt types.Optional[string] `json:"created_at"`
	UpdatedAt types.Optional[string] `json:"updated_at"`
}

func (g *GetGatewayRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetGatewayRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Gateway.IsValueSet() {
		structMap["gateway"] = g.Gateway.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Pgid.IsValueSet() {
		structMap["pgid"] = g.Pgid.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.UpdatedAt.IsValueSet() {
		structMap["updated_at"] = g.UpdatedAt.Value()
	}
	return structMap
}

func (g *GetGatewayRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Gateway   types.Optional[string] `json:"gateway"`
		Status    types.Optional[string] `json:"status"`
		Pgid      types.Optional[string] `json:"pgid"`
		CreatedAt types.Optional[string] `json:"created_at"`
		UpdatedAt types.Optional[string] `json:"updated_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Gateway = temp.Gateway
	g.Status = temp.Status
	g.Pgid = temp.Pgid
	g.CreatedAt = temp.CreatedAt
	g.UpdatedAt = temp.UpdatedAt
	return nil
}
