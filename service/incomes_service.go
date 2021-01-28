package service

import (
	"errors"
	"github.com/peterP1998/CostManagementSystem/models"
)

type IncomeService struct {
	IncomeRepositoryDB IncomeRepositoryInterface
}
type IncomeRepositoryInterface interface {
	CreateIncome(id int, desc string, value int, category string) error
	DeleteIncome(userId int) error
	SelectAllIncomesForUserById(id int) ([]models.Income, error)
	GetIncomesByCategoryAndUserId(id int, category string) ([]models.Income, error)
}

func (incomeService IncomeService) SelectAllIncomesForUser(id int) ([]models.Income, error) {
	res, err := incomeService.IncomeRepositoryDB.SelectAllIncomesForUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (incomeService IncomeService) CreateIncome(id int, desc string, value int, category string) error {
	if category != "Salary" && category != "Gift" && category != "Found" && category != "Sell" {
		return errors.New("Wrong category")
	}
	err := incomeService.IncomeRepositoryDB.CreateIncome(id, desc, value, category)
	if err != nil {
		return err
	}
	return nil
}
func (incomeService IncomeService) DeleteIncome(userId int) error {
	err := incomeService.IncomeRepositoryDB.DeleteIncome(userId)
	if err != nil {
		return err
	}
	return nil
}
