package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// interface Person
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// interface Address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Ebiz", LastName: "Lustria", Address: &Address{City: "Sleman", State: "Yogyakarta"}})
	people = append(people, Person{ID: "2", FirstName: "Luke", LastName: "Shaw", Address: &Address{City: "Bantul", State: "Yogyakarta"}})

	// endpoints
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", GetPeopleEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4200", router))
}
