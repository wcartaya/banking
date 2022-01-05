package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//router := http.NewServeMux()
	router := mux.NewRouter()
	//routers
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
