package service

import (
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberOfExpensesAndIncomes(t *testing.T) {
	var expenseRepo ExpenseRepositoryMock
	var incomeRepo IncomeRepositoryMock
	cnt := getValueOfExpensesOfOneCategory(2, "Other", expenseRepo)
	assert.Equal(t, 0.0, cnt, "Error should be nill")
	cntIncomes := getValueOfIncomesOfOneCategory(2, "Salary", incomeRepo)
	assert.Equal(t, 0.0, cntIncomes, "Error should be nill")
}

func TestBalance(t *testing.T) {
	incomes := makeIncomesArrary()
	expenses := makeExpensesArrary()
	balance := CalculateBalance(incomes, expenses)
	assert.Equal(t, float32(1), balance, "Wrong wroking balance function")
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
