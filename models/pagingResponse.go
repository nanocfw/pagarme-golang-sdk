package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Object used for returning lists of objects with pagination
type PagingResponse struct {
	Total    types.Optional[int]    `json:"total"`
	Previous types.Optional[string] `json:"previous"`
	Next     types.Optional[string] `json:"next"`
}

func (p *PagingResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

func (p *PagingResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Total.IsValueSet() {
		structMap["total"] = p.Total.Value()
	}
	if p.Previous.IsValueSet() {
		structMap["previous"] = p.Previous.Value()
	}
	if p.Next.IsValueSet() {
		structMap["next"] = p.Next.Value()
	}
	return structMap
}

func (p *PagingResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Total    types.Optional[int]    `json:"total"`
		Previous types.Optional[string] `json:"previous"`
		Next     types.Optional[string] `json:"next"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Total = temp.Total
	p.Previous = temp.Previous
	p.Next = temp.Next
	return nil
}
