package main

import (
	"net/http"

	"fmt"

	handler "github.com/MalikSaddique/go_learning/jwt-auth-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	router.HandleFunc("/protected", handler.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
}
