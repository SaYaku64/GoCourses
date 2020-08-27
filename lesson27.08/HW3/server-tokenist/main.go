package main

// Дз3: Добавить систему авторизации сервер сервер.
// Логика такова: вначале идет запрос авторизации от клиента ( по логину и паролю )
// В ответ приходит токен.
// (Токен существует 5 минут, по прошествии этого времени он становится не действительным со стороны сервера)

// Все следующие запросы должны содержать токен в хедере, иначе возврат ошибки, отсутствует авторизация.

// В случае этой ошибки, перезапросить токен.
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	name     string
	email    string
	password string
}

// type Claims struct {
// 	jwt.StandardClaims
// 	Username string `json:"name"`
// }

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
	myHash2, _ := HashPassword("qq")
	myHash3, _ := HashPassword("qwerty")
	usr[1] = user{"SaYaku", "example@gmail.com", myHash}
	usr[2] = user{"qqq", "qq@qq.qq", myHash2}
	usr[3] = user{"qwerty", "qwerty@qwer.ty", myHash3}

	// start page, that uses html-file (localhost/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/postlog", func(w http.ResponseWriter, r *http.Request) {
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

		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.Claims{
		// 	StandardClaims: jwt.StandardClaims{},
		// 	Username:       "ussseerr",
		// })
		// fmt.Println(token)

		name := resultBody["name"]
		//email := resultBody["email"]
		pass := resultBody["password"]

		//checking name and password
		for i := 1; i < len(usr)+1; i++ {
			if usr[i].name == name && CheckPasswordHash(pass, usr[i].password) {
				//place for creating TOKEN and sending it...
				//ok, I don't understand, jwt-tokens, so let's do it together
				//this "token" is silly try to make something that will be sent to client
				hash, err := HashPassword("AcCeSsAcCePtEd")
				if err != nil {
					fmt.Printf("Trouble in hashing!")
					fmt.Fprintf(w, "Trouble in hashing!")
				} else {
					fmt.Fprintf(w, hash)
				}
				printed = true
			} else if i == len(usr) && printed == false {
				fmt.Fprintf(w, "Wrong name or password!")
			}
		}

		printed = false
		return
	})

	http.ListenAndServe(":80", nil)
}
