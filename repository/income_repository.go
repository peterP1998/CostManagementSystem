package repository

import (
	"database/sql"
	"github.com/peterP1998/CostManagementSystem/models"
)

type IncomeRepositoryInterface interface {
	CreateIncome(id int, desc string, value int, category string) error
	DeleteIncome(userId int) error
	SelectAllIncomesForUserById(id int) (*sql.Rows, error)
	GetIncomesByCategoryAndUserId(id int, category string) (*sql.Rows, error)
}
type IncomeRepository struct {
}

func (ir IncomeRepository) SelectAllIncomesForUserById(id int) (*sql.Rows, error) {
	res, err := models.DB.Query("select * from Income where userid=?;", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (ir IncomeRepository) DeleteIncome(userId int) error {
	_, err := models.DB.Query("delete from Income where userid=?;", userId)
	if err != nil {
		return err
	}
	return nil
}
func (ir IncomeRepository) CreateIncome(id int, desc string, value int, category string) error {
	_, err := models.DB.Query("insert into Income(description,value,category,userid) Values(?,?,?,?);", desc, value, category, id)
	if err != nil {
		return err
	}
	return nil
}
func (ir IncomeRepository) GetIncomesByCategoryAndUserId(id int, category string) (*sql.Rows, error) {
	res, err := models.DB.Query(`select * from Income where userid=? and category=?;`, id, category)
	if err != nil {
		return nil, err
	}
	return res, nil
}
