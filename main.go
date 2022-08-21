package main

import (
	"fmt"
	"net/http"

	"github.com/denizedizcan/REST-API-User-Management-Service/db"
	"github.com/denizedizcan/REST-API-User-Management-Service/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// start the app and handle routes
func main() {
	fmt.Println("Starting App..")
	godotenv.Load()
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
