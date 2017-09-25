package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person : struct that represents the users
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address : struct that represents some location
type Address struct {
	City  string `json:"firstname,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// GetPeopleEndpoint : Get all users in local list
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint : Get some user by id in local list
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePersonEndpoint : Insert user in local list
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

// DeletePersonEndpoint : Remove some user in local list
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people,
		Person{
			ID:        "1",
			Firstname: "Gandalf",
			Lastname:  "The Gray"})

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
