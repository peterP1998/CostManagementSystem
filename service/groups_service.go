package service

import (
	"github.com/peterP1998/CostManagementSystem/models"
	"net/http"
	"strconv"
	"strings"
)

type GroupService struct {
}

func (groupService GroupService) SelectGroupById(groupId int) (models.Group, error) {
	var group models.Group
	err := models.DB.QueryRow("select * from Groupp where id=?;", groupId).Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
	return group, err
}
func (groupService GroupService) SelectGroupByName(name string) (models.Group, error) {
	var group models.Group
	err := models.DB.QueryRow("select * from Groupp where groupname=?;", name).Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
	return group, err
}
func (groupService GroupService) UpdateGroupMoney(id int, value int) error {
	_, err := models.DB.Query("Update Groupp set moneybynow=? where id=?;", value, id)
	return err
}
func (groupService GroupService) SelectAllGroups() ([]models.Group, error) {
	res, err := models.DB.Query("select * from Groupp;")
	groups := make([]models.Group, 0)
	for res.Next() {
		var group models.Group
		res.Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
		groups = append(groups, group)
	}
	defer res.Close()
	return groups, err
}
func (groupService GroupService) DeleteGroup(groupId int) error{
	_, err := models.DB.Query("delete from Groupp where id=?;", groupId)
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
	err = createGroupDB(i, group)
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

func createGroupDB(targetmoney int, groupname string) error {
	_, err := models.DB.Query("insert into Groupp(groupname,moneybynow,targetmoney) Values(?,?,?);", groupname, 0, targetmoney)
	if err != nil {
		return err
	}
	return nil
}

func DonateToGroup(groupId int, money int, w http.ResponseWriter) {
	_, err := models.DB.Query("UPDATE Groupp SET moneybynow='?' WHERE id=?", money, groupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
