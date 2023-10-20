package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for updating a charge due date
type UpdateChargeDueDateRequest struct {
	DueAt *time.Time `json:"due_at,omitempty"`
}

func (u *UpdateChargeDueDateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateChargeDueDateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.DueAt != nil {
		structMap["due_at"] = u.DueAt.Format(time.RFC3339)
	}
	return structMap
}

func (u *UpdateChargeDueDateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		DueAt *string `json:"due_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.DueAt != nil {
		DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		u.DueAt = &DueAtVal
	}
	return nil
}
