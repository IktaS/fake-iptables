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
	if d.Src == "http://localhost:4000/" {
		send := &data{
			Src:    "http://localhost:8081/",
			Dest:   d.Src,
			HWAddr: d.HWAddr,
			Msg:    "Not Permitted",
		}
		b, _ := json.Marshal(send)
		req, _ := http.NewRequest("POST", send.Dest, bytes.NewReader(b))
		w.WriteHeader(http.StatusOK)
		go sendMessage(req)
		return
	}
	send := &data{
		Src:    "http://localhost:8081/",
		Dest:   d.Src,
		HWAddr: d.HWAddr,
		Msg:    "Hello!",
	}
	b, _ := json.Marshal(send)
	fmt.Println(string(b))
	req, _ := http.NewRequest("POST", send.Dest, bytes.NewReader(b))
	w.WriteHeader(http.StatusOK)
	go sendMessage(req)
}

func main() {
	http.HandleFunc("/", recieveMessage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
