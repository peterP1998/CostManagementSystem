package service

import (
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberOfExpensesAndIncomes(t *testing.T) {
	var userService UserService
	user, err := userService.SelectUserByName("test")
	var incomeService IncomeService
	err = incomeService.CreateIncome(user.ID, "test", 3, "Salary")
	var expenseService ExpenseService
	err = expenseService.CreateExpense(user.ID, "test", 3, "Other")
	cnt := getValueOfExpensesOfOneCategory(user.ID, "Other")
	assert.Equal(t, 3.0, cnt, "Error should be nill")
	cntIncomes := getValueOfIncomesOfOneCategory(user.ID, "Salary")
	assert.Equal(t, 3.0, cntIncomes, "Error should be nill")
	err = DeleteIncome(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
	err = DeleteExpense(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
}

func TestBalance(t *testing.T) {
	incomes := makeIncomesArrary()
	expenses := makeExpensesArrary()
	balance := CalculateBalance(incomes, expenses)
	assert.Equal(t, float32(1), balance, "Wrong wroking balanc function")
}

func makeExpensesArrary() []models.Expense {
	expenses := make([]models.Expense, 0)
	var expense models.Expense
	expense.Value = 3
	expenses = append(expenses, expense)
	return expenses
}
func makeIncomesArrary() []models.Income {
	incomes := make([]models.Income, 0)
	var income models.Income
	income.Value = 4
	incomes = append(incomes, income)
	return incomes
}
