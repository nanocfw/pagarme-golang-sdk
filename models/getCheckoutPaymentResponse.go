package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Resposta das configurações de pagamento do checkout
type GetCheckoutPaymentResponse struct {
	Id                      types.Optional[string]                                 `json:"id"`
	Amount                  types.Optional[int]                                    `json:"amount"`
	DefaultPaymentMethod    types.Optional[string]                                 `json:"default_payment_method"`
	SuccessUrl              types.Optional[string]                                 `json:"success_url"`
	PaymentUrl              types.Optional[string]                                 `json:"payment_url"`
	GatewayAffiliationId    types.Optional[string]                                 `json:"gateway_affiliation_id"`
	AcceptedPaymentMethods  types.Optional[[]string]                               `json:"accepted_payment_methods"`
	Status                  types.Optional[string]                                 `json:"status"`
	SkipCheckoutSuccessPage types.Optional[bool]                                   `json:"skip_checkout_success_page"`
	CreatedAt               types.Optional[time.Time]                              `json:"created_at"`
	UpdatedAt               types.Optional[time.Time]                              `json:"updated_at"`
	CanceledAt              types.Optional[time.Time]                              `json:"canceled_at"`
	CustomerEditable        types.Optional[bool]                                   `json:"customer_editable"`
	Customer                types.Optional[GetCustomerResponse]                    `json:"customer"`
	Billingaddress          types.Optional[GetAddressResponse]                     `json:"billingaddress"`
	CreditCard              types.Optional[GetCheckoutCreditCardPaymentResponse]   `json:"credit_card"`
	Boleto                  types.Optional[GetCheckoutBoletoPaymentResponse]       `json:"boleto"`
	BillingAddressEditable  types.Optional[bool]                                   `json:"billing_address_editable"`
	Shipping                types.Optional[GetShippingResponse]                    `json:"shipping"`
	Shippable               types.Optional[bool]                                   `json:"shippable"`
	ClosedAt                types.Optional[time.Time]                              `json:"closed_at"`
	ExpiresAt               types.Optional[time.Time]                              `json:"expires_at"`
	Currency                types.Optional[string]                                 `json:"currency"`
	DebitCard               types.Optional[GetCheckoutDebitCardPaymentResponse]    `json:"debit_card"`
	BankTransfer            types.Optional[GetCheckoutBankTransferPaymentResponse] `json:"bank_transfer"`
	AcceptedBrands          types.Optional[[]string]                               `json:"accepted_brands"`
	Pix                     types.Optional[GetCheckoutPixPaymentResponse]          `json:"pix"`
}

