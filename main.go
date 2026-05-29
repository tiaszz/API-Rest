package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

var clients []Client

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateClient).Methods("POST")
	router.HandleFunc("/client/{id}", GetClient).Methods("GET")
	router.HandleFunc("/client/update/{id}", GetClient).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	log.Fatal(srv.ListenAndServe())
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var newClient Client

	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients = append(clients, newClient)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClient)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	clientID := vars["id"]

	id, err := strconv.Atoi(clientID)
	if err != nil {
		http.Error(w, "Error converting string to int", http.StatusInternalServerError)
		return
	}

	client := clients[id-1]
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	clientID := vars["id"]

	id, err := strconv.Atoi(clientID)
	if err != nil {
		http.Error(w, "Error converting string to int", http.StatusInternalServerError)
		return
	}

	var updatedClient Client

	err = json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients[id-1] = updatedClient

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClient)
}
