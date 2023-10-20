package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response for the listing recipient method
type ListRecipientResponse struct {
	Data   types.Optional[[]GetRecipientResponse] `json:"data"`
	Paging types.Optional[PagingResponse]         `json:"paging"`
}

func (l *ListRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

func (l *ListRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   types.Optional[[]GetRecipientResponse] `json:"data"`
		Paging types.Optional[PagingResponse]         `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
