package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}
type Status struct {
	Status string `json:"status"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{
		Text: "Hello, backend World",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := Status{
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/status", statusHandler)

	http.ListenAndServe(":8080", nil)
}
