package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Generic response object for getting a MovementObjectBase.
type GetMovementObjectBaseResponse struct {
	Object    *string                `json:"object,omitempty"`
	Id        types.Optional[string] `json:"id"`
	Status    types.Optional[string] `json:"status"`
	Amount    types.Optional[string] `json:"amount"`
	CreatedAt types.Optional[string] `json:"created_at"`
	Type      types.Optional[string] `json:"type"`
	ChargeId  types.Optional[string] `json:"charge_id"`
	GatewayId types.Optional[string] `json:"gateway_id"`
}

func (g *GetMovementObjectBaseResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetMovementObjectBaseResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "MovementObject"
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.ChargeId.IsValueSet() {
		structMap["charge_id"] = g.ChargeId.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	return structMap
}

func (g *GetMovementObjectBaseResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Object    *string                `json:"object,omitempty"`
		Id        types.Optional[string] `json:"id"`
		Status    types.Optional[string] `json:"status"`
		Amount    types.Optional[string] `json:"amount"`
		CreatedAt types.Optional[string] `json:"created_at"`
		Type      types.Optional[string] `json:"type"`
		ChargeId  types.Optional[string] `json:"charge_id"`
		GatewayId types.Optional[string] `json:"gateway_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Object = temp.Object
	g.Id = temp.Id
	g.Status = temp.Status
	g.Amount = temp.Amount
	g.CreatedAt = temp.CreatedAt
	g.Type = temp.Type
	g.ChargeId = temp.ChargeId
	g.GatewayId = temp.GatewayId
	return nil
}

// Getter for object.
func (g *GetMovementObjectBaseResponse) GetObject() *string {
	return g.Object
}

// Getter for id.
func (g *GetMovementObjectBaseResponse) GetId() types.Optional[string] {
	return g.Id
}

// Getter for status.
func (g *GetMovementObjectBaseResponse) GetStatus() types.Optional[string] {
	return g.Status
}

// Getter for amount.
func (g *GetMovementObjectBaseResponse) GetAmount() types.Optional[string] {
	return g.Amount
}

// Getter for created_at.
func (g *GetMovementObjectBaseResponse) GetCreatedAt() types.Optional[string] {
	return g.CreatedAt
}

// Getter for type.
func (g *GetMovementObjectBaseResponse) GetType() types.Optional[string] {
	return g.Type
}

// Getter for charge_id.
func (g *GetMovementObjectBaseResponse) GetChargeId() types.Optional[string] {
	return g.ChargeId
}

// Getter for gateway_id.
func (g *GetMovementObjectBaseResponse) GetGatewayId() types.Optional[string] {
	return g.GatewayId
}

func UnmarshalGetMovementObjectBaseResponse(input []byte) (
	GetMovementObjectBaseResponseInterface,
	error) {
	discrim := &struct {
		Object *string
	}{}

	err := json.Unmarshal(input, &discrim)
	if err != nil {
		return nil, err
	}
	if discrim.Object == nil {
		object := "MovementObject"
		discrim.Object = &object
	}

	var res GetMovementObjectBaseResponseInterface
	switch *discrim.Object {
	case "refund":
		res = &GetMovementObjectRefundResponse{}
	case "feeCollection":
		res = &GetMovementObjectFeeCollectionResponse{}
	case "payable":
		res = &GetMovementObjectPayableResponse{}
	case "transfer":
		res = &GetMovementObjectTransferResponse{}
	default:
		res = &GetMovementObjectBaseResponse{}
	}
	json.Unmarshal(input, res)
	return res, nil
}

func ToGetMovementObjectBaseResponseArray(getMovementObjectBaseResponse []GetMovementObjectBaseResponseField) []GetMovementObjectBaseResponseInterface {
	var items []GetMovementObjectBaseResponseInterface
	for _, temp := range getMovementObjectBaseResponse {
		items = append(items, temp.GetMovementObjectBaseResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetMovementObjectBaseResponseField struct {
	GetMovementObjectBaseResponseInterface
}

func (g *GetMovementObjectBaseResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectBaseResponseInterface, err = UnmarshalGetMovementObjectBaseResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetMovementObjectBaseResponseSlice struct {
	Slice []GetMovementObjectBaseResponseInterface
}

func (ga *GetMovementObjectBaseResponseSlice) UnmarshalJSON(input []byte) error {
	tempStruct := struct {
		Slice []json.RawMessage
	}{}
	temp := []json.RawMessage{}
	err := json.Unmarshal(input, &tempStruct)
	if err != nil {
		err := json.Unmarshal(input, &temp)
		if err != nil {
			return err
		}
	} else {
		temp = tempStruct.Slice
	}

	ga.Slice = nil
	if temp != nil {
		ga.Slice = []GetMovementObjectBaseResponseInterface{}
		for _, getMovementObjectBaseResponse := range temp {
			if getMovementObjectBaseResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectBaseResponse
			err := json.Unmarshal(getMovementObjectBaseResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
