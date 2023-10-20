package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a charge
type GetChargeResponse struct {
	Id                  types.Optional[string]                          `json:"id"`
	Code                types.Optional[string]                          `json:"code"`
	GatewayId           types.Optional[string]                          `json:"gateway_id"`
	Amount              types.Optional[int]                             `json:"amount"`
	Status              types.Optional[string]                          `json:"status"`
	Currency            types.Optional[string]                          `json:"currency"`
	PaymentMethod       types.Optional[string]                          `json:"payment_method"`
	DueAt               types.Optional[time.Time]                       `json:"due_at"`
	CreatedAt           types.Optional[time.Time]                       `json:"created_at"`
	UpdatedAt           types.Optional[time.Time]                       `json:"updated_at"`
	LastTransaction     types.Optional[GetTransactionResponseInterface] `json:"last_transaction"`
	Invoice             types.Optional[GetInvoiceResponse]              `json:"invoice"`
	Order               types.Optional[GetOrderResponse]                `json:"order"`
	Customer            types.Optional[GetCustomerResponse]             `json:"customer"`
	Metadata            types.Optional[map[string]string]               `json:"metadata"`
	PaidAt              types.Optional[time.Time]                       `json:"paid_at"`
	CanceledAt          types.Optional[time.Time]                       `json:"canceled_at"`
	CanceledAmount      types.Optional[int]                             `json:"canceled_amount"`
	PaidAmount          types.Optional[int]                             `json:"paid_amount"`
	InterestAndFinePaid types.Optional[int]                             `json:"interest_and_fine_paid"`
	RecurrencyCycle     types.Optional[string]                          `json:"recurrency_cycle"`
}

func (g *GetChargeResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetChargeResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.PaymentMethod.IsValueSet() {
		structMap["payment_method"] = g.PaymentMethod.Value()
	}
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
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
	if g.LastTransaction.IsValueSet() {
		structMap["last_transaction"] = g.LastTransaction.Value()
	}
	if g.Invoice.IsValueSet() {
		structMap["invoice"] = g.Invoice.Value()
	}
	if g.Order.IsValueSet() {
		structMap["order"] = g.Order.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.PaidAt.IsValueSet() {
		var PaidAtVal *string = nil
		if g.PaidAt.Value() != nil {
			val := g.PaidAt.Value().Format(time.RFC3339)
			PaidAtVal = &val
		}
		structMap["paid_at"] = PaidAtVal
	}
	if g.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if g.CanceledAt.Value() != nil {
			val := g.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if g.CanceledAmount.IsValueSet() {
		structMap["canceled_amount"] = g.CanceledAmount.Value()
	}
	if g.PaidAmount.IsValueSet() {
		structMap["paid_amount"] = g.PaidAmount.Value()
	}
	if g.InterestAndFinePaid.IsValueSet() {
		structMap["interest_and_fine_paid"] = g.InterestAndFinePaid.Value()
	}
	if g.RecurrencyCycle.IsValueSet() {
		structMap["recurrency_cycle"] = g.RecurrencyCycle.Value()
	}
	return structMap
}

func (g *GetChargeResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                  types.Optional[string]                          `json:"id"`
		Code                types.Optional[string]                          `json:"code"`
		GatewayId           types.Optional[string]                          `json:"gateway_id"`
		Amount              types.Optional[int]                             `json:"amount"`
		Status              types.Optional[string]                          `json:"status"`
		Currency            types.Optional[string]                          `json:"currency"`
		PaymentMethod       types.Optional[string]                          `json:"payment_method"`
		DueAt               types.Optional[string]                          `json:"due_at"`
		CreatedAt           types.Optional[string]                          `json:"created_at"`
		UpdatedAt           types.Optional[string]                          `json:"updated_at"`
		LastTransaction     types.Optional[GetTransactionResponseInterface] `json:"last_transaction"`
		Invoice             types.Optional[GetInvoiceResponse]              `json:"invoice"`
		Order               types.Optional[GetOrderResponse]                `json:"order"`
		Customer            types.Optional[GetCustomerResponse]             `json:"customer"`
		Metadata            types.Optional[map[string]string]               `json:"metadata"`
		PaidAt              types.Optional[string]                          `json:"paid_at"`
		CanceledAt          types.Optional[string]                          `json:"canceled_at"`
		CanceledAmount      types.Optional[int]                             `json:"canceled_amount"`
		PaidAmount          types.Optional[int]                             `json:"paid_amount"`
		InterestAndFinePaid types.Optional[int]                             `json:"interest_and_fine_paid"`
		RecurrencyCycle     types.Optional[string]                          `json:"recurrency_cycle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.GatewayId = temp.GatewayId
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.Currency = temp.Currency
	g.PaymentMethod = temp.PaymentMethod
	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
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
	g.LastTransaction = temp.LastTransaction
	g.Invoice = temp.Invoice
	g.Order = temp.Order
	g.Customer = temp.Customer
	g.Metadata = temp.Metadata
	g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
	if temp.PaidAt.Value() != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt.SetValue(&PaidAtVal)
	}
	g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		g.CanceledAt.SetValue(&CanceledAtVal)
	}
	g.CanceledAmount = temp.CanceledAmount
	g.PaidAmount = temp.PaidAmount
	g.InterestAndFinePaid = temp.InterestAndFinePaid
	g.RecurrencyCycle = temp.RecurrencyCycle
	return nil
}
