package service
import(
	"github.com/peterP1998/CostManagementSystem/models"
)
func SelectAllExpensesForUser(id int)([]models.Expense,error){
	res,err :=models.DB.Query("select * from Expense where userid=?;",id)
	if err!=nil{
		return nil,err
	}
	expenses := make([]models.Expense, 0)
	for res.Next() {
		var expense models.Expense
		res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Userid)
		expenses = append(expenses, expense)
	}
	return expenses,nil
}
func CreateExpense(id int,desc string,value int,category string)(error){
	_,err :=models.DB.Query("insert into Expense(description,value,category,userid) Values(?,?,?,?);",desc,value,category,id)
	if err!=nil{
		return err
	}
	return nil
}
func GetNumberOfExpensesOfOneCategory(id int,category string)(float64){
	var cnt float64
    _= models.DB.QueryRow(`select count(*) from Expense where userid=? and category=?`,id,category).Scan(&cnt)
    return cnt
}