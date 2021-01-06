package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/peterP1998/CostManagementSystem/utils"
)
type UserController struct {
	accountService service.AccountService
	userService service.UserService
}
func (userController UserController)GetUserAcc(w http.ResponseWriter, r *http.Request){
	err:=views.CreateView(w,"static/templates/user.html",nil)
	utils.InternalServerError(err,w)
}
func (userController UserController)GetAdminAcc(w http.ResponseWriter, r *http.Request){
	err:=views.CreateView(w,"static/templates/admin.html",nil)
	utils.InternalServerError(err,w)
}
func (userController UserController)GetCreateUserPage(w http.ResponseWriter, r *http.Request){
	err:=views.CreateView(w,"static/templates/createuser.html",nil)
    utils.InternalServerError(err,w)
}
func (userController UserController) GetDeleteUserPage(w http.ResponseWriter, r *http.Request){
	token:=userController.accountService.CheckAuthBeforeOperate(r,w)
	username,admin,err:=userController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	users,err:=userController.userService.SelectAllUsersWithoutAdmins(username)
	utils.InternalServerError(err,w)
	err=views.CreateView(w,"static/templates/deleteuser.html",users)
	utils.InternalServerError(err,w)
}

func (userController UserController)DeleteUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	token:=userController.accountService.CheckAuthBeforeOperate(r,w)
	_,admin,err:=userController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user,err:=userController.userService.SelectUserByName(r.FormValue("name"))
	utils.InternalServerError(err,w)
	err=userController.userService.DeleteUserById(user.ID)
	utils.InternalServerError(err,w)
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}
func (userController UserController)CreateUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg":"User created succesfully"}
	err:=userController.userService.CreateUser(r.FormValue("username"),r.FormValue("email"),r.FormValue("password"),r.FormValue("admin"))
	if err != nil {
		err=views.CreateView(w,"static/templates/createuser.html",errresp)
		utils.InternalServerError(err,w)
		return
	}
	err=views.CreateView(w,"static/templates/createuser.html",ok)
	utils.InternalServerError(err,w)
}
func(userController UserController) RegisterUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	err:=userController.userService.RegisterUser(r.FormValue("username"),r.FormValue("email"),r.FormValue("password"))
	utils.InternalServerError(err,w)
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}
