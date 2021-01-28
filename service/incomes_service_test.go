package service

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IncomeRepositoryMock struct {
}

var arrIncomes []models.Income

func (er IncomeRepositoryMock) SelectAllIncomesForUserById(id int) ([]models.Income, error) {
	if id == 2 {
		var incomes []models.Income
		incomes = append(arrIncomes, models.Income{ID: 1, Description: "test1", Value: 3.0, Category: "Other", Userid: 2})
		incomes = append(arrIncomes, models.Income{ID: 2, Description: "test2", Value: 5.0, Category: "Other", Userid: 2})
		return incomes, nil
	}
	return nil, errors.New("No data for this user")
}
func (er IncomeRepositoryMock) DeleteIncome(userId int) error {
	if userId == 2 {
		arrIncomes = nil
		return nil
	}
	return nil
}
func (er IncomeRepositoryMock) CreateIncome(id int, desc string, value int, category string) error {
	if id == 2 {
		arrIncomes = append(arrIncomes, models.Income{ID: 2, Description: desc, Value: float32(value), Category: category, Userid: id})
		return nil
	}
	return nil
}
func (er IncomeRepositoryMock) GetIncomesByCategoryAndUserId(id int, category string) ([]models.Income, error) {
	if id == 2 {
		var incomes []models.Income
		incomes = append(arrIncomes, models.Income{ID: 1, Description: "test1", Value: 3.0, Category: "Other", Userid: 2})
		incomes = append(arrIncomes, models.Income{ID: 2, Description: "test2", Value: 5.0, Category: "Other", Userid: 2})
		return incomes, nil
	}
	return nil, errors.New("No data for this user")
}
func TestSelectAllIncomesForUser(t *testing.T) {
	var incomeService IncomeService = IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}
	res, err := incomeService.SelectAllIncomesForUser(2)
	assert.Equal(t, 1, len(res), "Select not working correctly")
	assert.Equal(t, nil, err, "Select not working correctly")
	_, err = incomeService.SelectAllIncomesForUser(3)
	assert.NotEqual(t, nil, err, "Select not working correctly")
}
func TestDeleteIncome(t *testing.T) {
	var incomeService IncomeService = IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}
	arrIncomes = append(arrIncomes, models.Income{ID: 2, Description: "", Value: 2.0, Category: "category", Userid: 3})
	err := incomeService.DeleteIncome(3)
	assert.Equal(t, nil, err, "Delete not working correctly")
	assert.NotEqual(t, nil, arrIncomes, "Delete not working correctly")
	err = incomeService.DeleteIncome(2)
	assert.Equal(t, nil, err, "Delete not working correctly")
	assert.Equal(t, []models.Income([]models.Income(nil)), arrIncomes, "Delete not working correctly")
}
func TestCreateIncome(t *testing.T) {
	var incomeService IncomeService = IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}
	err := incomeService.CreateIncome(2, "", 2.0, "category")
	assert.Equal(t, "Wrong category", err.Error(), "Create not working correctly")
	assert.NotEqual(t, nil, arrIncomes, "Create not working correctly")
	err = incomeService.CreateIncome(2, "", 2.0, "Salary")
	assert.Equal(t, nil, err, "Create not working correctly")
	assert.NotEqual(t, nil, arrIncomes, "Create not working correctly")
}
