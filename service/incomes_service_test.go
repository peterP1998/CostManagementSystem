package service
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSelectDeleteIncome(t *testing.T){
	var userService UserService
	user, err := userService.SelectUserByName("test")
	assert.Equal(t, err, nil, "Error should be nill")
	assert.Equal(t, user.Name, "test", "Select not working correctly")
	var incomeService IncomeService
	err=incomeService.CreateIncome(user.ID,"test",3,"Salary")
	assert.Equal(t, err, nil, "Error should be nill")
	incomes,err:=SelectAllIncomesForUser(user.ID)
	flag := false
	for _, b := range incomes {
		if b.Description == "test" && b.Value==3 && b.Category=="Salary" {
			flag = true
		}
	}
	assert.Equal(t, true, flag, "Select not working correctly")
	err=DeleteIncome(user.ID)
	assert.Equal(t, err, nil, "Error should be nill")
}