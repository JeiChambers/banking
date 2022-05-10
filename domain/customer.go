package domain

// Customer Domain|Business Object
type Customer struct {
	Id 			string
	Name 		string
	City 		string
	Zipcode 	string
	DateofBirth string
	Status 		string
}

// CustomerRepository Secondary port
// Remember: interfaces extend functionality to different types
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}




