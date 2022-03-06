package routes

import (
	"github.com/Ankita-Pohane/GE-employee/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterEmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/addEmployee", controllers.AddEmployee).Methods("POST")
	router.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{employeeId}", controllers.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee/{employeeId}", controllers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{employeeId}", controllers.DeleteEmployee).Methods("DELETE")
}
