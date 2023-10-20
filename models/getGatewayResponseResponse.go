package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// The Transaction Gateway Response
type GetGatewayResponseResponse struct {
	Code   types.Optional[string]                    `json:"code"`
	Errors types.Optional[[]GetGatewayErrorResponse] `json:"errors"`
}

func (g *GetGatewayResponseResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetGatewayResponseResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Errors.IsValueSet() {
		structMap["errors"] = g.Errors.Value()
	}
	return structMap
}

func (g *GetGatewayResponseResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code   types.Optional[string]                    `json:"code"`
		Errors types.Optional[[]GetGatewayErrorResponse] `json:"errors"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Code = temp.Code
	g.Errors = temp.Errors
	return nil
}
