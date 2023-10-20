package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting an Address
type GetAddressResponse struct {
	Id           types.Optional[string]              `json:"id"`
	Street       types.Optional[string]              `json:"street"`
	Number       types.Optional[string]              `json:"number"`
	Complement   types.Optional[string]              `json:"complement"`
	ZipCode      types.Optional[string]              `json:"zip_code"`
	Neighborhood types.Optional[string]              `json:"neighborhood"`
	City         types.Optional[string]              `json:"city"`
	State        types.Optional[string]              `json:"state"`
	Country      types.Optional[string]              `json:"country"`
	Status       types.Optional[string]              `json:"status"`
	CreatedAt    types.Optional[time.Time]           `json:"created_at"`
	UpdatedAt    types.Optional[time.Time]           `json:"updated_at"`
	Customer     types.Optional[GetCustomerResponse] `json:"customer"`
	Metadata     types.Optional[map[string]string]   `json:"metadata"`
	Line1        types.Optional[string]              `json:"line_1"`
	Line2        types.Optional[string]              `json:"line_2"`
	DeletedAt    types.Optional[time.Time]           `json:"deleted_at"`
}

func (g *GetAddressResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAddressResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Street.IsValueSet() {
		structMap["street"] = g.Street.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.Complement.IsValueSet() {
		structMap["complement"] = g.Complement.Value()
	}
	if g.ZipCode.IsValueSet() {
		structMap["zip_code"] = g.ZipCode.Value()
	}
	if g.Neighborhood.IsValueSet() {
		structMap["neighborhood"] = g.Neighborhood.Value()
	}
	if g.City.IsValueSet() {
		structMap["city"] = g.City.Value()
	}
	if g.State.IsValueSet() {
		structMap["state"] = g.State.Value()
	}
	if g.Country.IsValueSet() {
		structMap["country"] = g.Country.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Line1.IsValueSet() {
		structMap["line_1"] = g.Line1.Value()
	}
	if g.Line2.IsValueSet() {
		structMap["line_2"] = g.Line2.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	return structMap
}

func (g *GetAddressResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id           types.Optional[string]              `json:"id"`
		Street       types.Optional[string]              `json:"street"`
		Number       types.Optional[string]              `json:"number"`
		Complement   types.Optional[string]              `json:"complement"`
		ZipCode      types.Optional[string]              `json:"zip_code"`
		Neighborhood types.Optional[string]              `json:"neighborhood"`
		City         types.Optional[string]              `json:"city"`
		State        types.Optional[string]              `json:"state"`
		Country      types.Optional[string]              `json:"country"`
		Status       types.Optional[string]              `json:"status"`
		CreatedAt    types.Optional[string]              `json:"created_at"`
		UpdatedAt    types.Optional[string]              `json:"updated_at"`
		Customer     types.Optional[GetCustomerResponse] `json:"customer"`
		Metadata     types.Optional[map[string]string]   `json:"metadata"`
		Line1        types.Optional[string]              `json:"line_1"`
		Line2        types.Optional[string]              `json:"line_2"`
		DeletedAt    types.Optional[string]              `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Street = temp.Street
	g.Number = temp.Number
	g.Complement = temp.Complement
	g.ZipCode = temp.ZipCode
	g.Neighborhood = temp.Neighborhood
	g.City = temp.City
	g.State = temp.State
	g.Country = temp.Country
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.Customer = temp.Customer
	g.Metadata = temp.Metadata
	g.Line1 = temp.Line1
	g.Line2 = temp.Line2
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	return nil
}
