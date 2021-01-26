package service

import (
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/repository"
	"net/http"
	"strconv"
	"strings"
)

type GroupService struct {
	GroupRepositoryDB repository.GroupRepositoryInterface
}

func (groupService GroupService) SelectGroupById(groupId int) (models.Group, error) {
	group, err := groupService.GroupRepositoryDB.SelectGroupById(groupId)
	if err != nil {
		return group, err
	}
	return group, nil
}
func (groupService GroupService) SelectGroupByName(name string) (models.Group, error) {
	group, err := groupService.GroupRepositoryDB.SelectGroupByName(name)
	if err != nil {
		return group, err
	}
	return group, err
}
func (groupService GroupService) UpdateGroupMoney(id int, value int) error {
	err := groupService.GroupRepositoryDB.UpdateGroupMoney(id, value)
	if err != nil {
		return err
	}
	return nil
}
func (groupService GroupService) SelectAllGroups() ([]models.Group, error) {
	groups, err := groupService.GroupRepositoryDB.SelectAllGroups()
	if err != nil {
		return nil, err
	}
	return groups, nil
}
func (groupService GroupService) DeleteGroup(groupId int) error {
	err := groupService.GroupRepositoryDB.DeleteGroup(groupId)
	if err != nil {
		return err
	}
	return nil
}
func (groupService GroupService) CreateGroup(money string, group string) error {
	i, err := strconv.Atoi(money)
	if err != nil {
		return err
	}
	err = createGroupDB(i, group, groupService)
	if err != nil {
		return err
	}
	return nil
}
func SplitUrlGroup(r *http.Request) (int, error) {
	p := strings.Split(r.URL.Path, "/group/")
	groupId, err := strconv.Atoi(p[len(p)-1])
	return groupId, err
}

func createGroupDB(targetmoney int, groupname string, groupService GroupService) error {
	err := groupService.GroupRepositoryDB.CreateGroup(targetmoney, groupname)
	if err != nil {
		return err
	}
	return nil
}
