package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary int
	EOY    int
}

func main() {
	employeeDetails := Employee{"Avishek", 24, 46000, 1}
	employeeDetails.Salary = 46400
	fmt.Println(employeeDetails)
}
