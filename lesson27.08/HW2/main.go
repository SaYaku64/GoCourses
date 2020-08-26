package main

// Дз 2 Создать элементарные формы регистрации и авторизации на html.
// Создать эндпоинты и обработчики что бы:
// Регистрироваться( пароли хранить в хешированом виде)
// Авторизоваться.
// Отправлять запросы из этих форм.

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	name     string
	email    string
	password string
}

var printed bool
var exist bool

// HashPassword - encoding string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - checking if encoded password == password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	fmt.Println("[Server started]")
	var usr = make(map[int]user)

	myHash, _ := HashPassword("123")
	usr[1] = user{"SaYaku", "example@gmail.com", myHash}

	// start page, that uses html-file (localhost/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})

	http.HandleFunc("/postreg", func(w http.ResponseWriter, r *http.Request) {
		newName := r.FormValue("newname")
		email := r.FormValue("email")
		newPass := r.FormValue("newpass")

		// checking empty boxes
		if newName != "" && email != "" && newPass != "" {

			// checking existing account
			for i := 1; i < len(usr)+1; i++ {
				if usr[i].name == newName || usr[i].email == email {
					fmt.Fprintf(w, "Oops! This name or email is already taken!\n")
					exist = true
				}
			}

			//register new user
			if exist == false {
				hash, err := HashPassword(newPass)
				if err != nil {
					fmt.Fprintf(w, "Error in encoding password!\n%v\n", err)
				} else {
					usr[len(usr)+1] = user{newName, email, hash}
					fmt.Fprintf(w, "Successfully registered!\n")
				}
			}
		} else {
			fmt.Fprintf(w, "You haven't filled all the boxes!\n")
		}

		exist = false
		return
	})

	http.HandleFunc("/postlog", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		pass := r.FormValue("pass")

		if name != "" && pass != "" {
			//checking name and password
			for i := 1; i < len(usr)+1; i++ {
				if usr[i].name == name && CheckPasswordHash(pass, usr[i].password) {
					fmt.Fprintf(w, "Welcome back, %s\n", name)
					printed = true
				} else if i == len(usr) && printed == false {
					fmt.Fprintf(w, "Wrong name or password!\n")
				}
			}
		} else {
			fmt.Fprintf(w, "You haven't filled all the boxes!\n")
		}

		printed = false
		return
	})

	http.ListenAndServe(":80", nil)
}
