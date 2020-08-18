package main

// Реализовать хендлер на роуте "/"
// с функционалом чтения(get)
// Передаем в  query параметр ID получаем структуру.
//  и создания (post):
// Передаем в теле структуру, получаем в ответ ее ID

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var printed bool

func main() {
	fmt.Println("[Server started]")

	// map of employees
	var emp = make(map[int]employee)
	emp[1] = employee{"Alexander Yakushev", 20, 3}
	emp[2] = employee{"Tony Montana", 40, 15}
	emp[3] = employee{"abc", 1, 1}

	// start page, that uses html-file (localhost/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})

	// handle for get query
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello Master!\n")

		id := ""
		query := r.URL.Query()
		id = query.Get("id")

		if id == "" {
			fmt.Fprintf(w, "You haven't printed the ID\n")
		} else {
			// checking if ID is a number
			intID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Check your entered ID!\n")
			} else {
				fmt.Fprintf(w, "Entered ID: %v\n", id)
				fmt.Fprintf(w, "Structure with this ID: %v\n", emp[intID])
			}
		}
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		// getting values from post query
		var result http.Header
		result = r.Header
		// contentTypes := map[int][]string{
		// 	1: []string{"application/json"},
		// 	2: []string{"application/x-www-form-urlencoded"},
		// }
		//temp := []string{}
		// json.NewDecoder(r.Body).Decode(&result)

		// log.Println(result)

		// когда делаем пост-запрос через браузер, то не нужно декодировать с байтов в строку
		// вероятно я не всё до конца понимаю, хз, буду разбираться как посплю
		contentType := result["Content-Type"]
		fmt.Println(contentType)

		if contentType[0] == "application/json" {
			var resultB map[string]interface{}
			json.NewDecoder(r.Body).Decode(&resultB)

			name := (resultB["username"]).(string)
			age := (resultB["userage"]).(float64)
			exp := (resultB["userexp"]).(float64)

			var intAge int = int(age)
			var intExp int = int(exp)

			fmt.Printf("Entered   ||   Full name: %v Age: %v Experience: %v\n", name, age, exp)

			for i := 1; i < len(emp)+1; i++ {
				if emp[i].fullName == name && emp[i].age == intAge && emp[i].experience == intExp {
					fmt.Printf("ID: %v\n", i)
					printed = true
				} else if i == len(emp) && printed == false {
					fmt.Printf("No matches! Check mistakes!\n")
				}
			}

		} else if contentType[0] == "application/x-www-form-urlencoded" {
			name := r.FormValue("username")
			age := r.FormValue("userage")
			exp := r.FormValue("userexp")

			fmt.Fprintf(w, "Entered   ||   Full name: %s Age: %s Experience: %s\n", name, age, exp)
			//fmt.Printf("Entered   ||   Full name: %s Age: %s Experience: %s\n", name, age, exp)
			// making int versions of age and exp for manipulations
			intAge, err1 := strconv.Atoi(age)
			intExp, err2 := strconv.Atoi(exp)
			if err1 != nil || err2 != nil {
				fmt.Println(err1)
				fmt.Println(err2)
				fmt.Fprintf(w, "Something wrong with age or exp!\n")
			} else {
				for i := 1; i < len(emp)+1; i++ {
					if emp[i].fullName == name && emp[i].age == intAge && emp[i].experience == intExp {
						fmt.Fprintf(w, "ID: %v\n", i)
						printed = true
					} else if i == len(emp) && printed == false {
						fmt.Fprintf(w, "No matches! Check mistakes!\n")
					}
				}
			}
		}

	})

	http.ListenAndServe(":80", nil)
}
