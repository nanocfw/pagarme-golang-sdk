package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request to update the end date of the current subscription cycle
type UpdateCurrentCycleEndDateRequest struct {
	EndAt *time.Time `json:"end_at,omitempty"`
}

func (u *UpdateCurrentCycleEndDateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateCurrentCycleEndDateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.EndAt != nil {
		structMap["end_at"] = u.EndAt.Format(time.RFC3339)
	}
	return structMap
}

func (u *UpdateCurrentCycleEndDateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EndAt *string `json:"end_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.EndAt != nil {
		EndAtVal, err := time.Parse(time.RFC3339, *temp.EndAt)
		if err != nil {
			log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
		}
		u.EndAt = &EndAtVal
	}
	return nil
}
