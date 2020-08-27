package main

// Создать клиента, который будет отправлять запросы на ваш сервер, на эндпоинты:
// Получение, создание, вашего воркера , с помощью функций пакета "net/http" :
// func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
// func PostForm(url string, data url.Values) (resp *Response, err error)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// server is used from the "lesson17.08/main.go" package/file
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func main() {
	makeRequest()
}

// CheckPasswordHash - checking if encoded password == password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func makeRequest() {

	message := struct {
		NAME     string `json:"name"`
		EMAIL    string `json:"email"`
		PASSWORD string `json:"password"`
	}{
		NAME:     "qqq",
		EMAIL:    "qq@qq.qq",
		PASSWORD: "qq",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:80/postlog", "application/json", bytes.NewReader(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	} else {
		if CheckPasswordHash("AcCeSsAcCePtEd", string(body)) {
			fmt.Println("Yep, Access accepted!")
			req, err := http.NewRequest("POST", "http://localhost:81/postlog", bytes.NewReader(bytesRepresentation))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Token: ", string(body))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Token", string(body))

				defer req.Body.Close()

				if body, err := ioutil.ReadAll(req.Body); err != nil {
					log.Fatalln(err)
				} else {
					fmt.Println(string(body), "dddddd")
				}
			}

		} else {
			fmt.Println("Access denied!")
			fmt.Println(string(body))
		}
	}
}
