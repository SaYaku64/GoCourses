package main

// Создать клиента, который будет отправлять запросы на ваш сервер, на эндпоинты:
// Получение, создание, вашего воркера , с помощью функций пакета "net/http" :
// func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
// func PostForm(url string, data url.Values) (resp *Response, err error)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	message := map[string]interface{}{
		"username": "abc",
		"userage":  1,
		"userexp":  1,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost/postform", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	//log.Println(result["data"])
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	data := []byte(`{"ping"}`)
// 	r := bytes.NewReader(data)
// 	resp, err := http.Post("http://localhost/pingpost", "application/json", r)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	for true {

// 		bs := make([]byte, 1014)
// 		n, err := resp.Body.Read(bs)
// 		fmt.Println(string(bs[:n]))

// 		if n == 0 || err != nil {
// 			break
// 		}
// 	}
// }
