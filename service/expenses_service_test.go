package service

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExpenseRepositoryMock struct {
}

var arrExpenses []models.Expense

func (er ExpenseRepositoryMock) SelectAllExpensesForUserById(id int) ([]models.Expense, error) {
	if id == 2 {
		var expenses []models.Expense
		expenses = append(expenses, models.Expense{ID: 1, Description: "test1", Value: 2.0, Category: "Other", Userid: 2})
		expenses = append(expenses, models.Expense{ID: 2, Description: "test2", Value: 3.0, Category: "Other", Userid: 2})
		return expenses, nil
	}
	return nil, errors.New("No data for this user")
}
func (er ExpenseRepositoryMock) DeleteExpense(userId int) error {
	if userId == 2 {
		arrExpenses = nil
		return nil
	}
	return nil
}
func (er ExpenseRepositoryMock) CreateExpense(id int, desc string, value int, category string) error {
	if id == 2 {
		arrExpenses = append(arrExpenses, models.Expense{ID: 2, Description: desc, Value: float32(value), Category: category, Userid: id})
		return nil
	}
	return nil
}
func (er ExpenseRepositoryMock) GetExpensesByCategoryAndUserId(id int, category string) ([]models.Expense, error) {
	if id == 2 && category == "Other" {
		var expenses []models.Expense
		expenses = append(expenses, models.Expense{ID: 1, Description: "test1", Value: 2.0, Category: "Other", Userid: 2})
		expenses = append(expenses, models.Expense{ID: 2, Description: "test2", Value: 3.0, Category: "Other", Userid: 2})
		return expenses, nil
	}
	return nil, nil
}

func TestSelectAllExpensesForUser(t *testing.T) {
	var expenseService ExpenseService = ExpenseService{ExpenseRepositoryDB: ExpenseRepositoryMock{}}
	res, err := expenseService.SelectAllExpensesForUser(2)
	assert.Equal(t, nil, err, "Select not working correctly")
	assert.Equal(t, 2, len(res), "Select not working correctly")
	_, err = expenseService.SelectAllExpensesForUser(3)
	assert.NotEqual(t, nil, err, "Select not working correctly")
}
func TestDeleteExpense(t *testing.T) {
	var expenseService ExpenseService = ExpenseService{ExpenseRepositoryDB: ExpenseRepositoryMock{}}
	arrExpenses = append(arrExpenses, models.Expense{ID: 2, Description: "test", Value: 3.0, Category: "Other", Userid: 2})
	err := expenseService.DeleteExpense(3)
	assert.Equal(t, nil, err, "Delete not working correctly")
	assert.NotEqual(t, nil, arrExpenses, "Delete not working correctly")
	err = expenseService.DeleteExpense(2)
	assert.Equal(t, nil, err, "Delete not working correctly")
	assert.Equal(t, []models.Expense([]models.Expense(nil)), arrExpenses, "Delete not working correctly")
}
func TestCreateExpense(t *testing.T) {
	var expenseService ExpenseService = ExpenseService{ExpenseRepositoryDB: ExpenseRepositoryMock{}, IncomeServiceWired: IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}}
	err := expenseService.CreateExpense(2, "", 2.0, "category")
	assert.Equal(t, "Not enough money", err.Error(), "Delete not working correctly")
	assert.NotEqual(t, nil, arrExpenses, "Create not working correctly")
}
