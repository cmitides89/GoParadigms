package employee

import "fmt"

//Displaying inheritance
type employee struct {
	firstName    string
	lastName     string
	totalTimeOff int
	timeOffTaken int
}

//Constructor workaround since go doesnt have constructors
func New(firstName string, lastName string, totalTimeOff int, timeOffTaken int) employee {
	e := employee{firstName, lastName, totalTimeOff, timeOffTaken}
	return e
}

//adding this function to the struct Employee allows for the similar functionality of a class in java
func (e employee) TimeOffRemaining() {
	fmt.Printf("%s %s has %d leaves remaining",
		e.firstName, e.lastName, (e.totalTimeOff - e.timeOffTaken))
}
