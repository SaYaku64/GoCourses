package main

// Создать функцию которая приводит работников к человеку(значение типов)
// Создать кеш начальников и работников не приступать опросу ( перебору) работников пока не будут перебраны все начальники.
// Перебирать кеши в рутинах, конкурентно. С RWMutex.

import (
	"fmt"
	"time"
)

var mux = myMutex{}

// Func that shows human from the employee
func empToHum(e employee) human {
	return e.info
}

// func readHead(m map[int]head) {
// 	fmt.Println("Checking heads:")
// 	for i := range m {
// 		fmt.Println(m[i])
// 	}
// }

func readEmp(m map[int]employee) {
	// b := false
	// for b == true {
	// 	if m[0].experience == 222 {
	// 		b = true
	// 	}
	// }
	fmt.Println(m[0].head)
	if m[0].head == false {
		time.Sleep(3 * time.Second)
		mux.mutex.Lock()
		for i := range m {
			fmt.Println(m[i])
		}
		mux.mutex.Unlock()
	} else {
		mux.mutex.Lock()
		for i := range m {
			fmt.Println(m[i])
		}
		mux.mutex.Unlock()
	}

}

func main() {

	var head1 = employee{
		jobName:    "Head of Human Resources Dep",
		experience: 8.5,
		head:       true,
		info: human{
			birthday: "26.11.1982",
			name:     "Michael",
			surname:  "McCalisto",
		},
	}

	var head2 = employee{
		jobName:    "Head of Connections Dep",
		experience: 10.25,
		head:       true,
		info: human{
			birthday: "14.10.1999",
			name:     "Tom",
			surname:  "Angelo",
		},
	}

	var emp1 = employee{
		jobName:    "Actor",
		experience: 7.45,
		info: human{
			birthday: "06.04.2000",
			name:     "Alex",
			surname:  "Cake",
		},
	}

	var emp2 = employee{
		jobName:    "Cleaner",
		experience: 2,
		info: human{
			birthday: "01.01.1995",
			name:     "Dmitriy",
			surname:  "Klyak",
		},
	}

	//fmt.Println("Human-employee-1:", empToHum(emp1))

	heads := make(map[int]employee)
	{
		heads[1] = head1
		heads[2] = head2
	}

	emps := make(map[int]employee)
	{
		emps[1] = emp1
		emps[2] = emp2
	}

	go readEmp(heads)
	go readEmp(emps)

	var input string
	fmt.Scanln(&input)
}

// Создать функцию которая приводит работников к человеку(значение типов)
// Создать кеш начальников и работников не приступать опросу ( перебору) работников, пока
// не будут перебраны все начальники.
// Перебирать кеши в рутинах, конкурентно. С RWMutex.
