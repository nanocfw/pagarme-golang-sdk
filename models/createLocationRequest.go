package models

import (
	"encoding/json"
)

// Request for creating a location
type CreateLocationRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *CreateLocationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateLocationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["latitude"] = c.Latitude
	structMap["longitude"] = c.Longitude
	return structMap
}

func (c *CreateLocationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Latitude = temp.Latitude
	c.Longitude = temp.Longitude
	return nil
}
