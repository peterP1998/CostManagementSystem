package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSelectDeleteGroup(t *testing.T) {
	var groupService GroupService
	testCreateGroup(t, groupService)
	groupbyname := testSelectGroupByName(t, groupService)
	testUpdateMoney(t, groupService, groupbyname.ID)
	testSelectGroupById(t, groupService, groupbyname.ID)
	testSelectAllGroups(t, groupService)
	testDeleteGroup(t, groupService, groupbyname.ID)
}

func testCreateGroup(t *testing.T, groupService GroupService) {
	err := groupService.CreateGroup("100", "tester")
	assert.Equal(t, err, nil, "Error should be nill")
}

func testSelectGroupByName(t *testing.T, groupService GroupService) models.Group {
	groupbyname, err := groupService.SelectGroupByName("tester")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, "tester", groupbyname.GroupName, "Wrong select")
	return groupbyname
}

func testUpdateMoney(t *testing.T, groupService GroupService, id int) {
	err := groupService.UpdateGroupMoney(id, 6)
	assert.Equal(t, err, nil, "Error should be nill")
}

func testSelectGroupById(t *testing.T, groupService GroupService, id int) {
	groupbyid, err := groupService.SelectGroupById(id)
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, "tester", groupbyid.GroupName, "Wrong select")
	assert.Equal(t, 6.0, groupbyid.MoneyByNow, "Wrong update money")
}

func testSelectAllGroups(t *testing.T, groupService GroupService) {
	groups, err := groupService.SelectAllGroups()
	flag := false
	for _, b := range groups {
		if b.MoneyByNow == 6 && b.GroupName == "tester" && b.TargetMoney == 100 {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "Select not working correctly")
	assert.Equal(t, err, nil, "Error should be nill")
}

func testDeleteGroup(t *testing.T, groupService GroupService, id int) {
	err := groupService.DeleteGroup(id)
	assert.Equal(t, err, nil, "Error should be nill")
}
