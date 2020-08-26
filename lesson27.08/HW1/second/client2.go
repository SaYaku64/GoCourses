package main

// Дз:

// Создать структуру работник, которая имеет вложенную структуру человек.
// Создать структуры пяти профессий которые будут иметь свои особые поля, и вложенную структуру работника.
// Создать теги для особых полей.
// Сменить в структурах вложение структур на вложение указателей.
// Создать Интерфейс работника умеющего называть свою работу, должность на ней и свое имя.
// Имплементировать этот интерфейс во все структуры профессий.
// Создать кеш таких работников ( из разных профессий)
// “Опросить” их всех в мапе о их работе должности и имени.
// Написать Функцию которая будет принимать кеш и возвращать типы значений каждого элемента.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var worker1 = doctor{"Pediatrician", &Employee{"Head of department", "23000", &Human{"Poshlaya Molly", 43}}}
	var worker2 = teacher{"Maths", &Employee{"Teacher", "13000", &Human{"Valentin Strikalo", 54}}}
	var worker3 = musician{"Rock", &Employee{"Singer", "40000", &Human{"Kis-Kis", 22}}}
	var worker4 = fireman{25, &Employee{"Head of squad", "28000", &Human{"Lyapis Trubetskoy", 35}}}
	var worker5 = priest{"Catholic", &Employee{"Chaplain", "100000", &Human{"Daite Tank", 67}}}

	// map with workers
	workers := make(map[int]Worker)
	{
		workers[1] = worker1
		workers[2] = worker2
		workers[3] = worker3
		workers[4] = worker4
		workers[5] = worker5
	}

	// output of name, profession and position of workers
	for _, worker := range workers {
		//fmt.Println(GetInfo(worker))
		makeRequest(GetInfo(worker))
	}

}

func makeRequest(emp string) {

	bytesRepresentation, err := json.Marshal(emp)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost/firstpost", "application/json", bytes.NewReader(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Struct has been sent.")
	}

	defer resp.Body.Close()
}
