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
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("/user", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user/expenses", handlers.GetExpenesesForUser).Methods("GET")
	router.HandleFunc("/user/expenses", handlers.AddExpenseForUser).Methods("PUT")
	router.HandleFunc("/user/incomes", handlers.GetIncomesForUser).Methods("GET")
	router.HandleFunc("/user/incomes", handlers.AddIncomeForUser).Methods("PUT")
	http.ListenAndServe(":8000", router)
}