func (g *GetCheckoutPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.DefaultPaymentMethod.IsValueSet() {
		structMap["default_payment_method"] = g.DefaultPaymentMethod.Value()
	}
	if g.SuccessUrl.IsValueSet() {
		structMap["success_url"] = g.SuccessUrl.Value()
	}
	if g.PaymentUrl.IsValueSet() {
		structMap["payment_url"] = g.PaymentUrl.Value()
	}
	if g.GatewayAffiliationId.IsValueSet() {
		structMap["gateway_affiliation_id"] = g.GatewayAffiliationId.Value()
	}
	if g.AcceptedPaymentMethods.IsValueSet() {
		structMap["accepted_payment_methods"] = g.AcceptedPaymentMethods.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.SkipCheckoutSuccessPage.IsValueSet() {
		structMap["skip_checkout_success_page"] = g.SkipCheckoutSuccessPage.Value()
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
	if g.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if g.CanceledAt.Value() != nil {
			val := g.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if g.CustomerEditable.IsValueSet() {
		structMap["customer_editable"] = g.CustomerEditable.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Billingaddress.IsValueSet() {
		structMap["billingaddress"] = g.Billingaddress.Value()
	}
	if g.CreditCard.IsValueSet() {
		structMap["credit_card"] = g.CreditCard.Value()
	}
	if g.Boleto.IsValueSet() {
		structMap["boleto"] = g.Boleto.Value()
	}
	if g.BillingAddressEditable.IsValueSet() {
		structMap["billing_address_editable"] = g.BillingAddressEditable.Value()
	}
	if g.Shipping.IsValueSet() {
		structMap["shipping"] = g.Shipping.Value()
	}
	if g.Shippable.IsValueSet() {
		structMap["shippable"] = g.Shippable.Value()
	}
	if g.ClosedAt.IsValueSet() {
		var ClosedAtVal *string = nil
		if g.ClosedAt.Value() != nil {
			val := g.ClosedAt.Value().Format(time.RFC3339)
			ClosedAtVal = &val
		}
		structMap["closed_at"] = ClosedAtVal
	}
	if g.ExpiresAt.IsValueSet() {
		var ExpiresAtVal *string = nil
		if g.ExpiresAt.Value() != nil {
			val := g.ExpiresAt.Value().Format(time.RFC3339)
			ExpiresAtVal = &val
		}
		structMap["expires_at"] = ExpiresAtVal
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.DebitCard.IsValueSet() {
		structMap["debit_card"] = g.DebitCard.Value()
	}
	if g.BankTransfer.IsValueSet() {
		structMap["bank_transfer"] = g.BankTransfer.Value()
	}
	if g.AcceptedBrands.IsValueSet() {
		structMap["accepted_brands"] = g.AcceptedBrands.Value()
	}
	if g.Pix.IsValueSet() {
		structMap["pix"] = g.Pix.Value()
	}
	return structMap
}

func (g *GetCheckoutPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                      types.Optional[string]                                 `json:"id"`
		Amount                  types.Optional[int]                                    `json:"amount"`
		DefaultPaymentMethod    types.Optional[string]                                 `json:"default_payment_method"`
		SuccessUrl              types.Optional[string]                                 `json:"success_url"`
		PaymentUrl              types.Optional[string]                                 `json:"payment_url"`
		GatewayAffiliationId    types.Optional[string]                                 `json:"gateway_affiliation_id"`
		AcceptedPaymentMethods  types.Optional[[]string]                               `json:"accepted_payment_methods"`
		Status                  types.Optional[string]                                 `json:"status"`
		SkipCheckoutSuccessPage types.Optional[bool]                                   `json:"skip_checkout_success_page"`
		CreatedAt               types.Optional[string]                                 `json:"created_at"`
		UpdatedAt               types.Optional[string]                                 `json:"updated_at"`
		CanceledAt              types.Optional[string]                                 `json:"canceled_at"`
		CustomerEditable        types.Optional[bool]                                   `json:"customer_editable"`
		Customer                types.Optional[GetCustomerResponse]                    `json:"customer"`
		Billingaddress          types.Optional[GetAddressResponse]                     `json:"billingaddress"`
		CreditCard              types.Optional[GetCheckoutCreditCardPaymentResponse]   `json:"credit_card"`
		Boleto                  types.Optional[GetCheckoutBoletoPaymentResponse]       `json:"boleto"`
		BillingAddressEditable  types.Optional[bool]                                   `json:"billing_address_editable"`
		Shipping                types.Optional[GetShippingResponse]                    `json:"shipping"`
		Shippable               types.Optional[bool]                                   `json:"shippable"`
		ClosedAt                types.Optional[string]                                 `json:"closed_at"`
		ExpiresAt               types.Optional[string]                                 `json:"expires_at"`
		Currency                types.Optional[string]                                 `json:"currency"`
		DebitCard               types.Optional[GetCheckoutDebitCardPaymentResponse]    `json:"debit_card"`
		BankTransfer            types.Optional[GetCheckoutBankTransferPaymentResponse] `json:"bank_transfer"`
		AcceptedBrands          types.Optional[[]string]                               `json:"accepted_brands"`
		Pix                     types.Optional[GetCheckoutPixPaymentResponse]          `json:"pix"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Amount = temp.Amount
	g.DefaultPaymentMethod = temp.DefaultPaymentMethod
	g.SuccessUrl = temp.SuccessUrl
	g.PaymentUrl = temp.PaymentUrl
	g.GatewayAffiliationId = temp.GatewayAffiliationId
	g.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
	g.Status = temp.Status
	g.SkipCheckoutSuccessPage = temp.SkipCheckoutSuccessPage
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
	g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		g.CanceledAt.SetValue(&CanceledAtVal)
	}
	g.CustomerEditable = temp.CustomerEditable
	g.Customer = temp.Customer
	g.Billingaddress = temp.Billingaddress
	g.CreditCard = temp.CreditCard
	g.Boleto = temp.Boleto
	g.BillingAddressEditable = temp.BillingAddressEditable
	g.Shipping = temp.Shipping
	g.Shippable = temp.Shippable
	g.ClosedAt.ShouldSetValue(temp.ClosedAt.IsValueSet())
	if temp.ClosedAt.Value() != nil {
		ClosedAtVal, err := time.Parse(time.RFC3339, (*temp.ClosedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse closed_at as % s format.", time.RFC3339)
		}
		g.ClosedAt.SetValue(&ClosedAtVal)
	}
	g.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
	if temp.ExpiresAt.Value() != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		g.ExpiresAt.SetValue(&ExpiresAtVal)
	}
	g.Currency = temp.Currency
	g.DebitCard = temp.DebitCard
	g.BankTransfer = temp.BankTransfer
	g.AcceptedBrands = temp.AcceptedBrands
	g.Pix = temp.Pix
	return nil
}
