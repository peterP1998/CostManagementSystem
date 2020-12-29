package service
import (
	"github.com/peterP1998/CostManagementSystem/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)
type UserService struct {

}
func (userService UserService)SelectUserByName(username string)(models.User,error){
	var user models.User
	err:=models.DB.QueryRow("SELECT * FROM User where username=?",username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err!=nil{
		return user,err
	}
	return user,nil
}
func (userService UserService)SelectAllUsers()([]models.User,error){
	res,err:=selectAllUsersDB()
	if err != nil {
		return nil,err
	}
	defer res.Close()
	users := readUsersFromDB(res)
	return users,nil
}
func selectAllUsersDB()(*sql.Rows,error){
    res, err := models.DB.Query("SELECT * FROM User")
	if err != nil {
		return nil,err
	}
	return res,nil
}
func readUsersFromDB(res *sql.Rows)([]models.User){
	users := make([]models.User, 0)
    for res.Next() {
		var user models.User
		res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		users = append(users, user)
	}
	return users
}
func (userService UserService)SplitUrlUser(r *http.Request)(int,error){
	p := strings.Split(r.URL.Path, "/user/")
	userId,err:=strconv.Atoi(p[len(p)-1])
	return userId,err
}
func(userService UserService) SelectUserById(userId int)(models.User,error){
	var user models.User
	err :=models.DB.QueryRow("select * from User where id=?;",userId).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	return user,err
}
func (userService UserService) SelectAllUsersWithoutAdmins(username string)([]models.User,error){
	users,err:=userService.SelectAllUsers()
	if err != nil {
		return nil,err
	}
	users_without_admins:=make([]models.User, 0)
	for _, v := range users {
        if v.Name!=username&&v.Admin==false{
			users_without_admins = append(users_without_admins, v)
		}
	}
	return users_without_admins,nil
}
func(userService UserService) DeleteUserById(userId int,w http.ResponseWriter){
	_,err :=models.DB.Query("delete from Expense where userid=?;",userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_,err =models.DB.Query("delete from Income where userid=?;",userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_,err =models.DB.Query("delete from User where id=?;",userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (userService UserService)RegisterUser(w http.ResponseWriter, name string,email string,password string)(error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
	}
	err=registerUserDB(w,name,email,string(hashedPassword))
	if err != nil {
        return err
	}
	return nil
}
func registerUserDB(w http.ResponseWriter, name string,email string,password string)(error){
	_,err :=models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);",name,email,password,false)
	if err != nil {
		return err
	}
	return nil
}
func createUserDB(w http.ResponseWriter, name string,email string,password string,admin bool)(error){
	_,err :=models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);",name,email,password,admin)
	if err != nil {
		return err
	}
	return nil
}
func (userService UserService)CreateUser(w http.ResponseWriter, name string,email string,password string,admin string)(error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
       return err
	}
	adminval:=false
	if admin=="yes"{
		adminval=true
	}
	err=createUserDB(w,name,email,string(hashedPassword),adminval)
	if err != nil {
		return err
	}
    return nil
}
