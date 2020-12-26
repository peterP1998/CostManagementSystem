package handlers

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"html/template"
)

var jwtKey = []byte("my_secret_key")
func GetForm(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/index.html")
	t.Execute(w, nil)
}
func Welcome(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/welcome.html")
	t.Execute(w, nil)
}
func GetRegister(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/signup.html")
	t.Execute(w, nil)
}
func Signin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user,err :=service.SelectUserByName(r.FormValue("username"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if  user.Password != r.FormValue("password") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err=service.CreateAndConfigureToken(user,w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, "/api/welcome", http.StatusSeeOther)
}
func Logout(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)

	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}