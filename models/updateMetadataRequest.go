package models

import (
	"encoding/json"
)

// Request for updating an metadata
type UpdateMetadataRequest struct {
	Metadata map[string]string `json:"metadata"`
}

func (u *UpdateMetadataRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateMetadataRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["metadata"] = u.Metadata
	return structMap
}

func (u *UpdateMetadataRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Metadata map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Metadata = temp.Metadata
	return nil
}
