package models

import (
	"encoding/json"
	"fmt"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a debit card transaction
type GetDebitCardTransactionResponse struct {
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
	Mpi                     types.Optional[string]          `json:"mpi"`
	Eci                     types.Optional[string]          `json:"eci"`
	AuthenticationType      types.Optional[string]          `json:"authentication_type"`
	ThreedAuthenticationUrl types.Optional[string]          `json:"threed_authentication_url"`
	FundingSource           types.Optional[string]          `json:"funding_source"`
}

func (g *GetDebitCardTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetDebitCardTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "debit_card"
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
	if g.Mpi.IsValueSet() {
		structMap["mpi"] = g.Mpi.Value()
	}
	if g.Eci.IsValueSet() {
		structMap["eci"] = g.Eci.Value()
	}
	if g.AuthenticationType.IsValueSet() {
		structMap["authentication_type"] = g.AuthenticationType.Value()
	}
	if g.ThreedAuthenticationUrl.IsValueSet() {
		structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl.Value()
	}
	if g.FundingSource.IsValueSet() {
		structMap["funding_source"] = g.FundingSource.Value()
	}
	return structMap
}

func (g *GetDebitCardTransactionResponse) UnmarshalJSON(input []byte) error {
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
		Mpi                     types.Optional[string]          `json:"mpi"`
		Eci                     types.Optional[string]          `json:"eci"`
		AuthenticationType      types.Optional[string]          `json:"authentication_type"`
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
	g.Mpi = temp.Mpi
	g.Eci = temp.Eci
	g.AuthenticationType = temp.AuthenticationType
	g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
	g.FundingSource = temp.FundingSource
	return nil
}

// Getter for statement_descriptor.
func (g *GetDebitCardTransactionResponse) GetStatementDescriptor() types.Optional[string] {
	return g.StatementDescriptor
}

// Getter for acquirer_name.
func (g *GetDebitCardTransactionResponse) GetAcquirerName() types.Optional[string] {
	return g.AcquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetDebitCardTransactionResponse) GetAcquirerAffiliationCode() types.Optional[string] {
	return g.AcquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetDebitCardTransactionResponse) GetAcquirerTid() types.Optional[string] {
	return g.AcquirerTid
}

// Getter for acquirer_nsu.
func (g *GetDebitCardTransactionResponse) GetAcquirerNsu() types.Optional[string] {
	return g.AcquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetDebitCardTransactionResponse) GetAcquirerAuthCode() types.Optional[string] {
	return g.AcquirerAuthCode
}

// Getter for operation_type.
func (g *GetDebitCardTransactionResponse) GetOperationType() types.Optional[string] {
	return g.OperationType
}

// Getter for card.
func (g *GetDebitCardTransactionResponse) GetCard() types.Optional[GetCardResponse] {
	return g.Card
}

// Getter for acquirer_message.
func (g *GetDebitCardTransactionResponse) GetAcquirerMessage() types.Optional[string] {
	return g.AcquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetDebitCardTransactionResponse) GetAcquirerReturnCode() types.Optional[string] {
	return g.AcquirerReturnCode
}

// Getter for mpi.
func (g *GetDebitCardTransactionResponse) GetMpi() types.Optional[string] {
	return g.Mpi
}

// Getter for eci.
func (g *GetDebitCardTransactionResponse) GetEci() types.Optional[string] {
	return g.Eci
}

// Getter for authentication_type.
func (g *GetDebitCardTransactionResponse) GetAuthenticationType() types.Optional[string] {
	return g.AuthenticationType
}

// Getter for threed_authentication_url.
func (g *GetDebitCardTransactionResponse) GetThreedAuthenticationUrl() types.Optional[string] {
	return g.ThreedAuthenticationUrl
}

// Getter for funding_source.
func (g *GetDebitCardTransactionResponse) GetFundingSource() types.Optional[string] {
	return g.FundingSource
}

func UnmarshalGetDebitCardTransactionResponse(input []byte) (
	GetDebitCardTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getDebitCardTransactionResponse, ok := getTransactionResponse.(GetDebitCardTransactionResponseInterface); ok {
		return getDebitCardTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getDebitCardTransactionResponse %v", getTransactionResponse)
}

func ToGetDebitCardTransactionResponseArray(getDebitCardTransactionResponse []GetDebitCardTransactionResponseField) []GetDebitCardTransactionResponseInterface {
	var items []GetDebitCardTransactionResponseInterface
	for _, temp := range getDebitCardTransactionResponse {
		items = append(items, temp.GetDebitCardTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetDebitCardTransactionResponseField struct {
	GetDebitCardTransactionResponseInterface
}

func (g *GetDebitCardTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetDebitCardTransactionResponseInterface, err = UnmarshalGetDebitCardTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetDebitCardTransactionResponseSlice struct {
	Slice []GetDebitCardTransactionResponseInterface
}

func (ga *GetDebitCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetDebitCardTransactionResponseInterface{}
		for _, getDebitCardTransactionResponse := range temp {
			if getDebitCardTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetDebitCardTransactionResponse
			err := json.Unmarshal(getDebitCardTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
