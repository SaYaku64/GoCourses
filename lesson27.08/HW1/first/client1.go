package main

// ДЗ :
// 1 повторить го тур первые 3 раздела , рассмотренные на уроке , и те что надо было знать до курса.
// 2 настройка среды ( установка GO, установка IDE , запуск hello world
// 3  создать и инициализировать значениями структуру "работник" ( тип кастомны)
// разными способами задания структур ( значений)б несколько.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type employee struct {
	ID           int    `json:"id"`
	FULLNAME     string `json:"name"`
	KILLER       bool   `json:"killer"`
	WORKINGHOURS int    `json:"hours"`
}

func main() {
	emp1 := employee{
		ID:           1,
		FULLNAME:     "Uzhasnaya Dama",
		KILLER:       true,
		WORKINGHOURS: 999,
	}
	emp2 := employee{2, "Neschastniy Valet", false, 333}
	emp3 := employee{}

	makeRequest(emp1)
	makeRequest(emp2)
	makeRequest(emp3)
}

func makeRequest(emp employee) {

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

	resp.Body.Close()
}
