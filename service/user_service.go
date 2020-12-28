package service
import (
	"github.com/peterP1998/CostManagementSystem/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
	"net/http"
)
func SelectUserByName(username string)(models.User,error){
	var user models.User
	err:=models.DB.QueryRow("SELECT * FROM User where username=?",username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err!=nil{
		return user,err
	}
	return user,nil
}
func SelectAllUsers()([]models.User,error){
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
func SplitUrlUser(r *http.Request)(int,error){
	p := strings.Split(r.URL.Path, "/user/")
	userId,err:=strconv.Atoi(p[len(p)-1])
	return userId,err
}
func SelectUserById(userId int)(models.User,error){
	var user models.User
	err :=models.DB.QueryRow("select * from User where id=?;",userId).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	return user,err
}
func DeleteUserById(userId int,w http.ResponseWriter){
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
func RegisterUserDB(w http.ResponseWriter, name string,email string,password string){
	_,err :=models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);",name,email,password,false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func CreateUserDB(w http.ResponseWriter, name string,email string,password string,admin bool){
	_,err :=models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);",name,email,password,admin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}