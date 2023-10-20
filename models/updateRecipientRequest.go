package models

import (
	"encoding/json"
)

// Request for updating a Recipient
type UpdateRecipientRequest struct {
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Status      string            `json:"status"`
	Metadata    map[string]string `json:"metadata"`
}

func (u *UpdateRecipientRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateRecipientRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = u.Name
	structMap["email"] = u.Email
	structMap["description"] = u.Description
	structMap["type"] = u.Type
	structMap["status"] = u.Status
	structMap["metadata"] = u.Metadata
	return structMap
}

func (u *UpdateRecipientRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name        string            `json:"name"`
		Email       string            `json:"email"`
		Description string            `json:"description"`
		Type        string            `json:"type"`
		Status      string            `json:"status"`
		Metadata    map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Email = temp.Email
	u.Description = temp.Description
	u.Type = temp.Type
	u.Status = temp.Status
	u.Metadata = temp.Metadata
	return nil
}
