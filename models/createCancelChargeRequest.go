package models

import (
	"encoding/json"
)

// Request for canceling a charge.
type CreateCancelChargeRequest struct {
	Amount             *int                                   `json:"amount,omitempty"`
	SplitRules         *[]CreateCancelChargeSplitRulesRequest `json:"split_rules,omitempty"`
	Split              *[]CreateSplitRequest                  `json:"split,omitempty"`
	OperationReference string                                 `json:"operation_reference"`
	BankAccount        *CreateBankAccountRefundingDTO         `json:"bank_account,omitempty"`
}

func (c *CreateCancelChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCancelChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.SplitRules != nil {
		structMap["split_rules"] = c.SplitRules
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	structMap["operation_reference"] = c.OperationReference
	if c.BankAccount != nil {
		structMap["bank_account"] = c.BankAccount
	}
	return structMap
}

func (c *CreateCancelChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount             *int                                   `json:"amount,omitempty"`
		SplitRules         *[]CreateCancelChargeSplitRulesRequest `json:"split_rules,omitempty"`
		Split              *[]CreateSplitRequest                  `json:"split,omitempty"`
		OperationReference string                                 `json:"operation_reference"`
		BankAccount        *CreateBankAccountRefundingDTO         `json:"bank_account,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.SplitRules = temp.SplitRules
	c.Split = temp.Split
	c.OperationReference = temp.OperationReference
	c.BankAccount = temp.BankAccount
	return nil
}
