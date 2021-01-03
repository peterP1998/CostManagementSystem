package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/models"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/routes"
)

func main() {
	fmt.Printf("Starting server at port 8090\n")
	models.DB,_=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/expenses_system")
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", router)
	var routes routes.Route
    routes.AllRoutes(router)
	http.ListenAndServe(":8090", router)
}
