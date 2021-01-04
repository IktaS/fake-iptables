package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type data struct {
	Sender string `json:"sender"`
	Server string `json:"server"`
	Msg    string `json:"msg"`
}

func recieveMessage(w http.ResponseWriter, r *http.Request) {
	var d data
	json.NewDecoder(r.Body).Decode(&d)
	fmt.Println(d)
	if d.Sender == "192.168.0.0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/", recieveMessage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
