package modle

import (
	"github.com/Ankita-Pohane/GE-employee/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Employee struct {
	gorm.Model
	FirstName   string "gorm:\"\"json:\"firstName\""
	LastName    string "json:\"lastName\""
	EmailId     string "json:\"emailId\""
	Designation string "json:\"designation\""
	EmployeeId  string "json:\"employeeId\""
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Employee{})
}

func (b *Employee) AddEmployee() *Employee {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllEmployees() []Employee {
	var Employees []Employee
	db.Find(&Employees)
	return Employees
}

func GetEmployeeById(EmployeeId int64) (*Employee, *gorm.DB) {
	var getEmployee Employee
	db := db.Where("employee_id=?", EmployeeId).Find(&getEmployee)
	return &getEmployee, db
}

func DeleteEmployee(EmployeeId int64) (*Employee, *gorm.DB) {
	var getEmployee Employee
	db := db.Where("employee_id=?", EmployeeId).Delete(&getEmployee)
	return &getEmployee, db
}
