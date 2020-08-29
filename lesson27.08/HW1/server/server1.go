package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var printed bool

type employee struct {
	ID       int    `json:"id"`
	FULLNAME string `json:"name"`
	KILLER   bool   `json:"killer"`
	HouRs    int    `json:"hours"`
}

func main() {
	fmt.Println("[Server-printer started]")

	http.HandleFunc("/firstpost", func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		// JSON-content-type requesting
		if contentType == "application/json" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//resultBody := make(map[string]interface{})
			fmt.Println("JSON: ", string(body))
			resultBody := employee{}
			//var resultBody interface{}

			err = json.Unmarshal(body, &resultBody)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//fmt.Println(resultBody)
			fmt.Printf("Struct: %+v\n", resultBody)
		} else {
			fmt.Fprintf(w, "\nUnknown content-type! I understand only JSON.")
		}

	})

	http.ListenAndServe(":80", nil)
}
