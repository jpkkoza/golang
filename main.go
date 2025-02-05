package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	task = req.Task

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if task == "" {
		fmt.Fprintf(w, "Hello, task not set")
	} else {
		fmt.Fprintf(w, "Hello, %s", task)
	}
}

func main() {
	task = "test task"
	router := mux.NewRouter()
	router.HandleFunc("/api/task", SetTaskHandler).Methods("POST")
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
