package routes
import ("github.com/gorilla/mux"
"github.com/peterP1998/CostManagementSystem/controller")
func ExpensesRoutes(router *mux.Router) {
	
	router.HandleFunc("/api/user", controller.GetCreateUserPage).Methods("GET")
	router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/delete", controller.DeleteUser).Methods("POST")
	router.HandleFunc("/api/user/delete", controller.GetDeleteUserPage).Methods("GET")
}