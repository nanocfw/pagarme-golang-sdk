package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponse struct {
	GetTransactionResponse
	StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
	AcquirerName            *string                         `json:"acquirer_name,omitempty"`
	AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
	AcquirerTid             *string                         `json:"acquirer_tid,omitempty"`
	AcquirerNsu             *string                         `json:"acquirer_nsu,omitempty"`
	AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
	OperationType           types.Optional[string]          `json:"operation_type"`
	Card                    types.Optional[GetCardResponse] `json:"card"`
	AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
	AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
	Installments            types.Optional[int]             `json:"installments"`
	ThreedAuthenticationUrl types.Optional[string]          `json:"threed_authentication_url"`
	FundingSource           types.Optional[string]          `json:"funding_source"`
}

func (g *GetCreditCardTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCreditCardTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "credit_card"
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.AcquirerName != nil {
		structMap["acquirer_name"] = g.AcquirerName
	}
	if g.AcquirerAffiliationCode.IsValueSet() {
		structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode.Value()
	}
	if g.AcquirerTid != nil {
		structMap["acquirer_tid"] = g.AcquirerTid
	}
	if g.AcquirerNsu != nil {
		structMap["acquirer_nsu"] = g.AcquirerNsu
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
	if g.ThreedAuthenticationUrl.IsValueSet() {
		structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl.Value()
	}
	if g.FundingSource.IsValueSet() {
		structMap["funding_source"] = g.FundingSource.Value()
	}
	return structMap
}

func (g *GetCreditCardTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     types.Optional[string]          `json:"statement_descriptor"`
		AcquirerName            *string                         `json:"acquirer_name,omitempty"`
		AcquirerAffiliationCode types.Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             *string                         `json:"acquirer_tid,omitempty"`
		AcquirerNsu             *string                         `json:"acquirer_nsu,omitempty"`
		AcquirerAuthCode        types.Optional[string]          `json:"acquirer_auth_code"`
		OperationType           types.Optional[string]          `json:"operation_type"`
		Card                    types.Optional[GetCardResponse] `json:"card"`
		AcquirerMessage         types.Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      types.Optional[string]          `json:"acquirer_return_code"`
		Installments            types.Optional[int]             `json:"installments"`
		ThreedAuthenticationUrl types.Optional[string]          `json:"threed_authentication_url"`
		FundingSource           types.Optional[string]          `json:"funding_source"`
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
	g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
	g.FundingSource = temp.FundingSource
	return nil
}

// Getter for statement_descriptor.
func (g *GetCreditCardTransactionResponse) GetStatementDescriptor() types.Optional[string] {
	return g.StatementDescriptor
}

// Getter for acquirer_name.
func (g *GetCreditCardTransactionResponse) GetAcquirerName() *string {
	return g.AcquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerAffiliationCode() types.Optional[string] {
	return g.AcquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetCreditCardTransactionResponse) GetAcquirerTid() *string {
	return g.AcquirerTid
}

// Getter for acquirer_nsu.
func (g *GetCreditCardTransactionResponse) GetAcquirerNsu() *string {
	return g.AcquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerAuthCode() types.Optional[string] {
	return g.AcquirerAuthCode
}

// Getter for operation_type.
func (g *GetCreditCardTransactionResponse) GetOperationType() types.Optional[string] {
	return g.OperationType
}

// Getter for card.
func (g *GetCreditCardTransactionResponse) GetCard() types.Optional[GetCardResponse] {
	return g.Card
}

// Getter for acquirer_message.
func (g *GetCreditCardTransactionResponse) GetAcquirerMessage() types.Optional[string] {
	return g.AcquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerReturnCode() types.Optional[string] {
	return g.AcquirerReturnCode
}

// Getter for installments.
func (g *GetCreditCardTransactionResponse) GetInstallments() types.Optional[int] {
	return g.Installments
}

// Getter for threed_authentication_url.
func (g *GetCreditCardTransactionResponse) GetThreedAuthenticationUrl() types.Optional[string] {
	return g.ThreedAuthenticationUrl
}

// Getter for funding_source.
func (g *GetCreditCardTransactionResponse) GetFundingSource() types.Optional[string] {
	return g.FundingSource
}

func UnmarshalGetCreditCardTransactionResponse(input []byte) (
	GetCreditCardTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getCreditCardTransactionResponse, ok := getTransactionResponse.(GetCreditCardTransactionResponseInterface); ok {
		return getCreditCardTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getCreditCardTransactionResponse %v", getTransactionResponse)
}

func ToGetCreditCardTransactionResponseArray(getCreditCardTransactionResponse []GetCreditCardTransactionResponseField) []GetCreditCardTransactionResponseInterface {
	var items []GetCreditCardTransactionResponseInterface
	for _, temp := range getCreditCardTransactionResponse {
		items = append(items, temp.GetCreditCardTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetCreditCardTransactionResponseField struct {
	GetCreditCardTransactionResponseInterface
}

func (g *GetCreditCardTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetCreditCardTransactionResponseInterface, err = UnmarshalGetCreditCardTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetCreditCardTransactionResponseSlice struct {
	Slice []GetCreditCardTransactionResponseInterface
}

func (ga *GetCreditCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetCreditCardTransactionResponseInterface{}
		for _, getCreditCardTransactionResponse := range temp {
			if getCreditCardTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetCreditCardTransactionResponse
			err := json.Unmarshal(getCreditCardTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
