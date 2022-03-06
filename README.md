# GE-employee
Demo for crud operation using Go lang and dockerize application

# Command to run
Run below comand to run Employee record application

```
docker-compose up

```

# Request for using API's

# Get all employees
GET API - to get all employees data

URL - http://localhost:9010/employees

# Add an employee
POST API - to add an employee

URL - http://localhost:9010/addEmployee

Request - 

```
{
    "FirstName": "Ankita",
    "LastName": "Pohane",
    "EmailId": "Ankita.Pohane@quest-global.com",
    "Designation": "Software Developer",
    "EmployeeId":"1003"
}
```

# Get by employee Id
GET API - to get any employee data by employee id

URL - http://localhost:9010/employee/1003

# Update employee data by using employee id
PUT API -  to update employee data by employee id

URL - http://localhost:9010/employee/1001

Request - 

```
{
    "FirstName": "Ankita New",
    "LastName": "Pohane",
    "EmailId": "Ankita.Pohane@quest-global.com",
    "Designation": "Software Developer",
    "EmployeeId":"1003"
}
```

# Delete employee data using employee id
DELETE API - to delete employee by employee id

URL - http://localhost:9010/employee/1003




