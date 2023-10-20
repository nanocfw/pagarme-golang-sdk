package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting the setup from a subscription
type GetSetupResponse struct {
	Id          types.Optional[string] `json:"id"`
	Description types.Optional[string] `json:"description"`
	Amount      types.Optional[int]    `json:"amount"`
	Status      types.Optional[string] `json:"status"`
}

func (g *GetSetupResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSetupResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	return structMap
}

func (g *GetSetupResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          types.Optional[string] `json:"id"`
		Description types.Optional[string] `json:"description"`
		Amount      types.Optional[int]    `json:"amount"`
		Status      types.Optional[string] `json:"status"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Description = temp.Description
	g.Amount = temp.Amount
	g.Status = temp.Status
	return nil
}
