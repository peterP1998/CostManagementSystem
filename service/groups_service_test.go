package service
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestCreateSelectDeleteGroup(t *testing.T){
	var groupService GroupService
	err:=groupService.CreateGroup("100","test")
	assert.Equal(t, err, nil, "Error should be nill")
	groupbyname,err:=groupService.SelectGroupByName("test")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, "test", groupbyname.GroupName, "Wrong select")
	/*flag := false
	for _, b := range incomes {
		if b.Description == "test" && b.Value==3 && b.Category=="Salary" {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "Select not working correctly")
	err=DeleteIncome(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")*/
}