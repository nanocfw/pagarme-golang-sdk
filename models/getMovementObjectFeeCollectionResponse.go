package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Generic response object for getting a MovementObjectFeeCollection.
type GetMovementObjectFeeCollectionResponse struct {
	GetMovementObjectBaseResponse
	Description types.Optional[string] `json:"description"`
	PaymentDate types.Optional[string] `json:"payment_date"`
	RecipientId types.Optional[string] `json:"recipient_id"`
}

func (g *GetMovementObjectFeeCollectionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetMovementObjectFeeCollectionResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "feeCollection"
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.PaymentDate.IsValueSet() {
		structMap["payment_date"] = g.PaymentDate.Value()
	}
	if g.RecipientId.IsValueSet() {
		structMap["recipient_id"] = g.RecipientId.Value()
	}
	return structMap
}

func (g *GetMovementObjectFeeCollectionResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		Description types.Optional[string] `json:"description"`
		PaymentDate types.Optional[string] `json:"payment_date"`
		RecipientId types.Optional[string] `json:"recipient_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Description = temp.Description
	g.PaymentDate = temp.PaymentDate
	g.RecipientId = temp.RecipientId
	return nil
}

// Getter for description.
func (g *GetMovementObjectFeeCollectionResponse) GetDescription() types.Optional[string] {
	return g.Description
}

// Getter for payment_date.
func (g *GetMovementObjectFeeCollectionResponse) GetPaymentDate() types.Optional[string] {
	return g.PaymentDate
}

// Getter for recipient_id.
func (g *GetMovementObjectFeeCollectionResponse) GetRecipientId() types.Optional[string] {
	return g.RecipientId
}

func UnmarshalGetMovementObjectFeeCollectionResponse(input []byte) (
	GetMovementObjectFeeCollectionResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectFeeCollectionResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectFeeCollectionResponseInterface); ok {
		return getMovementObjectFeeCollectionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectFeeCollectionResponse %v", getMovementObjectBaseResponse)
}

func ToGetMovementObjectFeeCollectionResponseArray(getMovementObjectFeeCollectionResponse []GetMovementObjectFeeCollectionResponseField) []GetMovementObjectFeeCollectionResponseInterface {
	var items []GetMovementObjectFeeCollectionResponseInterface
	for _, temp := range getMovementObjectFeeCollectionResponse {
		items = append(items, temp.GetMovementObjectFeeCollectionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetMovementObjectFeeCollectionResponseField struct {
	GetMovementObjectFeeCollectionResponseInterface
}

func (g *GetMovementObjectFeeCollectionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectFeeCollectionResponseInterface, err = UnmarshalGetMovementObjectFeeCollectionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetMovementObjectFeeCollectionResponseSlice struct {
	Slice []GetMovementObjectFeeCollectionResponseInterface
}

func (ga *GetMovementObjectFeeCollectionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectFeeCollectionResponseInterface{}
		for _, getMovementObjectFeeCollectionResponse := range temp {
			if getMovementObjectFeeCollectionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectFeeCollectionResponse
			err := json.Unmarshal(getMovementObjectFeeCollectionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
