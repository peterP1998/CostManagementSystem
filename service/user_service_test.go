package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	models.DB, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/expenses_system")
	models.DB.Query("insert into User(username,email,password,admin) Values(?,?,?,?);", "test", "test", "test", false)
	exitVal := m.Run()
	models.DB.Query("delete from User where username=?", "test")
	os.Exit(exitVal)
}

func TestSelectAllUsers(t *testing.T) {
	var userService UserService
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
	var userService UserService
	users, _ := userService.SelectAllUsersWithoutAdmins("test1234")
	flag := false
	for _, b := range users {
		if b.Name == "test" {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "SelectAllUsersWithoutAdmins not working correctly")
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
	var userService UserService
	err := userService.CreateUser("testtest", "test@abv.bg", "test", "no")
	assert.Equal(t, err, nil, "Error should be nill")
	user, err := userService.SelectUserByName("testtest")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, user.Name, "testtest", "Select not working correctly")
	err = userService.DeleteUserById(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
}
