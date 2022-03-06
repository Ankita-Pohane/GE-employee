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
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r)) // used to map system port to docker
}
