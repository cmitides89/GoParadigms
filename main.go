package main

import "oopExample/employee"

func main() {
	//calling New() lets us create an instance of the employee structure and
	//store it in a var called e
	e := employee.New("Sam", "Adolf", 30, 20)
	e.TimeOffRemaining()

	// var e employee.Employee
	// e.TimeOffRemaining()

}
