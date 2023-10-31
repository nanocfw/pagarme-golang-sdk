package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Request for creating a new Address
type CreateAddressRequest struct {
	Street       string                            `json:"street"`
	Number       string                            `json:"number"`
	ZipCode      string                            `json:"zip_code"`
	Neighborhood string                            `json:"neighborhood"`
	City         string                            `json:"city"`
	State        string                            `json:"state"`
	Country      string                            `json:"country"`
	Complement   string                            `json:"complement,omitempty"`
	Metadata     types.Optional[map[string]string] `json:"metadata,omitempty"`
	Line1        string                            `json:"line_1,omitempty"`
	Line2        string                            `json:"line_2,omitempty"`
}

func (c *CreateAddressRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateAddressRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["street"] = c.Street
	structMap["number"] = c.Number
	structMap["zip_code"] = c.ZipCode
	structMap["neighborhood"] = c.Neighborhood
	structMap["city"] = c.City
	structMap["state"] = c.State
	structMap["country"] = c.Country
	structMap["complement"] = c.Complement
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	structMap["line_1"] = c.Line1
	structMap["line_2"] = c.Line2
	return structMap
}

func (c *CreateAddressRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Street       string                            `json:"street"`
		Number       string                            `json:"number"`
		ZipCode      string                            `json:"zip_code"`
		Neighborhood string                            `json:"neighborhood"`
		City         string                            `json:"city"`
		State        string                            `json:"state"`
		Country      string                            `json:"country"`
		Complement   string                            `json:"complement"`
		Metadata     types.Optional[map[string]string] `json:"metadata"`
		Line1        string                            `json:"line_1"`
		Line2        string                            `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Street = temp.Street
	c.Number = temp.Number
	c.ZipCode = temp.ZipCode
	c.Neighborhood = temp.Neighborhood
	c.City = temp.City
	c.State = temp.State
	c.Country = temp.Country
	c.Complement = temp.Complement
	c.Metadata = temp.Metadata
	c.Line1 = temp.Line1
	c.Line2 = temp.Line2
	return nil
}
