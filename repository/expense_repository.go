package repository

import (
	"database/sql"
	"github.com/peterP1998/CostManagementSystem/models"
)

type ExpenseRepositoryInterface interface {
	CreateExpense(id int, desc string, value int, category string) error
	DeleteExpense(userId int) error
	SelectAllExpensesForUserById(id int) (*sql.Rows, error)
	GetExpensesByCategoryAndUserId(id int, category string) (*sql.Rows, error)
}
type ExpenseRepository struct {
}

func (er ExpenseRepository) SelectAllExpensesForUserById(id int) (*sql.Rows, error) {
	res, err := models.DB.Query("select * from Expense where userid=?;", id)
	if err != nil {
		return nil, err
	}
	return res, nil
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
func (er ExpenseRepository) GetExpensesByCategoryAndUserId(id int, category string) (*sql.Rows, error) {
	res, err := models.DB.Query(`select * from Expense where userid=? and category=?;`, id, category)
	if err != nil {
		return nil, err
	}
	return res, nil
}
