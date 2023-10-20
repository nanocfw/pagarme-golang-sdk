package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Checkout pix payment response
type GetCheckoutPixPaymentResponse struct {
	ExpiresAt             types.Optional[time.Time]                  `json:"expires_at"`
	AdditionalInformation types.Optional[[]PixAdditionalInformation] `json:"additional_information"`
}

func (g *GetCheckoutPixPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutPixPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.ExpiresAt.IsValueSet() {
		var ExpiresAtVal *string = nil
		if g.ExpiresAt.Value() != nil {
			val := g.ExpiresAt.Value().Format(time.RFC3339)
			ExpiresAtVal = &val
		}
		structMap["expires_at"] = ExpiresAtVal
	}
	if g.AdditionalInformation.IsValueSet() {
		structMap["additional_information"] = g.AdditionalInformation.Value()
	}
	return structMap
}

func (g *GetCheckoutPixPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresAt             types.Optional[string]                     `json:"expires_at"`
		AdditionalInformation types.Optional[[]PixAdditionalInformation] `json:"additional_information"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
	if temp.ExpiresAt.Value() != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		g.ExpiresAt.SetValue(&ExpiresAtVal)
	}
	g.AdditionalInformation = temp.AdditionalInformation
	return nil
}
