package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Gateway Response
type GetGatewayErrorResponse struct {
	Message types.Optional[string] `json:"message"`
}

func (g *GetGatewayErrorResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetGatewayErrorResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Message.IsValueSet() {
		structMap["message"] = g.Message.Value()
	}
	return structMap
}

func (g *GetGatewayErrorResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Message types.Optional[string] `json:"message"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Message = temp.Message
	return nil
}
