package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
)

type UserController struct {
	accountService service.AccountService
	userService    service.UserService
}

func (userController UserController) GetUserAcc(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/accounts/user.html", nil)
	utils.InternalServerError(err, w)
}
func (userController UserController) GetAdminAcc(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/accounts/admin.html", nil)
	utils.InternalServerError(err, w)
}
func (userController UserController) GetCreateUserPage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/user/createuser.html", nil)
	utils.InternalServerError(err, w)
}
func (userController UserController) GetDeleteUserPage(w http.ResponseWriter, r *http.Request) {
	token := userController.accountService.CheckAuthBeforeOperate(r, w)
	username, admin, err := userController.accountService.ParseToken(token.Value)
	if admin == false || err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	users, err := userController.userService.SelectAllUsersWithoutAdmins(username)
	utils.InternalServerError(err, w)
	usersmap := map[string]interface{}{
		"messg": "",
		"users": users,
	}
	err = views.CreateView(w, "static/templates/user/deleteuser.html", usersmap)
	utils.InternalServerError(err, w)
}

func (userController UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	users, err := userController.userService.SelectAllUsersWithoutAdmins(r.FormValue("name"))
	errresp := map[string]interface{}{"users": users, "messg": "Something went wrong!Try again!"}
	okresp := map[string]interface{}{"users": users, "messg": "User deleted succesfully!"}
	if err != nil {
		views.CreateView(w, "static/templates/user/deleteuser.html", errresp)
	}
	token := userController.accountService.CheckAuthBeforeOperate(r, w)
	_, admin, err := userController.accountService.ParseToken(token.Value)
	if admin == false || err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := userController.userService.SelectUserByName(r.FormValue("name"))
	if err != nil {
		err = views.CreateView(w, "static/templates/user/deleteuser.html", errresp)
		utils.InternalServerError(err, w)
		return
	}
	err = userController.userService.DeleteUserById(user.ID)
	if err != nil {
		err = views.CreateView(w, "static/templates/user/deleteuser.html", errresp)
		utils.InternalServerError(err, w)
		return
	}
	views.CreateView(w, "static/templates/user/deleteuser.html", okresp)
}
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg": "User created succesfully"}
	valid, resp := userController.userService.CheckInputsBeforeCreating(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	if !valid {
		err := views.CreateView(w, "static/templates/user/createuser.html", map[string]interface{}{"messg": resp})
		utils.InternalServerError(err, w)
		return
	}
	err := userController.userService.CreateUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("admin"))
	if err != nil {
		err = views.CreateView(w, "static/templates/user/createuser.html", errresp)
		utils.InternalServerError(err, w)
		return
	}
	err = views.CreateView(w, "static/templates/user/createuser.html", ok)
	utils.InternalServerError(err, w)
}
func (userController UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	valid, resp := userController.userService.CheckInputsBeforeCreating(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	if !valid {
		err := views.CreateView(w, "static/templates/accounts/signup.html", map[string]interface{}{"messg": resp})
		utils.InternalServerError(err, w)
		return
	}
	err := userController.userService.RegisterUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	utils.InternalServerError(err, w)
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}
