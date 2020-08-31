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

			client := &http.Client{}
			r, err := http.NewRequest(http.MethodPost, "http://localhost:81/postlog", bytes.NewReader(bytesRepresentation)) // URL-encoded payload
			if err != nil {
				fmt.Println(err)
			} else {
				r.Header.Add("Token", string(body))
				r.Header.Add("Content-Type", "application/json")

				resp, err := client.Do(r)
				if err != nil {
					log.Println(err)
				}

				if body, err := ioutil.ReadAll(resp.Body); err != nil {
					log.Fatalln(err)
				} else {
					fmt.Println("Ok? ", string(body))
				}
			}

		} else {
			fmt.Println("Access denied!")
			fmt.Println(string(body))
		}
	}
}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	//apiUrl := "http://localhost:81/postlog"
// 	//resource := "/postlog/"

// 	message := struct {
// 		NAME     string `json:"name"`
// 		EMAIL    string `json:"email"`
// 		PASSWORD string `json:"password"`
// 	}{
// 		NAME:     "qqq",
// 		EMAIL:    "qq@qq.qq",
// 		PASSWORD: "qq",
// 	}

// 	bytesRepresentation, err := json.Marshal(message)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// data := url.Values{}
// 	// data.Set("name", "foo")
// 	// data.Set("surname", "bar")

// 	client := &http.Client{}
// 	r, _ := http.NewRequest(http.MethodPost, "http://localhost:81/postlog", bytes.NewReader(bytesRepresentation)) // URL-encoded payload
// 	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
// 	r.Header.Add("Content-Type", "application/json")
// 	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

// 	resp, err := client.Do(r)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		fmt.Println(resp.Status)
// 	}
// }
