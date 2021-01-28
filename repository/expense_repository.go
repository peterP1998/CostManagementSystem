package repository

import (
	"github.com/peterP1998/CostManagementSystem/models"
)

type ExpenseRepository struct {
}

func (er ExpenseRepository) SelectAllExpensesForUserById(id int) ([]models.Expense, error) {
	res, err := models.DB.Query("select * from Expense where userid=?;", id)
	if err != nil {
		return nil, err
	}
	var expenses []models.Expense
	if res != nil {
		for res.Next() {
			var expense models.Expense
			res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Category, &expense.Userid)
			expenses = append(expenses, expense)
		}
	}
	return expenses, nil
}
func (er ExpenseRepository) DeleteExpense(userId int) error {
	_, err := models.DB.Query("delete from Expense where userid=?;", userId)
	if err != nil {
		return err
	}
	return nil
}
func (er ExpenseRepository) CreateExpense(id int, desc string, value int, category string) error {
	_, err := models.DB.Query("insert into Expense(description,value,category,userid) Values(?,?,?,?);", desc, value, category, id)
	if err != nil {
		return err
	}
	return nil
}
func (er ExpenseRepository) GetExpensesByCategoryAndUserId(id int, category string) ([]models.Expense, error) {
	res, err := models.DB.Query(`select * from Expense where userid=? and category=?;`, id, category)
	if err != nil {
		return nil, err
	}
	var expenses []models.Expense
	if res != nil {
		for res.Next() {
			var expense models.Expense
			res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Category, &expense.Userid)
			expenses = append(expenses, expense)
		}
	}
	return expenses, nil
}
