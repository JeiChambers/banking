package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{Id: "1001", Name: "Chandi", City: "New Delhi", Zipcode: "1110075", DateofBirth: "1987-03-23", Status:"1"},
		{Id: "1002", Name: "Allen", City: "Taipei", Zipcode: "114", DateofBirth: "1985-05-23", Status:"1"},
		{Id: "1003", Name: "Elon", City: "Austin", Zipcode: "78701", DateofBirth: "1971-06-28", Status:"1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
