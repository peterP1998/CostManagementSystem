package routes

import (
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/controller"
)

type Route struct {
	userController    controller.UserController
	accountController controller.AccountController
	balanceController controller.BalanceController
	expenseController controller.ExpenseController
	incomeController  controller.IncomeController
	groupController   controller.GroupController
}

func (route Route) AllRoutes(router *mux.Router) {
	route.UserRoutes(router)
	route.AccountRoutes(router)
	route.IncomeRoutes(router)
	route.ExpenseRoutes(router)
	route.GroupRoutes(router)
	route.BalanceRoutes(router)
}
func (route Route) UserRoutes(router *mux.Router) {
	router.HandleFunc("/api/user", route.userController.GetCreateUserPage).Methods("GET")
	router.HandleFunc("/api/user", route.userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/delete", route.userController.DeleteUser).Methods("POST")
	router.HandleFunc("/api/register", route.userController.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/delete", route.userController.GetDeleteUserPage).Methods("GET")
	// router.HandleFunc("/api/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
}
func (route Route) AccountRoutes(router *mux.Router) {
	router.HandleFunc("/", route.accountController.Welcome).Methods("GET")
	router.HandleFunc("/api/login", route.accountController.GetLoginForm).Methods("GET")
	router.HandleFunc("/api/login", route.accountController.Signin).Methods("POST")
	router.HandleFunc("/api/logout", route.accountController.Logout).Methods("GET")
	router.HandleFunc("/api/register", route.accountController.GetRegister).Methods("GET")
	router.HandleFunc("/api/account", route.accountController.Account).Methods("GET")
}
func (route Route) IncomeRoutes(router *mux.Router) {
	router.HandleFunc("/api/income", route.incomeController.IncomePage).Methods("GET")
	router.HandleFunc("/api/user/incomes", route.incomeController.GetIncomesForUser).Methods("GET")
	router.HandleFunc("/api/user/incomes", route.incomeController.AddIncomeForUser).Methods("POST")

}
func (route Route) ExpenseRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/expenses", route.expenseController.GetExpenesesForUser).Methods("GET")
	router.HandleFunc("/api/user/expenses", route.expenseController.AddExpenseForUser).Methods("POST")
	router.HandleFunc("/api/expense", route.expenseController.ExpensePage).Methods("GET")
}
func (route Route) GroupRoutes(router *mux.Router) {
	router.HandleFunc("/api/group", route.groupController.CreateGroup).Methods("POST")
	router.HandleFunc("/api/group/create", route.groupController.GetCreateGroupPage).Methods("GET")
	router.HandleFunc("/api/group/donate", route.groupController.GetDonateGroupPage).Methods("GET")
	router.HandleFunc("/api/group/donate", route.groupController.DonateMoney).Methods("POST")

}
func (route Route) BalanceRoutes(router *mux.Router) {
	router.HandleFunc("/api/balance", route.balanceController.GetBalanceForUser).Methods("GET")
}
