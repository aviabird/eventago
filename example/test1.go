package main

import (
	employee "dummy_bank/dummycqrs/eventago/example/test"
	"fmt"
)

type Employeedetails struct {
	First  string
	Last   string
	Total  int
	Leaves int
}

func main() {
	Employeedetails := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	fmt.Println(Employeedetails)
	Employeedetails.LeavesRemaining()
}
