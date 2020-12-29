package service
import(
	"github.com/peterP1998/CostManagementSystem/models"
)
type ExpenseService struct {

}
func SelectAllExpensesForUser(id int)([]models.Expense,error){
	res,err :=models.DB.Query("select * from Expense where userid=?;",id)
	if err!=nil{
		return nil,err
	}
	expenses := make([]models.Expense, 0)
	for res.Next() {
		var expense models.Expense
		res.Scan(&expense.ID, &expense.Description, &expense.Value,&expense.Category, &expense.Userid)
		expenses = append(expenses, expense)
	}
	return expenses,nil
}
func (expenseService ExpenseService) CreateExpense(id int,desc string,value int,category string)(error){
	_,err :=models.DB.Query("insert into Expense(description,value,category,userid) Values(?,?,?,?);",desc,value,category,id)
	if err!=nil{
		return err
	}
	return nil
}

