package main

import (
	"fmt"
)

// PREMIUM - maximum value of premium. Employee can get part of it if norm
// will be overfilled. (1-20 hours)
const PREMIUM = 2500

// Employee - structure, that have information about working time of employees
// normalHours (nH) - normal count of working hours per month
// actualHours (aH) - actual count of working hours per month
// coefficient (c) - coefficient that will be multiplied to PREMIUM (0.0-1.0)
type employee struct {
	id     int
	nH, aH int
	c      float64
}

// prints greeting string
func greeting() {
	fmt.Println("Hello world!")
}

func main() {

	// printing "Hello world!"
	greeting()

	// creating variable with type Employee
	emp := employee{id: 1, nH: 170, aH: 171}
	// condition for cases
	cond := (emp.aH - emp.nH)

	// checking overfulfilling
	switch {
	// overfilled
	case cond > 0:
		emp.c = float64(cond) / 20.0
		if emp.c > 1 {
			emp.c = 1.0
		}
		fmt.Println("Employee overfulfilled the norm and deserve the premium!")
	// fulfilled
	case cond == 0:
		fmt.Println("Employee has exactly fulfilled the norm!")
	// didn't fulfilled
	default:
		fmt.Println("Employee should work harder next month! The norm hasn't been fulfilled!")
	}
	fmt.Println("Employee's premium: ", (PREMIUM * emp.c))

	// creating two another employees by differ methods
	emp2 := employee{2, 160, 120, 0.75}
	emp3 := employee{}
	fmt.Println(emp, emp2, emp3)
}
