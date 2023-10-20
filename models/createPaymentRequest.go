package models

import (
	"encoding/json"
)

// Payment data
type CreatePaymentRequest struct {
	PaymentMethod        string                            `json:"payment_method"`
	CreditCard           *CreateCreditCardPaymentRequest   `json:"credit_card,omitempty"`
	DebitCard            *CreateDebitCardPaymentRequest    `json:"debit_card,omitempty"`
	Boleto               *CreateBoletoPaymentRequest       `json:"boleto,omitempty"`
	Currency             *string                           `json:"currency,omitempty"`
	Voucher              *CreateVoucherPaymentRequest      `json:"voucher,omitempty"`
	Split                *[]CreateSplitRequest             `json:"split,omitempty"`
	BankTransfer         *CreateBankTransferPaymentRequest `json:"bank_transfer,omitempty"`
	GatewayAffiliationId *string                           `json:"gateway_affiliation_id,omitempty"`
	Amount               *int                              `json:"amount,omitempty"`
	Checkout             *CreateCheckoutPaymentRequest     `json:"checkout,omitempty"`
	CustomerId           *string                           `json:"customer_id,omitempty"`
	Customer             *CreateCustomerRequest            `json:"customer,omitempty"`
	Metadata             map[string]*string                `json:"metadata,omitempty"`
	Cash                 *CreateCashPaymentRequest         `json:"cash,omitempty"`
	PrivateLabel         *CreatePrivateLabelPaymentRequest `json:"private_label,omitempty"`
	Pix                  *CreatePixPaymentRequest          `json:"pix,omitempty"`
}

func (c *CreatePaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_method"] = c.PaymentMethod
	if c.CreditCard != nil {
		structMap["credit_card"] = c.CreditCard
	}
	if c.DebitCard != nil {
		structMap["debit_card"] = c.DebitCard
	}
	if c.Boleto != nil {
		structMap["boleto"] = c.Boleto
	}
	if c.Currency != nil {
		structMap["currency"] = c.Currency
	}
	if c.Voucher != nil {
		structMap["voucher"] = c.Voucher
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	if c.BankTransfer != nil {
		structMap["bank_transfer"] = c.BankTransfer
	}
	if c.GatewayAffiliationId != nil {
		structMap["gateway_affiliation_id"] = c.GatewayAffiliationId
	}
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.Checkout != nil {
		structMap["checkout"] = c.Checkout
	}
	if c.CustomerId != nil {
		structMap["customer_id"] = c.CustomerId
	}
	if c.Customer != nil {
		structMap["customer"] = c.Customer
	}
	if c.Metadata != nil {
		structMap["metadata"] = c.Metadata
	}
	if c.Cash != nil {
		structMap["cash"] = c.Cash
	}
	if c.PrivateLabel != nil {
		structMap["private_label"] = c.PrivateLabel
	}
	if c.Pix != nil {
		structMap["pix"] = c.Pix
	}
	return structMap
}

func (c *CreatePaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentMethod        string                            `json:"payment_method"`
		CreditCard           *CreateCreditCardPaymentRequest   `json:"credit_card,omitempty"`
		DebitCard            *CreateDebitCardPaymentRequest    `json:"debit_card,omitempty"`
		Boleto               *CreateBoletoPaymentRequest       `json:"boleto,omitempty"`
		Currency             *string                           `json:"currency,omitempty"`
		Voucher              *CreateVoucherPaymentRequest      `json:"voucher,omitempty"`
		Split                *[]CreateSplitRequest             `json:"split,omitempty"`
		BankTransfer         *CreateBankTransferPaymentRequest `json:"bank_transfer,omitempty"`
		GatewayAffiliationId *string                           `json:"gateway_affiliation_id,omitempty"`
		Amount               *int                              `json:"amount,omitempty"`
		Checkout             *CreateCheckoutPaymentRequest     `json:"checkout,omitempty"`
		CustomerId           *string                           `json:"customer_id,omitempty"`
		Customer             *CreateCustomerRequest            `json:"customer,omitempty"`
		Metadata             map[string]*string                `json:"metadata,omitempty"`
		Cash                 *CreateCashPaymentRequest         `json:"cash,omitempty"`
		PrivateLabel         *CreatePrivateLabelPaymentRequest `json:"private_label,omitempty"`
		Pix                  *CreatePixPaymentRequest          `json:"pix,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PaymentMethod = temp.PaymentMethod
	c.CreditCard = temp.CreditCard
	c.DebitCard = temp.DebitCard
	c.Boleto = temp.Boleto
	c.Currency = temp.Currency
	c.Voucher = temp.Voucher
	c.Split = temp.Split
	c.BankTransfer = temp.BankTransfer
	c.GatewayAffiliationId = temp.GatewayAffiliationId
	c.Amount = temp.Amount
	c.Checkout = temp.Checkout
	c.CustomerId = temp.CustomerId
	c.Customer = temp.Customer
	c.Metadata = temp.Metadata
	c.Cash = temp.Cash
	c.PrivateLabel = temp.PrivateLabel
	c.Pix = temp.Pix
	return nil
}
