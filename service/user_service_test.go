package service

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	//"os"
	"testing"
)

type UserRepositoryMock struct {
}

var users []models.User

func (ur UserRepositoryMock) SelectAllUsers() ([]models.User, error) {
	users = append(users, models.User{ID: 1, Name: "test", Email: "petar@abv.bg", Password: "test", Admin: true})
	return users, nil
}
func (ur UserRepositoryMock) SelectUserByName(username string) (models.User, error) {
	var user models.User
	if username == "test" {
		return models.User{ID: 1, Name: "test", Email: "petar@abv.bg", Password: "test", Admin: false}, nil
	}
	return user, nil
}
func (ur UserRepositoryMock) DeleteUserById(id int) error {
	return nil
}
func (ur UserRepositoryMock) InsertUser(name string, email string, password string, admin bool) error {
	users = append(users, models.User{ID: 1, Name: name, Email: email, Password: password, Admin: admin})
	return nil
}
func TestSelectAllUsers(t *testing.T) {
	var userService UserService = UserService{ExpenseService{}, IncomeService{}, UserRepositoryMock{}}
	users, _ := userService.SelectAllUsers()
	flag := false
	for _, b := range users {
		if b.Name == "test" {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "Select not working correctly")
}
func TestSelectAllUsersWithoutAdmins(t *testing.T) {
	var userService UserService = UserService{ExpenseService{}, IncomeService{}, UserRepositoryMock{}}
	users, _ := userService.SelectAllUsersWithoutAdmins("test1234")
	flag := false
	for _, b := range users {
		if b.Name == "test" {
			flag = true
		}
	}
	assert.Equal(t, false, flag, "SelectAllUsersWithoutAdmins not working correctly")
	users, _ = userService.SelectAllUsersWithoutAdmins("test")
	flag = false
	for _, b := range users {
		if b.Name == "test" {
			flag = true
		}
	}
	assert.Equal(t, false, flag, "SelectAllUsersWithoutAdmins not working correctly")
}
func TestCreateSelectDeleteUser(t *testing.T) {
	var userService UserService = UserService{ExpenseService{ExpenseRepositoryMock{}, IncomeService{}}, IncomeService{IncomeRepositoryMock{}}, UserRepositoryMock{}}
	err := userService.CreateUser("testtest", "test@abv.bg", "test", "no")
	assert.Equal(t, err, nil, "Error should be nill")
	user, err := userService.SelectUserByName("test")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, user.Name, "test", "Select not working correctly")
	err = userService.DeleteUserById(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
}
func TestEmailValidate(t *testing.T) {
	assert.Equal(t, false, checkEmail("test not email"))
	assert.Equal(t, true, checkEmail("test@gmail.com"))
}
func TestPasswordValidate(t *testing.T) {
	assert.Equal(t, false, checkPassword("test not email"))
	assert.Equal(t, true, checkPassword("test1gmail.com"))
	assert.Equal(t, false, checkPassword("test1"))
}
