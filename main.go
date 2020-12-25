package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/handlers"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	router := mux.NewRouter()
	router.HandleFunc("/login", handlers.Signin).Methods("GET")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	http.ListenAndServe(":8000", router)
}
