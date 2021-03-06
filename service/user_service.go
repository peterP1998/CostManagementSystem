package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type UserService struct {
	ExpenseS       ExpenseService
	IncomeS        IncomeService
	UserRepository UserRepositoryInterface
}
type UserRepositoryInterface interface {
	SelectAllUsers() ([]models.User, error)
	SelectUserByName(username string) (models.User, error)
	DeleteUserById(id int) error
	InsertUser(name string, email string, password string, admin bool) error
}

func (userService UserService) SelectUserByName(username string) (models.User, error) {
	user, err := userService.UserRepository.SelectUserByName(username)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (userService UserService) SelectAllUsers() ([]models.User, error) {
	res, err := userService.UserRepository.SelectAllUsers()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func readUsersFromDB(res *sql.Rows) []models.User {
	users := make([]models.User, 0)
	for res.Next() {
		var user models.User
		res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		users = append(users, user)
	}
	return users
}

func (userService UserService) SelectAllUsersWithoutAdmins(username string) ([]models.User, error) {
	users, err := userService.SelectAllUsers()
	if err != nil {
		return nil, err
	}
	users_without_admins := make([]models.User, 0)
	for _, v := range users {
		if v.Name != username && v.Admin == false {
			users_without_admins = append(users_without_admins, v)
		}
	}
	return users_without_admins, nil
}
func (userService UserService) DeleteUserById(userId int) error {
	err := userService.ExpenseS.DeleteExpense(userId)
	if err != nil {
		return err
	}
	err = userService.IncomeS.DeleteIncome(userId)
	if err != nil {
		return err
	}
	err = userService.UserRepository.DeleteUserById(userId)
	if err != nil {
		return err
	}
	return nil
}
func (userService UserService) RegisterUser(name string, email string, password string) error {
	err := createUser(name, email, password, false, userService.UserRepository)
	if err != nil {
		return err
	}
	return nil
}
func createUserDB(name string, email string, password string, admin bool, userRepository UserRepositoryInterface) error {
	err := userRepository.InsertUser(name, email, password, admin)
	if err != nil {
		return err
	}
	return nil
}
func createUser(name string, email string, password string, admin bool, userRepository UserRepositoryInterface) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = createUserDB(name, email, string(hashedPassword), admin, userRepository)
	if err != nil {
		return err
	}
	return nil
}
func (userService UserService) CreateUser(name string, email string, password string, admin string) error {
	adminval := false
	if admin == "yes" {
		adminval = true
	}
	err := createUser(name, email, password, adminval, userService.UserRepository)
	if err != nil {
		return err
	}
	return nil
}
func (userService UserService) CheckInputsBeforeCreating(username string, email string, password string) (bool, string) {
	if checkDoesUserWithThatNameExists(username, userService) {
		return false, "User with that name already exists"
	}
	if !checkEmail(email) {
		return false, "Invalid email adress"
	}
	if !checkPassword(password) {
		return false, "Invalid password"
	}
	return true, ""
}
func checkDoesUserWithThatNameExists(name string, userService UserService) bool {
	users, _ := userService.SelectAllUsers()
	for _, v := range users {
		if v.Name == name {
			return true
		}
	}
	return false
}
func checkEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}
func checkPassword(pass string) bool {
	var passRegex = regexp.MustCompile("([A-Za-z]+[0-9]|[0-9]+[A-Za-z])[A-Za-z0-9]*")
	if len(pass) < 6 {
		return false
	}
	return passRegex.MatchString(pass)
}
