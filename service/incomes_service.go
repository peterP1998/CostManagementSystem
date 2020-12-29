package service
import(
	"github.com/peterP1998/CostManagementSystem/models"
)
type IncomeService struct {

}
func SelectAllIncomesForUser(id int)([]models.Income,error){
	res,err :=models.DB.Query("select * from Income where userid=?;",id)
	if err!=nil{
		return nil,err
	}
	incomes := make([]models.Income, 0)
	for res.Next() {
		var income models.Income
		res.Scan(&income.ID, &income.Description, &income.Value, &income.Userid)
		incomes = append(incomes, income)
	}
	return incomes,nil
}
func (incomeService IncomeService) CreateIncome(id int,desc string,value int)(error){
	_,err :=models.DB.Query("insert into Income(description,value,userid) Values(?,?,?);",desc,value,id)
	if err!=nil{
		return err
	}
	return nil
}