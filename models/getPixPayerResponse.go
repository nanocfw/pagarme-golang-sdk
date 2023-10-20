package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Pix payer data.
type GetPixPayerResponse struct {
	Name         types.Optional[string]                    `json:"name"`
	Document     types.Optional[string]                    `json:"document"`
	DocumentType types.Optional[string]                    `json:"document_type"`
	BankAccount  types.Optional[GetPixBankAccountResponse] `json:"bank_account"`
}

func (g *GetPixPayerResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPixPayerResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Document.IsValueSet() {
		structMap["document"] = g.Document.Value()
	}
	if g.DocumentType.IsValueSet() {
		structMap["document_type"] = g.DocumentType.Value()
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	return structMap
}

func (g *GetPixPayerResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name         types.Optional[string]                    `json:"name"`
		Document     types.Optional[string]                    `json:"document"`
		DocumentType types.Optional[string]                    `json:"document_type"`
		BankAccount  types.Optional[GetPixBankAccountResponse] `json:"bank_account"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Name = temp.Name
	g.Document = temp.Document
	g.DocumentType = temp.DocumentType
	g.BankAccount = temp.BankAccount
	return nil
}
