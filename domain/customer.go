package domain

import (
	"github.com/pedramaghasian/hexagonal-in-go/dto"
	"github.com/pedramaghasian/hexagonal-in-go/errs"
)

// Domain Object
type Customer struct {
	Id          string `json:"customer_id" xml:"id" db:"customer_id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipcode" `
	DateOfBirth string `json:"date_of_birth" xml:"birthDate" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// Secondary Port
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
