package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
