package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	modle "github.com/Ankita-Pohane/GE-employee/pkg/models"
	"github.com/Ankita-Pohane/GE-employee/pkg/utils"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	allEmployees := modle.GetAllEmployees()
	res, _ := json.Marshal(allEmployees)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	EmployeeId := vars["employeeId"]
	fmt.Println("employeeId: ", EmployeeId)
	ID, err := strconv.ParseInt(EmployeeId, 0, 0)
	fmt.Println("Id to get: ", ID)
	if err != nil {
		fmt.Println("error while parsing")
	}
	employeeId, _ := modle.GetEmployeeById(ID)
	res, _ := json.Marshal(employeeId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var employee modle.Employee
	err := decoder.Decode(&employee)
	if err != nil {
		panic(err)
	}
	b := employee.AddEmployee()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	ID, err := strconv.ParseInt(employeeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	employee, _ := modle.DeleteEmployee(ID)
	res, _ := json.Marshal(employee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var updateEmployee = &modle.Employee{}
	utils.ParseBody(r, updateEmployee)
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	ID, err := strconv.ParseInt(employeeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	employeeDetails, db := modle.GetEmployeeById(ID)
	if updateEmployee.FirstName != "" {
		employeeDetails.FirstName = updateEmployee.FirstName
	}
	if updateEmployee.LastName != "" {
		employeeDetails.LastName = updateEmployee.LastName
	}
	if updateEmployee.EmailId != "" {
		employeeDetails.EmailId = updateEmployee.EmailId
	}
	if updateEmployee.Designation != "" {
		employeeDetails.Designation = updateEmployee.Designation
	}
	db.Save(&employeeDetails)
	res, _ := json.Marshal(employeeDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
