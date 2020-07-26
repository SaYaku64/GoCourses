package main

// Создать пару указателей  на существующие структуры, вызвать панику.
// Сменить в структурах вложение структур на вложение указателей.
// Написать методы для указателя на человека. (несколько)
// func (receiver *human) Happy birthday{}
// Вызвать эти методы для работников.

import (
	"fmt"
	"strings"
	"time"
)

// human - struct with fields (birthday - date in format 22.06.2008)
type human struct {
	birthday string
	name     string
	surname  string
}

// employee - struct, that are used only for keeping "human" inside :)
// experience - number of working experience in years (4.5  -  4 yrs and 6 mnths)
type employee struct {
	jobName    string
	experience float64
	info       human
}

// firstPointer - first of two methods for our pointers
// we change data in our human
func (h *human) firstPointer() {
	fmt.Println("Well... This is the first method.")
	fmt.Println("What if we change the human fields?")
	*h = human{
		birthday: "22.12.1833",
	}
	h.name = "Marko"
	h.surname = "Vovchok"
}

// happyBirthdayPointer - second method, that checks, if birthday is today
// in program calling of this function goes after calling "firstPointer", so
// you can change date in it and see, that this method works
func (h *human) happyBirthdayPointer() {
	fmt.Println("Let's check, maybe today is birthday of this human?")
	arr := strings.Split(h.birthday, ".")
	now := time.Now().Format("2006-01-02")
	timeArr := strings.Split(now, "-")

	if arr[0] == timeArr[2] && arr[1] == timeArr[1] {
		fmt.Println("Yeah! Happy birthday!")
	} else {
		fmt.Println("Oh... Em... Not today...")
	}
}

func main() {
	// creating the employee
	var emp = employee{
		jobName:    "Actor",
		experience: 7.45,
		info: human{
			birthday: "06.04.2000",
			name:     "Alex",
			surname:  "Cake",
		},
	}
	fmt.Println("Our employee:", emp)

	// creating pointer to human through employee
	var pHum *human
	pHum = &emp.info
	fmt.Println("Hmm... Let's see the adress of human's birthday info:", &pHum.birthday)

	// creating pointer to employee
	var pEmp *employee
	pEmp = &emp
	fmt.Println("And through pointer to employee let's check surname:", pEmp.info.surname)

	// calling panic? Ok...
	//panic("Oh no! It's PANIC!!!")

	// struct of new human
	newHuman := human{
		birthday: "12.12.1915",
		name:     "Frank",
		surname:  "Sinatra",
	}

	// replacing the old human structure with new by
	*pHum = newHuman

	fmt.Println("Wow! New human in emp!", emp.info)

	// rename, just for epic
	newHuman.name = "Francis Albert"
	*pHum = newHuman
	fmt.Println("But, yeah, better to change name to full name, isn't it?\n", emp.info)

	// calling methods
	fmt.Println("Alright, two methods:")
	emp.info.firstPointer()
	fmt.Println("Ok, results of changing:", emp.info)
	emp.info.happyBirthdayPointer()
}
