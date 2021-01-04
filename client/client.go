package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Sender string `json:"sender"`
	Server string `json:"server"`
	Msg    string `json:"msg"`
}

func main() {
	d := &data{
		Sender: "192.168.0.0",
		Server: "http://localhost:8081/",
		Msg:    "Hello Server",
	}
	jsn, _ := json.Marshal(d)
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewReader(jsn))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
