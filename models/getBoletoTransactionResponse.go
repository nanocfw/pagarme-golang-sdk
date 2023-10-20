package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a boleto transaction
type GetBoletoTransactionResponse struct {
	GetTransactionResponse
	Url                 types.Optional[string]                    `json:"url"`
	Barcode             types.Optional[string]                    `json:"barcode"`
	NossoNumero         types.Optional[string]                    `json:"nosso_numero"`
	Bank                types.Optional[string]                    `json:"bank"`
	DocumentNumber      types.Optional[string]                    `json:"document_number"`
	Instructions        types.Optional[string]                    `json:"instructions"`
	BillingAddress      types.Optional[GetBillingAddressResponse] `json:"billing_address"`
	DueAt               types.Optional[time.Time]                 `json:"due_at"`
	QrCode              types.Optional[string]                    `json:"qr_code"`
	Line                types.Optional[string]                    `json:"line"`
	PdfPassword         types.Optional[string]                    `json:"pdf_password"`
	Pdf                 types.Optional[string]                    `json:"pdf"`
	PaidAt              types.Optional[time.Time]                 `json:"paid_at"`
	PaidAmount          types.Optional[string]                    `json:"paid_amount"`
	Type                types.Optional[string]                    `json:"type"`
	CreditAt            types.Optional[time.Time]                 `json:"credit_at"`
	StatementDescriptor types.Optional[string]                    `json:"statement_descriptor"`
}

