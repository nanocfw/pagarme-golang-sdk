package models

import (
	"encoding/json"
)

type ListWithdrawals struct {
	Data   []GetWithdrawResponse `json:"data"`
	Paging PagingResponse        `json:"paging"`
}

func (l *ListWithdrawals) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListWithdrawals) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = l.Data
	structMap["paging"] = l.Paging
	return structMap
}

func (l *ListWithdrawals) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   []GetWithdrawResponse `json:"data"`
		Paging PagingResponse        `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
