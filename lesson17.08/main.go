package main

// Реализовать хендлер на роуте "/"
// с функционалом чтения(get)
// Передаем в  query параметр ID получаем структуру.
//  и создания (post):
// Передаем в теле структуру, получаем в ответ ее ID

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("[Server started]")

	var emp = make(map[int]employee)
	emp[1] = employee{"Alexander Yakushev", 20, 3}
	emp[2] = employee{"Tony Montana", 40, 15}
	emp[3] = employee{"abc", 1, 1}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Master!\n")
		id := ""
		query := r.URL.Query()
		id = query.Get("id")
		if id == "" {
			fmt.Fprintf(w, "You haven't printed the ID\n")
		} else {
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
		name := r.FormValue("username")
		age := r.FormValue("userage")
		exp := r.FormValue("userexp")

		intAge, err1 := strconv.Atoi(age)
		intExp, err2 := strconv.Atoi(exp)
		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			fmt.Fprintf(w, "Something wrong with age or exp!\n")
		}

		fmt.Fprintf(w, "Entered   ||   Full name: %s Age: %s Experience: %s\n", name, age, exp)
		for i := 0; i < len(emp); i++ {
			if emp[i].fullName == name && emp[i].age == intAge && emp[i].experience == intExp {
				fmt.Fprintf(w, "ID: %v\n", i)
			} else if i == len(emp)-1 {
				fmt.Fprintf(w, "No matches! Check mistakes!\n")
			}
		}
		fmt.Println(name, age, exp)
	})

	http.ListenAndServe(":80", nil)
}
