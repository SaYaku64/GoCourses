package main

// Создать функцию которая приводит работников к человеку(значение типов)
// Создать кеш начальников и работников не приступать опросу ( перебору) работников пока не будут перебраны все начальники.
// Перебирать кеши в рутинах, конкурентно. С RWMutex.

import (
	"fmt"
)

func empToHum(e employee) human {
	return e.info
}

func readHead(m map[int]head) {
	fmt.Println("Checking heads:")
	for i := range m {
		fmt.Println(m[i])
	}
}

func readEmp(m map[int]employee) {
	fmt.Println("Checking employees:")
	for i := range m {
		fmt.Println(m[i])
	}
}

// func (m map[int]head) reading() string {
// 	for i := range m {
// 		fmt.Println(m[i])
// 	}
// 	return ""
// }

func main() {
	var head1 = head{
		departmentName: "Human Resources",
		subordinates:   26,
		info: human{
			birthday: "26.11.1982",
			name:     "Michael",
			surname:  "McCalisto",
		},
	}

	var head2 = head{
		departmentName: "Connection",
		subordinates:   8,
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

	fmt.Println("Human-employee-1:", empToHum(emp1))

	heads := make(map[int]head)
	{
		heads[1] = head1
		heads[2] = head2
	}

	emps := make(map[int]employee)
	{
		emps[1] = emp1
		emps[2] = emp2
	}

	mux := myMutex{}

	mux.mutex.Lock()
	go readHead(heads)
	mux.mutex.Unlock()
	go readEmp(emps)

	var input string
	fmt.Scanln(&input)
}

// Создать функцию которая приводит работников к человеку(значение типов)
// Создать кеш начальников и работников не приступать опросу ( перебору) работников, пока
// не будут перебраны все начальники.
// Перебирать кеши в рутинах, конкурентно. С RWMutex.
