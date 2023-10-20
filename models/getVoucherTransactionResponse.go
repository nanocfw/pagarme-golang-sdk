package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Response for voucher transactions
type GetVoucherTransactionResponse struct {
	GetTransactionResponse
	StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
	AcquirerName            types.Optional[string]          `json:"acquirer_name"`
	AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
	AcquirerTid             types.Optional[string]          `json:"acquirer_tid"`
	AcquirerNsu             types.Optional[string]          `json:"acquirer_nsu"`
	AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
	AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
	AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
	OperationType           types.Optional[string]          `json:"operation_type"`
	Card                    types.Optional[GetCardResponse] `json:"card"`
}

func (g *GetVoucherTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetVoucherTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "voucher"
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
	if g.AcquirerMessage.IsValueSet() {
		structMap["acquirer_message"] = g.AcquirerMessage.Value()
	}
	if g.AcquirerReturnCode.IsValueSet() {
		structMap["acquirer_return_code"] = g.AcquirerReturnCode.Value()
	}
	if g.OperationType.IsValueSet() {
		structMap["operation_type"] = g.OperationType.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	return structMap
}

func (g *GetVoucherTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
		AcquirerName            types.Optional[string]          `json:"acquirer_name"`
		AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             types.Optional[string]          `json:"acquirer_tid"`
		AcquirerNsu             types.Optional[string]          `json:"acquirer_nsu"`
		AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
		AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
		OperationType           types.Optional[string]          `json:"operation_type"`
		Card                    types.Optional[GetCardResponse] `json:"card"`
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
	g.AcquirerMessage = temp.AcquirerMessage
	g.AcquirerReturnCode = temp.AcquirerReturnCode
	g.OperationType = temp.OperationType
	g.Card = temp.Card
	return nil
}

// Getter for statement_descriptor.
func (g *GetVoucherTransactionResponse) GetStatementDescriptor() types.Optional[string] {
	return g.StatementDescriptor
}

// Getter for acquirer_name.
func (g *GetVoucherTransactionResponse) GetAcquirerName() types.Optional[string] {
	return g.AcquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetVoucherTransactionResponse) GetAcquirerAffiliationCode() types.Optional[string] {
	return g.AcquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetVoucherTransactionResponse) GetAcquirerTid() types.Optional[string] {
	return g.AcquirerTid
}

// Getter for acquirer_nsu.
func (g *GetVoucherTransactionResponse) GetAcquirerNsu() types.Optional[string] {
	return g.AcquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetVoucherTransactionResponse) GetAcquirerAuthCode() types.Optional[string] {
	return g.AcquirerAuthCode
}

// Getter for acquirer_message.
func (g *GetVoucherTransactionResponse) GetAcquirerMessage() types.Optional[string] {
	return g.AcquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetVoucherTransactionResponse) GetAcquirerReturnCode() types.Optional[string] {
	return g.AcquirerReturnCode
}

// Getter for operation_type.
func (g *GetVoucherTransactionResponse) GetOperationType() types.Optional[string] {
	return g.OperationType
}

// Getter for card.
func (g *GetVoucherTransactionResponse) GetCard() types.Optional[GetCardResponse] {
	return g.Card
}

func UnmarshalGetVoucherTransactionResponse(input []byte) (
	GetVoucherTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getVoucherTransactionResponse, ok := getTransactionResponse.(GetVoucherTransactionResponseInterface); ok {
		return getVoucherTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getVoucherTransactionResponse %v", getTransactionResponse)
}

func ToGetVoucherTransactionResponseArray(getVoucherTransactionResponse []GetVoucherTransactionResponseField) []GetVoucherTransactionResponseInterface {
	var items []GetVoucherTransactionResponseInterface
	for _, temp := range getVoucherTransactionResponse {
		items = append(items, temp.GetVoucherTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetVoucherTransactionResponseField struct {
	GetVoucherTransactionResponseInterface
}

func (g *GetVoucherTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetVoucherTransactionResponseInterface, err = UnmarshalGetVoucherTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetVoucherTransactionResponseSlice struct {
	Slice []GetVoucherTransactionResponseInterface
}

func (ga *GetVoucherTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetVoucherTransactionResponseInterface{}
		for _, getVoucherTransactionResponse := range temp {
			if getVoucherTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetVoucherTransactionResponse
			err := json.Unmarshal(getVoucherTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
