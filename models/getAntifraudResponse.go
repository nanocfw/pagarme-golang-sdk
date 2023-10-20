package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetAntifraudResponse struct {
	Status        types.Optional[string] `json:"status"`
	ReturnCode    types.Optional[string] `json:"return_code"`
	ReturnMessage types.Optional[string] `json:"return_message"`
	ProviderName  types.Optional[string] `json:"provider_name"`
	Score         types.Optional[string] `json:"score"`
}

func (g *GetAntifraudResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAntifraudResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.ReturnCode.IsValueSet() {
		structMap["return_code"] = g.ReturnCode.Value()
	}
	if g.ReturnMessage.IsValueSet() {
		structMap["return_message"] = g.ReturnMessage.Value()
	}
	if g.ProviderName.IsValueSet() {
		structMap["provider_name"] = g.ProviderName.Value()
	}
	if g.Score.IsValueSet() {
		structMap["score"] = g.Score.Value()
	}
	return structMap
}

func (g *GetAntifraudResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Status        types.Optional[string] `json:"status"`
		ReturnCode    types.Optional[string] `json:"return_code"`
		ReturnMessage types.Optional[string] `json:"return_message"`
		ProviderName  types.Optional[string] `json:"provider_name"`
		Score         types.Optional[string] `json:"score"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Status = temp.Status
	g.ReturnCode = temp.ReturnCode
	g.ReturnMessage = temp.ReturnMessage
	g.ProviderName = temp.ProviderName
	g.Score = temp.Score
	return nil
}
