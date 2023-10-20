package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponse struct {
	GetTransactionResponse
	Url        *string    `json:"url,omitempty"`
	BankTid    *string    `json:"bank_tid,omitempty"`
	Bank       *string    `json:"bank,omitempty"`
	PaidAt     *time.Time `json:"paid_at,omitempty"`
	PaidAmount *int       `json:"paid_amount,omitempty"`
}

func (g *GetBankTransferTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetBankTransferTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "bank_transfer"
	}
	if g.Url != nil {
		structMap["url"] = g.Url
	}
	if g.BankTid != nil {
		structMap["bank_tid"] = g.BankTid
	}
	if g.Bank != nil {
		structMap["bank"] = g.Bank
	}
	if g.PaidAt != nil {
		structMap["paid_at"] = g.PaidAt.Format(time.RFC3339)
	}
	if g.PaidAmount != nil {
		structMap["paid_amount"] = g.PaidAmount
	}
	return structMap
}

func (g *GetBankTransferTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Url        *string `json:"url,omitempty"`
		BankTid    *string `json:"bank_tid,omitempty"`
		Bank       *string `json:"bank,omitempty"`
		PaidAt     *string `json:"paid_at,omitempty"`
		PaidAmount *int    `json:"paid_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.BankTid = temp.BankTid
	g.Bank = temp.Bank
	if temp.PaidAt != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, *temp.PaidAt)
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt = &PaidAtVal
	}
	g.PaidAmount = temp.PaidAmount
	return nil
}

// Getter for url.
func (g *GetBankTransferTransactionResponse) GetUrl() *string {
	return g.Url
}

// Getter for bank_tid.
func (g *GetBankTransferTransactionResponse) GetBankTid() *string {
	return g.BankTid
}

// Getter for bank.
func (g *GetBankTransferTransactionResponse) GetBank() *string {
	return g.Bank
}

// Getter for paid_at.
func (g *GetBankTransferTransactionResponse) GetPaidAt() *time.Time {
	return g.PaidAt
}

// Getter for paid_amount.
func (g *GetBankTransferTransactionResponse) GetPaidAmount() *int {
	return g.PaidAmount
}

func UnmarshalGetBankTransferTransactionResponse(input []byte) (
	GetBankTransferTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getBankTransferTransactionResponse, ok := getTransactionResponse.(GetBankTransferTransactionResponseInterface); ok {
		return getBankTransferTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getBankTransferTransactionResponse %v", getTransactionResponse)
}

func ToGetBankTransferTransactionResponseArray(getBankTransferTransactionResponse []GetBankTransferTransactionResponseField) []GetBankTransferTransactionResponseInterface {
	var items []GetBankTransferTransactionResponseInterface
	for _, temp := range getBankTransferTransactionResponse {
		items = append(items, temp.GetBankTransferTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetBankTransferTransactionResponseField struct {
	GetBankTransferTransactionResponseInterface
}

func (g *GetBankTransferTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetBankTransferTransactionResponseInterface, err = UnmarshalGetBankTransferTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetBankTransferTransactionResponseSlice struct {
	Slice []GetBankTransferTransactionResponseInterface
}

func (ga *GetBankTransferTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetBankTransferTransactionResponseInterface{}
		for _, getBankTransferTransactionResponse := range temp {
			if getBankTransferTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetBankTransferTransactionResponse
			err := json.Unmarshal(getBankTransferTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
