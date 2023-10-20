package models

import (
	"encoding/json"
)

type UpdateSubscriptionDueDaysRequest struct {
	BoletoDueDays int `json:"boleto_due_days"`
}

func (u *UpdateSubscriptionDueDaysRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionDueDaysRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["boleto_due_days"] = u.BoletoDueDays
	return structMap
}

func (u *UpdateSubscriptionDueDaysRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BoletoDueDays int `json:"boleto_due_days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.BoletoDueDays = temp.BoletoDueDays
	return nil
}
