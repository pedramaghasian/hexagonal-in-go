package dto

type CustomerResponse struct {
	Id          string `json:"customer_id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipcode" `
	DateOfBirth string `json:"date_of_birth" xml:"birthDate"`
	Status      string `json:"status" xml:"status"`
}
