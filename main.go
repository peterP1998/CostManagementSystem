package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/controller"
	"github.com/peterP1998/CostManagementSystem/models"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/routes"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	models.DB,_=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/expenses_system")
	//router.PathPrefix("/api/login").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    http.Handle("/", router)
	router.HandleFunc("/api/login", controller.GetForm).Methods("GET")
	router.HandleFunc("/api/login", controller.Signin).Methods("POST")
	router.HandleFunc("/",controller.Welcome).Methods("GET")
	router.HandleFunc("/api/register",controller.GetRegister).Methods("GET")
	router.HandleFunc("/api/register",controller.RegisterUser).Methods("POST")
	router.HandleFunc("/api/account",controller.Account).Methods("GET")
	router.HandleFunc("/api/income",controller.IncomePage).Methods("GET")
	router.HandleFunc("/api/expense",controller.ExpensePage).Methods("GET")
	router.HandleFunc("/api/balance",controller.GetBalanceForUser).Methods("GET")
	router.HandleFunc("/api/logout", controller.Logout).Methods("GET")
	routes.UserRoutes(router)
	router.HandleFunc("/api/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
	router.HandleFunc("/api/user/expenses", controller.GetExpenesesForUser).Methods("GET")
	router.HandleFunc("/api/user/expenses", controller.AddExpenseForUser).Methods("POST")
	router.HandleFunc("/api/user/incomes", controller.GetIncomesForUser).Methods("GET")
	router.HandleFunc("/api/user/incomes", controller.AddIncomeForUser).Methods("POST")
	router.HandleFunc("/api/group/{id:[0-9]+}", controller.GetGroup).Methods("GET")
	router.HandleFunc("/api/group", controller.CreateGroup).Methods("POST")
	router.HandleFunc("/api/group/create", controller.GetCreateGroupPage).Methods("GET")
	router.HandleFunc("/api/group/{id:[0-9]+}", controller.DeleteGroup).Methods("DELETE")
	router.HandleFunc("/api/group/{id:[0-9]+}", controller.AddIncomeForUser).Methods("POST")
	http.ListenAndServe(":8090", router)
}
