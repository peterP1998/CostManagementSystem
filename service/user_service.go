package service
import (
	"github.com/peterP1998/CostManagementSystem/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
)
func SelectUserByName(db *sql.DB,username string)(models.User,error){
	var user models.User
	err:=db.QueryRow("SELECT * FROM User where username=?",username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err!=nil{
		return user,err
	}
	return user,nil
}
func SelectAllUsers(db *sql.DB)([]models.User,error){
	res,err:=selectAllUsersDB(db)
	if err != nil {
		return nil,err
	}
	defer res.Close()
	users := readUsersFromDB(res)
	return users,nil
}
func selectAllUsersDB(db *sql.DB)(*sql.Rows,error){
    res, err := db.Query("SELECT * FROM User")
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
func SplitUrl(r *http.Request)(int,error){
	p := strings.Split(r.URL.Path, "/user/")
	userId,err:=strconv.Atoi(p[len(p)-1])
	return userId,err
}
func SelectUserById(db *sql.DB,userId int)(models.User,error){
	var user models.User
	err :=db.QueryRow("select * from User where id=?;",userId).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	return user,err
}
func DeleteUserById(db *sql.DB,userId int,w http.ResponseWriter){
	_,err :=db.Query("delete from User where id=?;",userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func CreateUserDB(db *sql.DB,w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_,err =db.Query("insert into User(username,email,password,admin) Values(?,?,?,?);",user.Name,user.Email,user.Password,false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}