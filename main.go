package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

var contacts []Contact

// Get all contacts
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Get a contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for _, item := range contacts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Add new contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, c := range contacts {
		if c.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = params["id"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

// Delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data
	contacts = append(contacts, Contact{ID: "1", Name: "Alonso Calle", Address: "Av. Bolognesi", Facebook: "AlonsoCN", Twitter: "@AlonsoCN"})
	contacts = append(contacts, Contact{ID: "2", Name: "Diego Abanto", Address: "Av. Alfonso Ugarte", Facebook: "DiegoAbanto", Twitter: "@DiegoAbanto"})

	// Route handles & endpoints
	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", deleteContact).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
