package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for listing subscriptions
type ListSubscriptionsResponse struct {
	Data   types.Optional[[]GetSubscriptionResponse] `json:"data"`
	Paging types.Optional[PagingResponse]            `json:"paging"`
}

func (l *ListSubscriptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

func (l *ListSubscriptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

func (l *ListSubscriptionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   types.Optional[[]GetSubscriptionResponse] `json:"data"`
		Paging types.Optional[PagingResponse]            `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
