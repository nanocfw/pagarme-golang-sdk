package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// List of paginated transfer objects
type ListTransferResponse struct {
	Data   types.Optional[[]GetTransferResponse] `json:"data"`
	Paging types.Optional[PagingResponse]        `json:"paging"`
}

func (l *ListTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListTransferResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

func (l *ListTransferResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   types.Optional[[]GetTransferResponse] `json:"data"`
		Paging types.Optional[PagingResponse]        `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
