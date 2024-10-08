package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Card token data
type GetCardTokenResponse struct {
	LastFourDigits types.Optional[string] `json:"last_four_digits"`
	HolderName     types.Optional[string] `json:"holder_name"`
	HolderDocument types.Optional[string] `json:"holder_document"`
	ExpMonth       types.Optional[int]    `json:"exp_month"`
	ExpYear        types.Optional[int]    `json:"exp_year"`
	Brand          types.Optional[string] `json:"brand"`
	Type           types.Optional[string] `json:"type"`
	Label          types.Optional[string] `json:"label"`
}

func (g *GetCardTokenResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCardTokenResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.LastFourDigits.IsValueSet() {
		structMap["last_four_digits"] = g.LastFourDigits.Value()
	}
	if g.HolderName.IsValueSet() {
		structMap["holder_name"] = g.HolderName.Value()
	}
	if g.HolderDocument.IsValueSet() {
		structMap["holder_document"] = g.HolderDocument.Value()
	}
	if g.ExpMonth.IsValueSet() {
		structMap["exp_month"] = g.ExpMonth.Value()
	}
	if g.ExpYear.IsValueSet() {
		structMap["exp_year"] = g.ExpYear.Value()
	}
	if g.Brand.IsValueSet() {
		structMap["brand"] = g.Brand.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Label.IsValueSet() {
		structMap["label"] = g.Label.Value()
	}
	return structMap
}

func (g *GetCardTokenResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		LastFourDigits types.Optional[string] `json:"last_four_digits"`
		HolderName     types.Optional[string] `json:"holder_name"`
		HolderDocument types.Optional[string] `json:"holder_document"`
		ExpMonth       types.Optional[int]    `json:"exp_month"`
		ExpYear        types.Optional[int]    `json:"exp_year"`
		Brand          types.Optional[string] `json:"brand"`
		Type           types.Optional[string] `json:"type"`
		Label          types.Optional[string] `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.LastFourDigits = temp.LastFourDigits
	g.HolderName = temp.HolderName
	g.HolderDocument = temp.HolderDocument
	g.ExpMonth = temp.ExpMonth
	g.ExpYear = temp.ExpYear
	g.Brand = temp.Brand
	g.Type = temp.Type
	g.Label = temp.Label
	return nil
}
