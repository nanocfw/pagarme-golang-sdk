package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetTransferSourceResponse struct {
	SourceId types.Optional[string] `json:"source_id"`
	Type     types.Optional[string] `json:"type"`
}

func (g *GetTransferSourceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTransferSourceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.SourceId.IsValueSet() {
		structMap["source_id"] = g.SourceId.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

func (g *GetTransferSourceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SourceId types.Optional[string] `json:"source_id"`
		Type     types.Optional[string] `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SourceId = temp.SourceId
	g.Type = temp.Type
	return nil
}
