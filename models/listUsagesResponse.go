package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response model for listing the usages from a subscription item
type ListUsagesResponse struct {
	Data   types.Optional[[]GetUsageResponse] `json:"data"`
	Paging types.Optional[PagingResponse]     `json:"paging"`
}

func (l *ListUsagesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListUsagesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

func (l *ListUsagesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   types.Optional[[]GetUsageResponse] `json:"data"`
		Paging types.Optional[PagingResponse]     `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
