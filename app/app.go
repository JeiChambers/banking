package app

import (
	"github.com/JeiChambers/banking/domain"
	"github.com/JeiChambers/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

