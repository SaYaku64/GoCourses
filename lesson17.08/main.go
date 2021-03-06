package main

// Реализовать хендлер на роуте "/"
// с функционалом чтения(get)
// Передаем в  query параметр ID получаем структуру.
//  и создания (post):
// Передаем в теле структуру, получаем в ответ ее ID

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		contentType := r.Header.Get("Content-Type")
		fmt.Println(contentType)

		// JSON-content-type requesting
		if contentType == "application/json" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			resultBody := make(map[string]string)

			err = json.Unmarshal(body, &resultBody)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(resultBody)

			name := resultBody["username"]
			age := resultBody["userage"]
			exp := resultBody["userexp"]

			intAge, err1 := strconv.Atoi(age)
			intExp, err2 := strconv.Atoi(exp)

			if err1 != nil || err2 != nil {
				fmt.Println(err1)
				fmt.Println(err2)
				fmt.Fprintf(w, "Something wrong with age or exp!\n")
			} else {
				//fmt.Printf("Entered   ||   Full name: %v Age: %v Experience: %v\n", name, age, exp)
				fmt.Fprintf(w, "Entered   ||   Full name: %v Age: %v Experience: %v\n", name, age, exp)
				for i := 1; i < len(emp)+1; i++ {
					if emp[i].fullName == name && emp[i].age == intAge && emp[i].experience == intExp {
						//fmt.Printf("ID: %v\n", i)
						fmt.Fprintf(w, "ID: %v\n", i)
						printed = true
					} else if i == len(emp) && printed == false {
						fmt.Fprintf(w, "No matches! Check mistakes!\n")
						//fmt.Printf("No matches! Check mistakes!\n")
					}
				}
			}
			return

		} else if contentType == "application/x-www-form-urlencoded" {
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
			return
		}

	})

	http.ListenAndServe(":80", nil)
}
