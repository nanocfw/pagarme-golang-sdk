package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// 3D-S payment authentication response
type GetThreeDSecureResponse struct {
	Mpi           types.Optional[string] `json:"mpi"`
	Eci           types.Optional[string] `json:"eci"`
	Cavv          types.Optional[string] `json:"cavv"`
	TransactionId types.Optional[string] `json:"transaction_Id"`
	SuccessUrl    types.Optional[string] `json:"success_url"`
}

func (g *GetThreeDSecureResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetThreeDSecureResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Mpi.IsValueSet() {
		structMap["mpi"] = g.Mpi.Value()
	}
	if g.Eci.IsValueSet() {
		structMap["eci"] = g.Eci.Value()
	}
	if g.Cavv.IsValueSet() {
		structMap["cavv"] = g.Cavv.Value()
	}
	if g.TransactionId.IsValueSet() {
		structMap["transaction_Id"] = g.TransactionId.Value()
	}
	if g.SuccessUrl.IsValueSet() {
		structMap["success_url"] = g.SuccessUrl.Value()
	}
	return structMap
}

func (g *GetThreeDSecureResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Mpi           types.Optional[string] `json:"mpi"`
		Eci           types.Optional[string] `json:"eci"`
		Cavv          types.Optional[string] `json:"cavv"`
		TransactionId types.Optional[string] `json:"transaction_Id"`
		SuccessUrl    types.Optional[string] `json:"success_url"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Mpi = temp.Mpi
	g.Eci = temp.Eci
	g.Cavv = temp.Cavv
	g.TransactionId = temp.TransactionId
	g.SuccessUrl = temp.SuccessUrl
	return nil
}
