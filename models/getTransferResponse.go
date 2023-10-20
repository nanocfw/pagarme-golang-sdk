package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Transfer response
type GetTransferResponse struct {
	Id          types.Optional[string]                 `json:"id"`
	Amount      types.Optional[int]                    `json:"amount"`
	Status      types.Optional[string]                 `json:"status"`
	CreatedAt   types.Optional[time.Time]              `json:"created_at"`
	UpdatedAt   types.Optional[time.Time]              `json:"updated_at"`
	BankAccount types.Optional[GetBankAccountResponse] `json:"bank_account"`
	Metadata    types.Optional[map[string]string]      `json:"metadata"`
}

func (g *GetTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTransferResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	return structMap
}

func (g *GetTransferResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          types.Optional[string]                 `json:"id"`
		Amount      types.Optional[int]                    `json:"amount"`
		Status      types.Optional[string]                 `json:"status"`
		CreatedAt   types.Optional[string]                 `json:"created_at"`
		UpdatedAt   types.Optional[string]                 `json:"updated_at"`
		BankAccount types.Optional[GetBankAccountResponse] `json:"bank_account"`
		Metadata    types.Optional[map[string]string]      `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.BankAccount = temp.BankAccount
	g.Metadata = temp.Metadata
	return nil
}
