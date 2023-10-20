package models

import (
	"encoding/json"
)

type ListTransfers struct {
	Data   []GetTransfer  `json:"data"`
	Paging PagingResponse `json:"paging"`
}

func (l *ListTransfers) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListTransfers) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = l.Data
	structMap["paging"] = l.Paging
	return structMap
}

func (l *ListTransfers) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   []GetTransfer  `json:"data"`
		Paging PagingResponse `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
