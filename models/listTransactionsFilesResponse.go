package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for listing of transactions files
type ListTransactionsFilesResponse struct {
	Data   types.Optional[[]GetTransactionReportFileResponse] `json:"data"`
	Paging types.Optional[PagingResponse]                     `json:"paging"`
}

func (l *ListTransactionsFilesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListTransactionsFilesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

func (l *ListTransactionsFilesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   types.Optional[[]GetTransactionReportFileResponse] `json:"data"`
		Paging types.Optional[PagingResponse]                     `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
