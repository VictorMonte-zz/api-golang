package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"firstname,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {

}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people,
		Person{
			ID:        "1",
			Firstname: "Gandalf",
			Lastname:  "The Gray",
			Address:   &Address{City: "Middle", State: "Earth"}})

	people = append(people,
		Person{
			ID:        "2",
			Firstname: "Bruce",
			Lastname:  "Wayne",
			Address:   &Address{City: "Gothan", State: "NYC"}})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
