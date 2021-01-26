package service

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GroupRepositoryMock struct {
}

func (gr GroupRepositoryMock) SelectGroupById(id int) (models.Group, error) {
	if id == 1 {
		return models.Group{ID: 1, GroupName: "test", MoneyByNow: 0, TargetMoney: 10}, nil
	}
	return models.Group{}, nil
}
func (gr GroupRepositoryMock) SelectGroupByName(name string) (models.Group, error) {
	if name == "test" {
		return models.Group{ID: 1, GroupName: "test", MoneyByNow: 0, TargetMoney: 10}, nil
	}
	return models.Group{}, nil
}
func (gr GroupRepositoryMock) DeleteGroup(id int) error {
	if id != 1 {
		return errors.New("No group with that id")
	}
	return nil
}
func (gr GroupRepositoryMock) CreateGroup(targetmoney int, groupname string) error {
	if groupname != "test" || targetmoney != 10 {
		return errors.New("Error creating")
	}
	return nil
}
func (gr GroupRepositoryMock) UpdateGroupMoney(id int, value int) error {
	if 2 != id {
		return errors.New("Error updating")
	}
	return nil
}
func (gr GroupRepositoryMock) SelectAllGroups() ([]models.Group, error) {
	groups := make([]models.Group, 0)
	groups = append(groups, models.Group{ID: 1, GroupName: "test", MoneyByNow: 0, TargetMoney: 10})
	groups = append(groups, models.Group{ID: 2, GroupName: "test2", MoneyByNow: 0, TargetMoney: 30})
	return groups, nil
}
func TestSelectGroupById(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	group, err := groupService.SelectGroupById(2)
	assert.Equal(t, nil, err, "Select by id not working correctly")
	assert.Equal(t, models.Group{}, group, "Select by id not working correctly")
	group, err = groupService.SelectGroupById(1)
	assert.Equal(t, nil, err, "Select by id not working correctly")
	assert.Equal(t, "test", group.GroupName, "Select by id not working correctly")
}
func TestSelectByGroupName(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	group, err := groupService.SelectGroupByName("test1")
	assert.Equal(t, nil, err, "Select by name not working correctly")
	assert.Equal(t, models.Group{}, group, "Select by name not working correctly")
	group, err = groupService.SelectGroupByName("test")
	assert.Equal(t, nil, err, "Select by name not working correctly")
	assert.Equal(t, 10.0, group.TargetMoney, "Select by name not working correctly")
}
func TestUpdateGroupMoney(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	err := groupService.UpdateGroupMoney(3, 10)
	assert.Equal(t, "Error updating", err.Error(), "Update not working correctly")
	err = groupService.UpdateGroupMoney(2, 10)
	assert.Equal(t, nil, err, "Update not working correctly")
}
func TestSelectAllGroups(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	groups, _ := groupService.SelectAllGroups()
	assert.Equal(t, 2, len(groups), "Select all not working correctly")
}
func TestDeleteGroup(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	err := groupService.DeleteGroup(2)
	assert.Equal(t, "No group with that id", err.Error(), "Delete not working correctly")
	err = groupService.DeleteGroup(1)
	assert.Equal(t, nil, err, "Delete not working correctly")
}
func TestCreateGroup(t *testing.T) {
	groupService := GroupService{GroupRepositoryDB: GroupRepositoryMock{}}
	err := groupService.CreateGroup("10", "test")
	assert.Equal(t, nil, err, "Create not working correctly")
	err = groupService.CreateGroup("11", "test")
	assert.Equal(t, "Error creating", err.Error(), "Create not working correctly")
}
