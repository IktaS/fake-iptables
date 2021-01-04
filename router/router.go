package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type data struct {
	Sender string `json:"sender"`
	Server string `json:"server"`
	Msg    string `json:"msg"`
}

func transformSender(w http.ResponseWriter, r *http.Request) {
	var d data
	err := json.NewDecoder(r.Body).Decode(&d)
	d.Sender = "192.168.2.1"
	b, _ := json.Marshal(d)
	req, _ := http.NewRequest(r.Method, d.Server, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", transformSender)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
