package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var tables = make(map[string]string)

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

func masqueradeRequest(w http.ResponseWriter, r *http.Request) {
	var d data
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		panic(err)
	}
	if d.Dest == "http://localhost:8080" {
		val := tables[d.HWAddr]
		d.Dest = val
	} else {
		tables[d.HWAddr] = d.Src
		d.Src = "http://localhost:8080"
	}
	b, _ := json.Marshal(d)
	req, _ := http.NewRequest(r.Method, d.Dest, bytes.NewReader(b))
	w.WriteHeader(http.StatusOK)
	go sendMessage(req)
}

func main() {
	http.HandleFunc("/", masqueradeRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
