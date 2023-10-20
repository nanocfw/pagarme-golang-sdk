package models

import (
	"encoding/json"
)

// Request for updating an address
type UpdateAddressRequest struct {
	Number     string            `json:"number"`
	Complement string            `json:"complement"`
	Metadata   map[string]string `json:"metadata"`
	Line2      string            `json:"line_2"`
}

func (u *UpdateAddressRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateAddressRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = u.Number
	structMap["complement"] = u.Complement
	structMap["metadata"] = u.Metadata
	structMap["line_2"] = u.Line2
	return structMap
}

func (u *UpdateAddressRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number     string            `json:"number"`
		Complement string            `json:"complement"`
		Metadata   map[string]string `json:"metadata"`
		Line2      string            `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Number = temp.Number
	u.Complement = temp.Complement
	u.Metadata = temp.Metadata
	u.Line2 = temp.Line2
	return nil
}
