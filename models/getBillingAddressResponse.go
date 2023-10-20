package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a billing address
type GetBillingAddressResponse struct {
	Street       types.Optional[string] `json:"street"`
	Number       types.Optional[string] `json:"number"`
	ZipCode      types.Optional[string] `json:"zip_code"`
	Neighborhood types.Optional[string] `json:"neighborhood"`
	City         types.Optional[string] `json:"city"`
	State        types.Optional[string] `json:"state"`
	Country      types.Optional[string] `json:"country"`
	Complement   types.Optional[string] `json:"complement"`
	Line1        types.Optional[string] `json:"line_1"`
	Line2        types.Optional[string] `json:"line_2"`
}

func (g *GetBillingAddressResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetBillingAddressResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Street.IsValueSet() {
		structMap["street"] = g.Street.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.ZipCode.IsValueSet() {
		structMap["zip_code"] = g.ZipCode.Value()
	}
	if g.Neighborhood.IsValueSet() {
		structMap["neighborhood"] = g.Neighborhood.Value()
	}
	if g.City.IsValueSet() {
		structMap["city"] = g.City.Value()
	}
	if g.State.IsValueSet() {
		structMap["state"] = g.State.Value()
	}
	if g.Country.IsValueSet() {
		structMap["country"] = g.Country.Value()
	}
	if g.Complement.IsValueSet() {
		structMap["complement"] = g.Complement.Value()
	}
	if g.Line1.IsValueSet() {
		structMap["line_1"] = g.Line1.Value()
	}
	if g.Line2.IsValueSet() {
		structMap["line_2"] = g.Line2.Value()
	}
	return structMap
}

func (g *GetBillingAddressResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Street       types.Optional[string] `json:"street"`
		Number       types.Optional[string] `json:"number"`
		ZipCode      types.Optional[string] `json:"zip_code"`
		Neighborhood types.Optional[string] `json:"neighborhood"`
		City         types.Optional[string] `json:"city"`
		State        types.Optional[string] `json:"state"`
		Country      types.Optional[string] `json:"country"`
		Complement   types.Optional[string] `json:"complement"`
		Line1        types.Optional[string] `json:"line_1"`
		Line2        types.Optional[string] `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Street = temp.Street
	g.Number = temp.Number
	g.ZipCode = temp.ZipCode
	g.Neighborhood = temp.Neighborhood
	g.City = temp.City
	g.State = temp.State
	g.Country = temp.Country
	g.Complement = temp.Complement
	g.Line1 = temp.Line1
	g.Line2 = temp.Line2
	return nil
}
