package service

import (
	"errors"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/repository"
)

type IncomeService struct {
	IncomeRepositoryDB repository.IncomeRepositoryInterface
}

func (incomeService IncomeService) SelectAllIncomesForUser(id int) ([]models.Income, error) {
	res, err := incomeService.IncomeRepositoryDB.SelectAllIncomesForUserById(id)
	if err != nil {
		return nil, err
	}
	incomes := make([]models.Income, 0)
	if res != nil {
		for res.Next() {
			var income models.Income
			res.Scan(&income.ID, &income.Description, &income.Value, &income.Category, &income.Userid)
			incomes = append(incomes, income)
		}
	}
	return incomes, nil
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
