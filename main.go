package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Sample model
type Sample struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func getWelcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Sample{Message: "Hello,World", Success: true})
}

func handelRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/api/welcome", getWelcome).Methods("GET")
	// start server
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	fmt.Println("App started")
	handelRequest()
}
