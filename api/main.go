package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/denizedizcan/REST-API-User-Management-Service/api/db"
	"github.com/denizedizcan/REST-API-User-Management-Service/api/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// start the app and handle routes
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("Starting App..")
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.SetMiddlewareJSON(h.CreateUser)).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.SetMiddlewareJSON(h.UpdateUser)).Methods("PATCH")
	router.HandleFunc("/users/{id}", handlers.SetMiddlewareJSON(h.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.SetMiddlewareJSON(h.ShowUser)).Methods("GET")
	router.HandleFunc("/users", handlers.SetMiddlewareJSON(h.ShowAllUsers)).Methods("GET")

	http.ListenAndServe(":8080", router)
}