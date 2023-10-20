package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a cash transaction
type GetCashTransactionResponse struct {
	GetTransactionResponse
	Description types.Optional[string] `json:"description"`
}

func (g *GetCashTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCashTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "cash"
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	return structMap
}

func (g *GetCashTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Description types.Optional[string] `json:"description"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Description = temp.Description
	return nil
}

// Getter for description.
func (g *GetCashTransactionResponse) GetDescription() types.Optional[string] {
	return g.Description
}

func UnmarshalGetCashTransactionResponse(input []byte) (
	GetCashTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getCashTransactionResponse, ok := getTransactionResponse.(GetCashTransactionResponseInterface); ok {
		return getCashTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getCashTransactionResponse %v", getTransactionResponse)
}

func ToGetCashTransactionResponseArray(getCashTransactionResponse []GetCashTransactionResponseField) []GetCashTransactionResponseInterface {
	var items []GetCashTransactionResponseInterface
	for _, temp := range getCashTransactionResponse {
		items = append(items, temp.GetCashTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetCashTransactionResponseField struct {
	GetCashTransactionResponseInterface
}

func (g *GetCashTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetCashTransactionResponseInterface, err = UnmarshalGetCashTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetCashTransactionResponseSlice struct {
	Slice []GetCashTransactionResponseInterface
}

func (ga *GetCashTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetCashTransactionResponseInterface{}
		for _, getCashTransactionResponse := range temp {
			if getCashTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetCashTransactionResponse
			err := json.Unmarshal(getCashTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
