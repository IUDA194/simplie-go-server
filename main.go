package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
	
type Person struct { 
	Name string
	Age int
}

var people []Person

func main() {
	var port string = "8080"

	http.HandleFunc("/people", peopleHendler)

	fmt.Println("Server Listening port " + port)
	err := http.ListenAndServe("localhost:" + port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func peopleHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPeople(w,r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func postPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
	//fmt.Fprintf(w, "POST: Person added: '%v'", person)
}
