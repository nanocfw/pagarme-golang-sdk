package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

type GetTransactionReportFileResponse struct {
	Name types.Optional[string]    `json:"name"`
	Date types.Optional[time.Time] `json:"date"`
}

func (g *GetTransactionReportFileResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTransactionReportFileResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Date.IsValueSet() {
		var DateVal *string = nil
		if g.Date.Value() != nil {
			val := g.Date.Value().Format(time.RFC3339)
			DateVal = &val
		}
		structMap["date"] = DateVal
	}
	return structMap
}

func (g *GetTransactionReportFileResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name types.Optional[string] `json:"name"`
		Date types.Optional[string] `json:"date"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Name = temp.Name
	g.Date.ShouldSetValue(temp.Date.IsValueSet())
	if temp.Date.Value() != nil {
		DateVal, err := time.Parse(time.RFC3339, (*temp.Date.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse date as % s format.", time.RFC3339)
		}
		g.Date.SetValue(&DateVal)
	}
	return nil
}
