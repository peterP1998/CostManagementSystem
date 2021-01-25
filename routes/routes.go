package routes

import (
	"github.com/gorilla/mux"
	"github.com/peterP1998/CostManagementSystem/controller"
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/repository"
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
	userService:=service.UserService{service.ExpenseService{repository.ExpenseRepository{},service.IncomeService{}},service.IncomeService{repository.IncomeRepository{}},repository.UserRepository{}}
	route.userController=controller.UserController{service.AccountService{},userService}
	router.HandleFunc("/api/user", route.userController.GetCreateUserPage).Methods("GET")
	router.HandleFunc("/api/user", route.userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/delete", route.userController.DeleteUser).Methods("POST")
	router.HandleFunc("/api/register", route.userController.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/delete", route.userController.GetDeleteUserPage).Methods("GET")
	// router.HandleFunc("/api/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
}
func (route Route) AccountRoutes(router *mux.Router) {
	userService:=service.UserService{service.ExpenseService{},service.IncomeService{},repository.UserRepository{}}
	accountService:=service.AccountService{userService}
	route.accountController=controller.AccountController{accountService}
	router.HandleFunc("/", route.accountController.Welcome).Methods("GET")
	router.HandleFunc("/api/login", route.accountController.GetLoginForm).Methods("GET")
	router.HandleFunc("/api/login", route.accountController.Signin).Methods("POST")
	router.HandleFunc("/api/logout", route.accountController.Logout).Methods("GET")
	router.HandleFunc("/api/register", route.accountController.GetRegister).Methods("GET")
	router.HandleFunc("/api/account", route.accountController.Account).Methods("GET")
}
func (route Route) IncomeRoutes(router *mux.Router) {
	userService:=service.UserService{service.ExpenseService{},service.IncomeService{},repository.UserRepository{}}
	incomeService:=service.IncomeService{repository.IncomeRepository{}}
	route.incomeController=controller.IncomeController{service.AccountService{},incomeService,userService}
	router.HandleFunc("/api/income", route.incomeController.IncomePage).Methods("GET")
	router.HandleFunc("/api/user/incomes", route.incomeController.GetIncomesForUser).Methods("GET")
	router.HandleFunc("/api/user/incomes", route.incomeController.AddIncomeForUser).Methods("POST")

}
func (route Route) ExpenseRoutes(router *mux.Router) {
	userService:=service.UserService{service.ExpenseService{},service.IncomeService{},repository.UserRepository{}}
	expenseService:=service.ExpenseService{repository.ExpenseRepository{},service.IncomeService{IncomeRepositoryDB: repository.IncomeRepository{}}}
	route.expenseController=controller.ExpenseController{service.AccountService{},expenseService,userService}
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
	userService:=service.UserService{service.ExpenseService{},service.IncomeService{},repository.UserRepository{}}
	incomeService:=service.IncomeService{repository.IncomeRepository{}}
	expenseService:=service.ExpenseService{repository.ExpenseRepository{},service.IncomeService{IncomeRepositoryDB: repository.IncomeRepository{}}}
	route.balanceController=controller.BalanceController{service.AccountService{},service.BalanceService{},userService,expenseService,incomeService}
	router.HandleFunc("/api/balance", route.balanceController.GetBalanceForUser).Methods("GET")
}
