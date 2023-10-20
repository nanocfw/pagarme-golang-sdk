package models

import (
	"encoding/json"
)

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

func (u *UpdateOrderStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateOrderStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = u.Status
	return structMap
}

func (u *UpdateOrderStatusRequest) UnmarshalJSON(input []byte) error {
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
