package app
// This file is where our Handler (Adapters) live

import (
	"encoding/json"
	"encoding/xml"
	"github.com/JeiChambers/banking/service"
	"net/http"
)

type Customer struct {
	Name string		`json:"full_name" xml:"full_name"`
	City string		`json:"city" xml:"city"`
	Zipcode string	`json:"zip_code"  xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Id: "1001", Name: "Chandi", City: "New Delhi", Zipcode: "1110075", DateofBirth: "1987-03-23", Status:"1"},
	//	{Id: "1002", Name: "Allen", City: "Taipei", Zipcode: "114", DateofBirth: "1985-05-23", Status:"1"},
	//	{Id: "1003", Name: "Elon", City: "Austin", Zipcode: "78701", DateofBirth: "1971-06-28", Status:"1"},
	//
	//}

		customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}