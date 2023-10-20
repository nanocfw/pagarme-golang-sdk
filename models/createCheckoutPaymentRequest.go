package models

import (
	"encoding/json"
)

// Checkout payment request
type CreateCheckoutPaymentRequest struct {
	AcceptedPaymentMethods      []string                                `json:"accepted_payment_methods"`
	AcceptedMultiPaymentMethods []interface{}                           `json:"accepted_multi_payment_methods"`
	SuccessUrl                  string                                  `json:"success_url"`
	DefaultPaymentMethod        *string                                 `json:"default_payment_method,omitempty"`
	GatewayAffiliationId        *string                                 `json:"gateway_affiliation_id,omitempty"`
	CreditCard                  *CreateCheckoutCreditCardPaymentRequest `json:"credit_card,omitempty"`
	DebitCard                   *CreateCheckoutDebitCardPaymentRequest  `json:"debit_card,omitempty"`
	Boleto                      *CreateCheckoutBoletoPaymentRequest     `json:"boleto,omitempty"`
	CustomerEditable            *bool                                   `json:"customer_editable,omitempty"`
	ExpiresIn                   *int                                    `json:"expires_in,omitempty"`
	SkipCheckoutSuccessPage     bool                                    `json:"skip_checkout_success_page"`
	BillingAddressEditable      bool                                    `json:"billing_address_editable"`
	BillingAddress              CreateAddressRequest                    `json:"billing_address"`
	BankTransfer                *CreateCheckoutBankTransferRequest      `json:"bank_transfer,omitempty"`
	AcceptedBrands              []string                                `json:"accepted_brands"`
	Pix                         *CreateCheckoutPixPaymentRequest        `json:"pix,omitempty"`
}

func (c *CreateCheckoutPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCheckoutPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["accepted_payment_methods"] = c.AcceptedPaymentMethods
	structMap["accepted_multi_payment_methods"] = c.AcceptedMultiPaymentMethods
	structMap["success_url"] = c.SuccessUrl
	if c.DefaultPaymentMethod != nil {
		structMap["default_payment_method"] = c.DefaultPaymentMethod
	}
	if c.GatewayAffiliationId != nil {
		structMap["gateway_affiliation_id"] = c.GatewayAffiliationId
	}
	if c.CreditCard != nil {
		structMap["credit_card"] = c.CreditCard
	}
	if c.DebitCard != nil {
		structMap["debit_card"] = c.DebitCard
	}
	if c.Boleto != nil {
		structMap["boleto"] = c.Boleto
	}
	if c.CustomerEditable != nil {
		structMap["customer_editable"] = c.CustomerEditable
	}
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	structMap["skip_checkout_success_page"] = c.SkipCheckoutSuccessPage
	structMap["billing_address_editable"] = c.BillingAddressEditable
	structMap["billing_address"] = c.BillingAddress
	if c.BankTransfer != nil {
		structMap["bank_transfer"] = c.BankTransfer
	}
	structMap["accepted_brands"] = c.AcceptedBrands
	if c.Pix != nil {
		structMap["pix"] = c.Pix
	}
	return structMap
}

func (c *CreateCheckoutPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AcceptedPaymentMethods      []string                                `json:"accepted_payment_methods"`
		AcceptedMultiPaymentMethods []interface{}                           `json:"accepted_multi_payment_methods"`
		SuccessUrl                  string                                  `json:"success_url"`
		DefaultPaymentMethod        *string                                 `json:"default_payment_method,omitempty"`
		GatewayAffiliationId        *string                                 `json:"gateway_affiliation_id,omitempty"`
		CreditCard                  *CreateCheckoutCreditCardPaymentRequest `json:"credit_card,omitempty"`
		DebitCard                   *CreateCheckoutDebitCardPaymentRequest  `json:"debit_card,omitempty"`
		Boleto                      *CreateCheckoutBoletoPaymentRequest     `json:"boleto,omitempty"`
		CustomerEditable            *bool                                   `json:"customer_editable,omitempty"`
		ExpiresIn                   *int                                    `json:"expires_in,omitempty"`
		SkipCheckoutSuccessPage     bool                                    `json:"skip_checkout_success_page"`
		BillingAddressEditable      bool                                    `json:"billing_address_editable"`
		BillingAddress              CreateAddressRequest                    `json:"billing_address"`
		BankTransfer                *CreateCheckoutBankTransferRequest      `json:"bank_transfer,omitempty"`
		AcceptedBrands              []string                                `json:"accepted_brands"`
		Pix                         *CreateCheckoutPixPaymentRequest        `json:"pix,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
	c.AcceptedMultiPaymentMethods = temp.AcceptedMultiPaymentMethods
	c.SuccessUrl = temp.SuccessUrl
	c.DefaultPaymentMethod = temp.DefaultPaymentMethod
	c.GatewayAffiliationId = temp.GatewayAffiliationId
	c.CreditCard = temp.CreditCard
	c.DebitCard = temp.DebitCard
	c.Boleto = temp.Boleto
	c.CustomerEditable = temp.CustomerEditable
	c.ExpiresIn = temp.ExpiresIn
	c.SkipCheckoutSuccessPage = temp.SkipCheckoutSuccessPage
	c.BillingAddressEditable = temp.BillingAddressEditable
	c.BillingAddress = temp.BillingAddress
	c.BankTransfer = temp.BankTransfer
	c.AcceptedBrands = temp.AcceptedBrands
	c.Pix = temp.Pix
	return nil
}
