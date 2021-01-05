package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type data struct {
	Src    string `json:"sender"`
	Dest   string `json:"server"`
	HWAddr string `json:"hwaddr"`
	Msg    string `json:"msg"`
}

func makeLog(d *data, prefix string) {
	fmt.Printf("%v\nSrc: %v\nDest: %v\nHWAddr: %v\nMsg: %v\n", prefix, d.Src, d.Dest, d.HWAddr, d.Msg)
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
	makeLog(&d, "Received")

	w.WriteHeader(http.StatusOK)
}
func sendMsg() {
	flag := true
	for {
		time.Sleep(4 * time.Second)
		d := &data{
			Src:    "http://localhost:4000/",
			Dest:   "http://localhost:4002/",
			HWAddr: "00:1b:63:84:45:e6",
			Msg:    "Hello Server",
		}
		makeLog(d, "Sent")
		jsn, _ := json.Marshal(d)
		if flag {
			req, err := http.NewRequest("POST", "http://localhost:4001/", bytes.NewReader(jsn))
			if err != nil {
				panic(err)
			}
			go sendMessage(req)
		} else {
			req, err := http.NewRequest("POST", "http://localhost:4002/", bytes.NewReader(jsn))
			if err != nil {
				panic(err)
			}
			go sendMessage(req)
		}
		flag = !flag
	}
}

func main() {
	go sendMsg()
	http.HandleFunc("/", recieveMessage)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
