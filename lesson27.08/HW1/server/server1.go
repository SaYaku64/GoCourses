package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var printed bool

func main() {
	fmt.Println("[Server-printer started]")

	http.HandleFunc("/firstpost", func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		// JSON-content-type requesting
		if contentType == "application/json" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			// resultBody := make(map[string]interface{})
			var resultBody interface{}

			err = json.Unmarshal(body, &resultBody)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(resultBody)
		} else {
			fmt.Fprintf(w, "\nUnknown content-type! I understand only JSON.")
		}

	})

	http.ListenAndServe(":80", nil)
}
