package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := Message{Message: "Hello, world. This is the Greeter service!", Version: "v3"}
	data, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
