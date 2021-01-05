package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type data struct {
	Src    string `json:"sender"`
	Dest   string `json:"server"`
	HWAddr string `json:"hwaddr"`
	Msg    string `json:"msg"`
}

func sendMessage(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func recieveMessage(w http.ResponseWriter, r *http.Request) {
	var d data
	json.NewDecoder(r.Body).Decode(&d)
	fmt.Println(d)
	w.WriteHeader(http.StatusOK)
}

func main() {
	d := &data{
		Src:    "http://localhost:4000/",
		Dest:   "http://localhost:8081/",
		HWAddr: "00:1b:63:84:45:e6",
		Msg:    "Hello Server",
	}
	jsn, _ := json.Marshal(d)
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewReader(jsn))
	if err != nil {
		panic(err)
	}
	go sendMessage(req)

	http.HandleFunc("/", recieveMessage)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
