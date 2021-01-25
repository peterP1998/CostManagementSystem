package service

import (
	"errors"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/repository"
)

type ExpenseService struct {
	ExpenseRepositoryDB repository.ExpenseRepositoryInterface
	IncomeServiceWired  IncomeService
}

func (expenseService ExpenseService) SelectAllExpensesForUser(id int) ([]models.Expense, error) {
	res, err := expenseService.ExpenseRepositoryDB.SelectAllExpensesForUserById(id)
	if err != nil {
		return nil, err
	}
	expenses := make([]models.Expense, 0)
	if res != nil {
		for res.Next() {
			var expense models.Expense
			res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Category, &expense.Userid)
			expenses = append(expenses, expense)
		}
	}
	return expenses, nil
}
func (expenseService ExpenseService) CreateExpense(id int, desc string, value int, category string) error {
	err := expenseService.BalanceForNewExpense(id, value)
	if err != nil {
		return err
	}
	err = expenseService.ExpenseRepositoryDB.CreateExpense(id, desc, value, category)
	if err != nil {
		return err
	}
	return nil
}
func (expenseService ExpenseService) DeleteExpense(userId int) error {
	err := expenseService.ExpenseRepositoryDB.DeleteExpense(userId)
	if err != nil {
		return err
	}
	return nil
}
func (expenseService ExpenseService) BalanceForNewExpense(id int, value int) error {
	incomes, err := expenseService.IncomeServiceWired.SelectAllIncomesForUser(id)
	if err != nil {
		return err
	}
	expenses, err := expenseService.SelectAllExpensesForUser(id)
	if err != nil {
		return err
	}
	balance := CalculateBalance(incomes, expenses)
	if int(balance)-value < 0 {
		return errors.New("Not enough money")
	}
	return nil
}
