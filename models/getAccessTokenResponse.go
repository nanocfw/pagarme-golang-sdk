package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a access token
type GetAccessTokenResponse struct {
	Id        types.Optional[string]              `json:"id"`
	Code      types.Optional[string]              `json:"code"`
	Status    types.Optional[string]              `json:"status"`
	CreatedAt types.Optional[time.Time]           `json:"created_at"`
	Customer  types.Optional[GetCustomerResponse] `json:"customer"`
}

func (g *GetAccessTokenResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAccessTokenResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	return structMap
}

func (g *GetAccessTokenResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id        types.Optional[string]              `json:"id"`
		Code      types.Optional[string]              `json:"code"`
		Status    types.Optional[string]              `json:"status"`
		CreatedAt types.Optional[string]              `json:"created_at"`
		Customer  types.Optional[GetCustomerResponse] `json:"customer"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.Customer = temp.Customer
	return nil
}
