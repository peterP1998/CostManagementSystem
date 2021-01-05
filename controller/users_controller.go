package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"html/template"
	"github.com/peterP1998/CostManagementSystem/views"
)
type UserController struct {
	accountService service.AccountService
	userService service.UserService
}
func (userController UserController)GetUserAcc(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/user.html",nil)
}
func (userController UserController)GetAdminAcc(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/admin.html",nil)
}
func (userController UserController)GetCreateUserPage(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/createuser.html",nil)
}
func (userController UserController) GetDeleteUserPage(w http.ResponseWriter, r *http.Request){
	token:=userController.accountService.CheckAuthBeforeOperate(r,w)
	username,admin,err:=userController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	users,err:=userController.userService.SelectAllUsersWithoutAdmins(username)
	views.CreateView(w,"static/templates/deleteuser.html",users)
}
/*func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token:=service.CheckAuthBeforeOperate(r,w)
	tknStr := token.Value
	_,admin,err:=service.ParseToken(tknStr)
	service.CheckAdminPermission(admin,w)
	users,err:=service.SelectAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrlUser(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user,err :=service.SelectUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}*/
func (userController UserController)DeleteUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	token:=userController.accountService.CheckAuthBeforeOperate(r,w)
	_,admin,err:=userController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user,err:=userController.userService.SelectUserByName(r.FormValue("name"))
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err=userController.userService.DeleteUserById(user.ID)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}
func (userController UserController)CreateUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	t, _ := template.ParseFiles("static/templates/createuser.html")
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg":"User created succesfully"}
	err:=userController.userService.CreateUser(r.FormValue("username"),r.FormValue("email"),r.FormValue("password"),r.FormValue("admin"))
	if err != nil {
        t.Execute(w, errresp)
		return
	}
	t.Execute(w, ok)
}
func(userController UserController) RegisterUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	err:=userController.userService.RegisterUser(r.FormValue("username"),r.FormValue("email"),r.FormValue("password"))
	if err != nil {
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}
