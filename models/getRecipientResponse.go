package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Recipient response
type GetRecipientResponse struct {
	Id                            types.Optional[string]                           `json:"id"`
	Name                          types.Optional[string]                           `json:"name"`
	Email                         types.Optional[string]                           `json:"email"`
	Document                      types.Optional[string]                           `json:"document"`
	Description                   types.Optional[string]                           `json:"description"`
	Type                          types.Optional[string]                           `json:"type"`
	Status                        types.Optional[string]                           `json:"status"`
	CreatedAt                     types.Optional[time.Time]                        `json:"created_at"`
	UpdatedAt                     types.Optional[time.Time]                        `json:"updated_at"`
	DeletedAt                     types.Optional[time.Time]                        `json:"deleted_at"`
	DefaultBankAccount            types.Optional[GetBankAccountResponse]           `json:"default_bank_account"`
	GatewayRecipients             types.Optional[[]GetGatewayRecipientResponse]    `json:"gateway_recipients"`
	Metadata                      types.Optional[map[string]string]                `json:"metadata"`
	AutomaticAnticipationSettings types.Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
	TransferSettings              types.Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
	Code                          types.Optional[string]                           `json:"code"`
	PaymentMode                   types.Optional[string]                           `json:"payment_mode"`
}

func (g *GetRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Email.IsValueSet() {
		structMap["email"] = g.Email.Value()
	}
	if g.Document.IsValueSet() {
		structMap["document"] = g.Document.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
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
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.DefaultBankAccount.IsValueSet() {
		structMap["default_bank_account"] = g.DefaultBankAccount.Value()
	}
	if g.GatewayRecipients.IsValueSet() {
		structMap["gateway_recipients"] = g.GatewayRecipients.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.AutomaticAnticipationSettings.IsValueSet() {
		structMap["automatic_anticipation_settings"] = g.AutomaticAnticipationSettings.Value()
	}
	if g.TransferSettings.IsValueSet() {
		structMap["transfer_settings"] = g.TransferSettings.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.PaymentMode.IsValueSet() {
		structMap["payment_mode"] = g.PaymentMode.Value()
	}
	return structMap
}

func (g *GetRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                            types.Optional[string]                           `json:"id"`
		Name                          types.Optional[string]                           `json:"name"`
		Email                         types.Optional[string]                           `json:"email"`
		Document                      types.Optional[string]                           `json:"document"`
		Description                   types.Optional[string]                           `json:"description"`
		Type                          types.Optional[string]                           `json:"type"`
		Status                        types.Optional[string]                           `json:"status"`
		CreatedAt                     types.Optional[string]                           `json:"created_at"`
		UpdatedAt                     types.Optional[string]                           `json:"updated_at"`
		DeletedAt                     types.Optional[string]                           `json:"deleted_at"`
		DefaultBankAccount            types.Optional[GetBankAccountResponse]           `json:"default_bank_account"`
		GatewayRecipients             types.Optional[[]GetGatewayRecipientResponse]    `json:"gateway_recipients"`
		Metadata                      types.Optional[map[string]string]                `json:"metadata"`
		AutomaticAnticipationSettings types.Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
		TransferSettings              types.Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
		Code                          types.Optional[string]                           `json:"code"`
		PaymentMode                   types.Optional[string]                           `json:"payment_mode"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Name = temp.Name
	g.Email = temp.Email
	g.Document = temp.Document
	g.Description = temp.Description
	g.Type = temp.Type
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
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.DefaultBankAccount = temp.DefaultBankAccount
	g.GatewayRecipients = temp.GatewayRecipients
	g.Metadata = temp.Metadata
	g.AutomaticAnticipationSettings = temp.AutomaticAnticipationSettings
	g.TransferSettings = temp.TransferSettings
	g.Code = temp.Code
	g.PaymentMode = temp.PaymentMode
	return nil
}
