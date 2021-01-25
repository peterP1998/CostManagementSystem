package repository

import (
	"github.com/peterP1998/CostManagementSystem/models"
)
type UserRepositoryInterface interface {
	SelectAllUsers()([]models.User, error)
	SelectUserByName(username string) (models.User, error)
	DeleteUserById(id int)error
	InsertUser(name string, email string, password string, admin bool)error
}
type UserRepository struct{
}
func (ur UserRepository)SelectAllUsers()([]models.User, error){
	res, err := models.DB.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for res.Next() {
		var user models.User
		res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		users = append(users, user)
	}
	return users,nil
}
func (ur UserRepository)SelectUserByName(username string) (models.User, error){
	var user models.User
	err := models.DB.QueryRow("SELECT * FROM User where username=?", username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (ur UserRepository)DeleteUserById(id int)error{
	_, err:= models.DB.Query("delete from User where id=?;", id)
	if err != nil {
		return err
	}
	return nil
}
func (ur UserRepository)InsertUser(name string, email string, password string, admin bool)error{
	_, err := models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);", name, email, password, admin)
	if err != nil {
		return err
	}
	return nil
}