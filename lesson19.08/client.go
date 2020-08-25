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
)

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// server is used from the "lesson17.08/main.go" package/file
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func main() {
	makeRequest()
}

func makeRequest() {

	// message := map[string]string{
	// 	"username": "abc",
	// 	"userage":  "1",
	// 	"userexp":  "1",
	// }

	message := struct {
		USERNAME string `json:"username"`
		USERAGE  string `json:"userage"`
		USEREXP  string `json:"userexp"`
	}{
		USERNAME: "abc",
		USERAGE:  "1",
		USEREXP:  "1",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost/postform", "application/json", bytes.NewReader(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(string(body))
	}
}
