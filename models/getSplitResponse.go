package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Split response
type GetSplitResponse struct {
	Type      types.Optional[string]                  `json:"type"`
	Amount    types.Optional[int]                     `json:"amount"`
	Recipient types.Optional[GetRecipientResponse]    `json:"recipient"`
	GatewayId types.Optional[string]                  `json:"gateway_id"`
	Options   types.Optional[GetSplitOptionsResponse] `json:"options"`
	Id        types.Optional[string]                  `json:"id"`
}

func (g *GetSplitResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSplitResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.Options.IsValueSet() {
		structMap["options"] = g.Options.Value()
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	return structMap
}

func (g *GetSplitResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      types.Optional[string]                  `json:"type"`
		Amount    types.Optional[int]                     `json:"amount"`
		Recipient types.Optional[GetRecipientResponse]    `json:"recipient"`
		GatewayId types.Optional[string]                  `json:"gateway_id"`
		Options   types.Optional[GetSplitOptionsResponse] `json:"options"`
		Id        types.Optional[string]                  `json:"id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Type = temp.Type
	g.Amount = temp.Amount
	g.Recipient = temp.Recipient
	g.GatewayId = temp.GatewayId
	g.Options = temp.Options
	g.Id = temp.Id
	return nil
}
