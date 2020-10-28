package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//declaring test global variable
var hello = "Hello World!"

func main() {

	// get ENV variable value for os
	var listenPort string
	listenPort = os.Getenv("PORT")

	// create mux router
	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld).Methods("GET") //use mux method for request
	http.ListenAndServe(":"+listenPort, r)       //create webserver

}

// create function for JSON format

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // setup JSON header
	json.NewEncoder(w).Encode(hello)                   // encoding hello variable to JSON

}
