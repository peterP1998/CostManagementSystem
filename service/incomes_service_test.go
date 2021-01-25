package service

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IncomeRepositoryMock struct {
}

var arrIncomes []models.Income

func (er IncomeRepositoryMock) SelectAllIncomesForUserById(id int) (*sql.Rows, error) {
	if id == 2 {
		return nil, errors.New("No data for this user")
	}
	return nil, nil
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
		arrIncomes = append(arrIncomes, models.Income{2, desc, float32(value), category, id})
		return nil
	}
	return nil
}
func (er IncomeRepositoryMock) GetIncomesByCategoryAndUserId(id int, category string) (*sql.Rows, error) {
	if id == 2 {
		return nil, errors.New("No data for this user")
	}
	return nil, nil
}
func TestSelectAllIncomesForUser(t *testing.T) {
	var incomeService IncomeService = IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}
	_, err := incomeService.SelectAllIncomesForUser(2)
	assert.NotEqual(t, nil, err, "Select not working correctly")
	_, err = incomeService.SelectAllIncomesForUser(3)
	assert.Equal(t, nil, err, "Select not working correctly")
}
func TestDeleteIncome(t *testing.T) {
	var incomeService IncomeService = IncomeService{IncomeRepositoryDB: IncomeRepositoryMock{}}
	arrIncomes = append(arrIncomes, models.Income{2, "", 2.0, "category", 3})
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
