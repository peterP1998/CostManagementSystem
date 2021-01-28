package repository

import (
	"github.com/peterP1998/CostManagementSystem/models"
)

type GroupRepository struct {
}

func (ur GroupRepository) SelectGroupById(id int) (models.Group, error) {
	var group models.Group
	err := models.DB.QueryRow("select * from Groupp where id=?;", id).Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
	return group, err
}
func (ur GroupRepository) SelectGroupByName(name string) (models.Group, error) {
	var group models.Group
	err := models.DB.QueryRow("select * from Groupp where groupname=?;", name).Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
	return group, err
}
func (ur GroupRepository) DeleteGroup(id int) error {
	_, err := models.DB.Query("delete from Groupp where id=?;", id)
	if err != nil {
		return err
	}
	return nil
}
func (ur GroupRepository) CreateGroup(targetmoney int, groupname string) error {
	_, err := models.DB.Query("insert into Groupp(groupname,moneybynow,targetmoney) Values(?,?,?);", groupname, 0, targetmoney)
	if err != nil {
		return err
	}
	return nil
}
func (ur GroupRepository) UpdateGroupMoney(id int, value int) error {
	_, err := models.DB.Query("Update Groupp set moneybynow=? where id=?;", value, id)
	return err
}
func (ur GroupRepository) SelectAllGroups() ([]models.Group, error) {
	res, err := models.DB.Query("select * from Groupp;")
	if err != nil {
		return nil, err
	}
	groups := make([]models.Group, 0)
	for res.Next() {
		var group models.Group
		res.Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
		groups = append(groups, group)
	}
	defer res.Close()
	return groups, err
}
