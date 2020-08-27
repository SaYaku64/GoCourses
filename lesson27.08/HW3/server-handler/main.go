package main

import (
	"fmt"
	"net/http"

	//jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

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

	// start page, that uses html-file (localhost/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/postlog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is my Token:", r.Header)
		//head, err := ioutil.ReadAll(r.Header.Get("Token"))
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		// resultBody := make(map[string]string)

		// // err = json.Unmarshal(head, &resultBody)
		// // if err != nil {
		// // 	fmt.Println(err.Error())
		// // }
		// fmt.Println(resultBody)
		return
	})

	http.ListenAndServe(":81", nil)
}
