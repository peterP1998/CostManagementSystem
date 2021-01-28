package repository

import (
	"github.com/peterP1998/CostManagementSystem/models"
)

type IncomeRepository struct {
}

func (ir IncomeRepository) SelectAllIncomesForUserById(id int) ([]models.Income, error) {
	res, err := models.DB.Query("select * from Income where userid=?;", id)
	if err != nil {
		return nil, err
	}
	var incomes []models.Income
	if res != nil {
		for res.Next() {
			var income models.Income
			res.Scan(&income.ID, &income.Description, &income.Value, &income.Category, &income.Userid)
			incomes = append(incomes, income)
		}
	}
	return incomes, nil
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
func (ir IncomeRepository) GetIncomesByCategoryAndUserId(id int, category string) ([]models.Income, error) {
	res, err := models.DB.Query(`select * from Income where userid=? and category=?;`, id, category)
	if err != nil {
		return nil, err
	}
	var incomes []models.Income
	if res != nil {
		for res.Next() {
			var income models.Income
			res.Scan(&income.ID, &income.Description, &income.Value, &income.Category, &income.Userid)
			incomes = append(incomes, income)
		}
	}
	return incomes, nil
}
