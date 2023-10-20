package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Token data
type GetTokenResponse struct {
	Id        types.Optional[string]               `json:"id"`
	Type      types.Optional[string]               `json:"type"`
	CreatedAt types.Optional[time.Time]            `json:"created_at"`
	ExpiresAt types.Optional[string]               `json:"expires_at"`
	Card      types.Optional[GetCardTokenResponse] `json:"card"`
}

func (g *GetTokenResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTokenResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.ExpiresAt.IsValueSet() {
		structMap["expires_at"] = g.ExpiresAt.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	return structMap
}

func (g *GetTokenResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id        types.Optional[string]               `json:"id"`
		Type      types.Optional[string]               `json:"type"`
		CreatedAt types.Optional[string]               `json:"created_at"`
		ExpiresAt types.Optional[string]               `json:"expires_at"`
		Card      types.Optional[GetCardTokenResponse] `json:"card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Type = temp.Type
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.ExpiresAt = temp.ExpiresAt
	g.Card = temp.Card
	return nil
}