func (g *GetBoletoTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetBoletoTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.TransactionType != nil {
		structMap["transaction_type"] = *g.TransactionType
	} else {
		structMap["transaction_type"] = "boleto"
	}
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.Barcode.IsValueSet() {
		structMap["barcode"] = g.Barcode.Value()
	}
	if g.NossoNumero.IsValueSet() {
		structMap["nosso_numero"] = g.NossoNumero.Value()
	}
	if g.Bank.IsValueSet() {
		structMap["bank"] = g.Bank.Value()
	}
	if g.DocumentNumber.IsValueSet() {
		structMap["document_number"] = g.DocumentNumber.Value()
	}
	if g.Instructions.IsValueSet() {
		structMap["instructions"] = g.Instructions.Value()
	}
	if g.BillingAddress.IsValueSet() {
		structMap["billing_address"] = g.BillingAddress.Value()
	}
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if g.QrCode.IsValueSet() {
		structMap["qr_code"] = g.QrCode.Value()
	}
	if g.Line.IsValueSet() {
		structMap["line"] = g.Line.Value()
	}
	if g.PdfPassword.IsValueSet() {
		structMap["pdf_password"] = g.PdfPassword.Value()
	}
	if g.Pdf.IsValueSet() {
		structMap["pdf"] = g.Pdf.Value()
	}
	if g.PaidAt.IsValueSet() {
		var PaidAtVal *string = nil
		if g.PaidAt.Value() != nil {
			val := g.PaidAt.Value().Format(time.RFC3339)
			PaidAtVal = &val
		}
		structMap["paid_at"] = PaidAtVal
	}
	if g.PaidAmount.IsValueSet() {
		structMap["paid_amount"] = g.PaidAmount.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.CreditAt.IsValueSet() {
		var CreditAtVal *string = nil
		if g.CreditAt.Value() != nil {
			val := g.CreditAt.Value().Format(time.RFC3339)
			CreditAtVal = &val
		}
		structMap["credit_at"] = CreditAtVal
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	return structMap
}

func (g *GetBoletoTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Url                 types.Optional[string]                    `json:"url"`
		Barcode             types.Optional[string]                    `json:"barcode"`
		NossoNumero         types.Optional[string]                    `json:"nosso_numero"`
		Bank                types.Optional[string]                    `json:"bank"`
		DocumentNumber      types.Optional[string]                    `json:"document_number"`
		Instructions        types.Optional[string]                    `json:"instructions"`
		BillingAddress      types.Optional[GetBillingAddressResponse] `json:"billing_address"`
		DueAt               types.Optional[string]                    `json:"due_at"`
		QrCode              types.Optional[string]                    `json:"qr_code"`
		Line                types.Optional[string]                    `json:"line"`
		PdfPassword         types.Optional[string]                    `json:"pdf_password"`
		Pdf                 types.Optional[string]                    `json:"pdf"`
		PaidAt              types.Optional[string]                    `json:"paid_at"`
		PaidAmount          types.Optional[string]                    `json:"paid_amount"`
		Type                types.Optional[string]                    `json:"type"`
		CreditAt            types.Optional[string]                    `json:"credit_at"`
		StatementDescriptor types.Optional[string]                    `json:"statement_descriptor"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.Barcode = temp.Barcode
	g.NossoNumero = temp.NossoNumero
	g.Bank = temp.Bank
	g.DocumentNumber = temp.DocumentNumber
	g.Instructions = temp.Instructions
	g.BillingAddress = temp.BillingAddress
	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
	g.QrCode = temp.QrCode
	g.Line = temp.Line
	g.PdfPassword = temp.PdfPassword
	g.Pdf = temp.Pdf
	g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
	if temp.PaidAt.Value() != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt.SetValue(&PaidAtVal)
	}
	g.PaidAmount = temp.PaidAmount
	g.Type = temp.Type
	g.CreditAt.ShouldSetValue(temp.CreditAt.IsValueSet())
	if temp.CreditAt.Value() != nil {
		CreditAtVal, err := time.Parse(time.RFC3339, (*temp.CreditAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse credit_at as % s format.", time.RFC3339)
		}
		g.CreditAt.SetValue(&CreditAtVal)
	}
	g.StatementDescriptor = temp.StatementDescriptor
	return nil
}

// Getter for url.
func (g *GetBoletoTransactionResponse) GetUrl() types.Optional[string] {
	return g.Url
}

// Getter for barcode.
func (g *GetBoletoTransactionResponse) GetBarcode() types.Optional[string] {
	return g.Barcode
}

// Getter for nosso_numero.
func (g *GetBoletoTransactionResponse) GetNossoNumero() types.Optional[string] {
	return g.NossoNumero
}

// Getter for bank.
func (g *GetBoletoTransactionResponse) GetBank() types.Optional[string] {
	return g.Bank
}

// Getter for document_number.
func (g *GetBoletoTransactionResponse) GetDocumentNumber() types.Optional[string] {
	return g.DocumentNumber
}

// Getter for instructions.
func (g *GetBoletoTransactionResponse) GetInstructions() types.Optional[string] {
	return g.Instructions
}

// Getter for billing_address.
func (g *GetBoletoTransactionResponse) GetBillingAddress() types.Optional[GetBillingAddressResponse] {
	return g.BillingAddress
}

// Getter for due_at.
func (g *GetBoletoTransactionResponse) GetDueAt() types.Optional[time.Time] {
	return g.DueAt
}

// Getter for qr_code.
func (g *GetBoletoTransactionResponse) GetQrCode() types.Optional[string] {
	return g.QrCode
}

// Getter for line.
func (g *GetBoletoTransactionResponse) GetLine() types.Optional[string] {
	return g.Line
}

// Getter for pdf_password.
func (g *GetBoletoTransactionResponse) GetPdfPassword() types.Optional[string] {
	return g.PdfPassword
}

// Getter for pdf.
func (g *GetBoletoTransactionResponse) GetPdf() types.Optional[string] {
	return g.Pdf
}

// Getter for paid_at.
func (g *GetBoletoTransactionResponse) GetPaidAt() types.Optional[time.Time] {
	return g.PaidAt
}

// Getter for paid_amount.
func (g *GetBoletoTransactionResponse) GetPaidAmount() types.Optional[string] {
	return g.PaidAmount
}

// Getter for type.
func (g *GetBoletoTransactionResponse) GetType() types.Optional[string] {
	return g.Type
}

// Getter for credit_at.
func (g *GetBoletoTransactionResponse) GetCreditAt() types.Optional[time.Time] {
	return g.CreditAt
}

// Getter for statement_descriptor.
func (g *GetBoletoTransactionResponse) GetStatementDescriptor() types.Optional[string] {
	return g.StatementDescriptor
}

func UnmarshalGetBoletoTransactionResponse(input []byte) (
	GetBoletoTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
	if err != nil {
		return nil, err
	}

	if getBoletoTransactionResponse, ok := getTransactionResponse.(GetBoletoTransactionResponseInterface); ok {
		return getBoletoTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getBoletoTransactionResponse %v", getTransactionResponse)
}

func ToGetBoletoTransactionResponseArray(getBoletoTransactionResponse []GetBoletoTransactionResponseField) []GetBoletoTransactionResponseInterface {
	var items []GetBoletoTransactionResponseInterface
	for _, temp := range getBoletoTransactionResponse {
		items = append(items, temp.GetBoletoTransactionResponseInterface)
	}
	return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetBoletoTransactionResponseField struct {
	GetBoletoTransactionResponseInterface
}

func (g *GetBoletoTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetBoletoTransactionResponseInterface, err = UnmarshalGetBoletoTransactionResponse(input)
	return err
}

// Utility struct that helps in the unmarshalling of slices.
type GetBoletoTransactionResponseSlice struct {
	Slice []GetBoletoTransactionResponseInterface
}

func (ga *GetBoletoTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetBoletoTransactionResponseInterface{}
		for _, getBoletoTransactionResponse := range temp {
			if getBoletoTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetBoletoTransactionResponse
			err := json.Unmarshal(getBoletoTransactionResponse, &obj)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, &obj)
		}
	}
	return nil
}
