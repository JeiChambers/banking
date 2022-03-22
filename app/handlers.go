package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func (ch * CustomerHandlers)getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "Chandi", City: "New Delhi", Zipcode: "1110075"},
	//	{Name: "Allen", City: "Taipei", Zipcode: "114"},
	//	{Name: "Elon", City: "Austin", Zipcode: "78701"},
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