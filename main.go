package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/handlers"
	"github.com/peterP1998/CostManagementSystem/models"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	models.DB,_=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/expenses_system")
	//router.PathPrefix("/api/login").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    http.Handle("/", router)
	router.HandleFunc("/api/login", handlers.GetForm).Methods("GET")
	router.HandleFunc("/api/login", handlers.Signin).Methods("POST")
	router.HandleFunc("/api/welcome",handlers.Welcome).Methods("GET")
	router.HandleFunc("/api/register",handlers.GetRegister).Methods("GET")
	router.HandleFunc("/api/register",handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user", handlers.CreateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/expenses", handlers.GetExpenesesForUser).Methods("GET")
	router.HandleFunc("/api/user/expenses", handlers.AddExpenseForUser).Methods("PUT")
	router.HandleFunc("/api/user/incomes", handlers.GetIncomesForUser).Methods("GET")
	router.HandleFunc("/api/user/incomes", handlers.AddIncomeForUser).Methods("PUT")
	router.HandleFunc("/api/group/{id:[0-9]+}", handlers.GetGroup).Methods("GET")
	router.HandleFunc("/api/group", handlers.CreateGroup).Methods("PUT")
	router.HandleFunc("/api/group/{id:[0-9]+}", handlers.DeleteGroup).Methods("DELETE")
	router.HandleFunc("/api/group/{id:[0-9]+}", handlers.AddIncomeForUser).Methods("POST")
	http.ListenAndServe(":8090", router)
}
