package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponse struct {
	GetTransactionResponse
	StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
	AcquirerName            types.Optional[string]          `json:"acquirer_name"`
	AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
	AcquirerTid             types.Optional[string]          `json:"acquirer_tid"`
	AcquirerNsu             types.Optional[string]          `json:"acquirer_nsu"`
	AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
	OperationType           types.Optional[string]          `json:"operation_type"`
	Card                    types.Optional[GetCardResponse] `json:"card"`
	AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
	AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
	Installments            types.Optional[int]             `json:"installments"`
}

func (g *GetPrivateLabelTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPrivateLabelTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "private_label"
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.AcquirerName.IsValueSet() {
		structMap["acquirer_name"] = g.AcquirerName.Value()
	}
	if g.AcquirerAffiliationCode.IsValueSet() {
		structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode.Value()
	}
	if g.AcquirerTid.IsValueSet() {
		structMap["acquirer_tid"] = g.AcquirerTid.Value()
	}
	if g.AcquirerNsu.IsValueSet() {
		structMap["acquirer_nsu"] = g.AcquirerNsu.Value()
	}
	if g.AcquirerAuthCode.IsValueSet() {
		structMap["acquirer_auth_code"] = g.AcquirerAuthCode.Value()
	}
	if g.OperationType.IsValueSet() {
		structMap["operation_type"] = g.OperationType.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	if g.AcquirerMessage.IsValueSet() {
		structMap["acquirer_message"] = g.AcquirerMessage.Value()
	}
	if g.AcquirerReturnCode.IsValueSet() {
		structMap["acquirer_return_code"] = g.AcquirerReturnCode.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	return structMap
}

func (g *GetPrivateLabelTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
		AcquirerName            types.Optional[string]          `json:"acquirer_name"`
		AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             types.Optional[string]          `json:"acquirer_tid"`
		AcquirerNsu             types.Optional[string]          `json:"acquirer_nsu"`
		AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
		OperationType           types.Optional[string]          `json:"operation_type"`
		Card                    types.Optional[GetCardResponse] `json:"card"`
		AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
		Installments            types.Optional[int]             `json:"installments"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StatementDescriptor = temp.StatementDescriptor
	g.AcquirerName = temp.AcquirerName
	g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
	g.AcquirerTid = temp.AcquirerTid
	g.AcquirerNsu = temp.AcquirerNsu
	g.AcquirerAuthCode = temp.AcquirerAuthCode
	g.OperationType = temp.OperationType
	g.Card = temp.Card
	g.AcquirerMessage = temp.AcquirerMessage
	g.AcquirerReturnCode = temp.AcquirerReturnCode
	g.Installments = temp.Installments
	return nil
}

// Getter for statement_descriptor.
func (g *GetPrivateLabelTransactionResponse) GetStatementDescriptor() types.Optional[string] {
	return g.StatementDescriptor
}

// Getter for acquirer_name.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerName() types.Optional[string] {
	return g.AcquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAffiliationCode() types.Optional[string] {
	return g.AcquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerTid() types.Optional[string] {
	return g.AcquirerTid
}

// Getter for acquirer_nsu.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerNsu() types.Optional[string] {
	return g.AcquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAuthCode() types.Optional[string] {
	return g.AcquirerAuthCode
}

// Getter for operation_type.
func (g *GetPrivateLabelTransactionResponse) GetOperationType() types.Optional[string] {
	return g.OperationType
}

// Getter for card.
func (g *GetPrivateLabelTransactionResponse) GetCard() types.Optional[GetCardResponse] {
	return g.Card
}

// Getter for acquirer_message.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerMessage() types.Optional[string] {
	return g.AcquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerReturnCode() types.Optional[string] {
	return g.AcquirerReturnCode
}

// Getter for installments.
func (g *GetPrivateLabelTransactionResponse) GetInstallments() types.Optional[int] {
	return g.Installments
}

func UnmarshalGetPrivateLabelTransactionResponse(input []byte) (
	GetPrivateLabelTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getPrivateLabelTransactionResponse, ok := getTransactionResponse.(GetPrivateLabelTransactionResponseInterface); ok {
		return getPrivateLabelTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getPrivateLabelTransactionResponse %v", getTransactionResponse)
}

func ToGetPrivateLabelTransactionResponseArray(getPrivateLabelTransactionResponse []GetPrivateLabelTransactionResponseField) []GetPrivateLabelTransactionResponseInterface {
	var items []GetPrivateLabelTransactionResponseInterface
	for _, temp := range getPrivateLabelTransactionResponse {
		items = append(items, temp.GetPrivateLabelTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetPrivateLabelTransactionResponseField struct {
	GetPrivateLabelTransactionResponseInterface
}

func (g *GetPrivateLabelTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetPrivateLabelTransactionResponseInterface, err = UnmarshalGetPrivateLabelTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetPrivateLabelTransactionResponseSlice struct {
	Slice []GetPrivateLabelTransactionResponseInterface
}

func (ga *GetPrivateLabelTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetPrivateLabelTransactionResponseInterface{}
		for _, getPrivateLabelTransactionResponse := range temp {
			if getPrivateLabelTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetPrivateLabelTransactionResponse
			err := json.Unmarshal(getPrivateLabelTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
