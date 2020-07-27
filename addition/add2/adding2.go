// Создать Интерфейс работника умеющего называть свою работу, должность на ней и свое имя.
// Имплементировать этот интерфейс во все структуры профессий.
// Создать кеш таких работников ( из разных профессий)
// “Опросить” их всех в мапе о их работе должности и имени.
// Написать Функцию которая будет принимать кеш и возвращать типы значений каждого элемента.

package main

import (
	"fmt"
)

// methods for interface for every professions
func (d doctor) GetNamePosition() string {
	return d.person.fullName + " : Doctor : " + d.position
}
func (t teacher) GetNamePosition() string {
	return t.person.fullName + " : Teacher : " + t.position
}
func (m musician) GetNamePosition() string {
	return m.person.fullName + " : Musician : " + m.position
}
func (fm fireman) GetNamePosition() string {
	return fm.person.fullName + " : Fireman : " + fm.position
}
func (pr priest) GetNamePosition() string {
	return pr.person.fullName + " : Priest : " + pr.position
}

type person struct {
	fullName string
	age      int
}

type worker struct {
	position string
	salary   string
	person   person
}

type doctor struct {
	speciality string
	worker
}

type teacher struct {
	mainSubject string
	worker
}

type musician struct {
	playingStyle string
	worker
}
type fireman struct {
	departuresOnCalls int
	worker
}
type priest struct {
	faith string
	worker
}

// Worker - interface, that is used to show name, profession and position
type Worker interface {
	GetNamePosition() string
}

// GetInfo function used to get info with interface
func GetInfo(w Worker) (info string) {
	info = w.GetNamePosition()
	return
}

func main() {

	// people with professions
	var worker1 = doctor{
		speciality: "Pediatrician",
		worker: worker{
			position: "Head of department",
			salary:   "23000",
			person: person{
				fullName: "Albert Lorghful",
				age:      43,
			},
		},
	}

	var worker2 = teacher{
		mainSubject: "Maths",
		worker: worker{
			position: "Teacher",
			salary:   "13000",
			person: person{
				fullName: "Lucie Brook",
				age:      54,
			},
		},
	}

	var worker3 = musician{
		playingStyle: "Rock",
		worker: worker{
			position: "Singer",
			salary:   "40000",
			person: person{
				fullName: "Birdy Angers",
				age:      22,
			},
		},
	}

	var worker4 = fireman{
		departuresOnCalls: 25,
		worker: worker{
			position: "Head of squad",
			salary:   "28000",
			person: person{
				fullName: "William Kornelf",
				age:      35,
			},
		},
	}

	var worker5 = priest{
		faith: "Catholic",
		worker: worker{
			position: "Chaplain",
			salary:   "100000",
			person: person{
				fullName: "Udendorfius Ptelofanirgen",
				age:      67,
			},
		},
	}

	// map with workers
	workers := make(map[int]Worker)
	{
		workers[1] = worker1
		workers[2] = worker2
		workers[3] = worker3
		workers[4] = worker4
		workers[5] = worker5
	}
	// fmt.Println(workers)

	// output of name, profession and position of workers
	for _, worker := range workers {
		fmt.Println(GetInfo(worker))
	}
}
