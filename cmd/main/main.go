package main

import (
	"log"
	"net/http"

	"github.com/Ankita-Pohane/GE-employee/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
