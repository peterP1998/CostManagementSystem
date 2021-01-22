package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSelectDeleteExpense(t *testing.T) {
	var userService UserService
	user, err := userService.SelectUserByName("test")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, user.Name, "test", "Select not working correctly")
	var expenseService ExpenseService
	err = expenseService.CreateExpense(user.ID, "test", 0, "Food")
	assert.Equal(t, err, nil, "Error should be nill")
	expenses, err := SelectAllExpensesForUser(user.ID)
	flag := false
	for _, b := range expenses {
		if b.Description == "test" && b.Value == 0 && b.Category == "Food" {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "Select not working correctly")
	err = DeleteExpense(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
}
