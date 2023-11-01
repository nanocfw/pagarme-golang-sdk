package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a bin info
type GetBinResponse struct {
	Brand     types.Optional[string] `json:"brand"`
	BrandName types.Optional[string] `json:"brandName"`
	Gaps      types.Optional[[]int]  `json:"gaps"`
	Lenghts   types.Optional[[]int]  `json:"lenghts"`
	Mask      types.Optional[string] `json:"mask"`
	Input     types.Optional[string] `json:"input"`
	Cvv       types.Optional[int]    `json:"cvv"`
}

func (g *GetBinResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetBinResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Brand.IsValueSet() {
		structMap["brand"] = g.Brand.Value()
	}
	if g.BrandName.IsValueSet() {
		structMap["brandName"] = g.BrandName.Value()
	}
	if g.Gaps.IsValueSet() {
		structMap["gaps"] = g.Gaps.Value()
	}
	if g.Lenghts.IsValueSet() {
		structMap["lenghts"] = g.Lenghts.Value()
	}
	if g.Mask.IsValueSet() {
		structMap["mask"] = g.Mask.Value()
	}
	if g.Input.IsValueSet() {
		structMap["input"] = g.Input.Value()
	}
	if g.Cvv.IsValueSet() {
		structMap["cvv"] = g.Cvv.Value()
	}

	return structMap
}

func (g *GetBinResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Brand     types.Optional[string] `json:"brand"`
		BrandName types.Optional[string] `json:"brandName"`
		Gaps      types.Optional[[]int]  `json:"gaps"`
		Lenghts   types.Optional[[]int]  `json:"lenghts"`
		Mask      types.Optional[string] `json:"mask"`
		Input     types.Optional[string] `json:"input"`
		Cvv       types.Optional[int]    `json:"cvv"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Brand = temp.Brand
	g.BrandName = temp.BrandName
	g.Gaps = temp.Gaps
	g.Lenghts = temp.Lenghts
	g.Mask = temp.Mask
	g.Input = temp.Input
	g.Cvv = temp.Cvv
	return nil
}
