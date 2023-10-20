package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for updating the start date from a subscription
type UpdateSubscriptionStartAtRequest struct {
	StartAt time.Time `json:"start_at"`
}

func (u *UpdateSubscriptionStartAtRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionStartAtRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_at"] = u.StartAt.Format(time.RFC3339)
	return structMap
}

func (u *UpdateSubscriptionStartAtRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartAt string `json:"start_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	StartAtVal, err := time.Parse(time.RFC3339, temp.StartAt)
	if err != nil {
		log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
	}
	u.StartAt = StartAtVal
	return nil
}
