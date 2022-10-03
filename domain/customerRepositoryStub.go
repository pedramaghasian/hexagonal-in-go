package domain

// Adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	var customers = []Customer{
		{Id: "1001", Name: "pedram", City: "tehran", ZipCode: "12", DateOfBirth: "2000-10-09", Status: "1"},
		{Id: "1002", Name: "ali", City: "tehran", ZipCode: "10", DateOfBirth: "1995-12-08", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}

// func (s CustomerRepositoryStub) ById(id string) (*Customer, error) {
// 	var customer *Customer
// 	for _, value := range s.customers {
// 		if value.Id == id {
// 			customer = *value
// 		}
// 	}
// 	return value,nil
// }
