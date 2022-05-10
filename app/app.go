package app
//This is the package for wiring hexagonal components

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
	// CustomerHandlers injected service into CustomerHandler. service, our helper function, takes a repository
	// so we give it a helper function to make a new repository stub.
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

