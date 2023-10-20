package models

import (
	"encoding/json"
)

// Invoice Update Status Request
type UpdateInvoiceStatusRequest struct {
	Status string `json:"status"`
}

func (u *UpdateInvoiceStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateInvoiceStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = u.Status
	return structMap
}

func (u *UpdateInvoiceStatusRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Status string `json:"status"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Status = temp.Status
	return nil
}
