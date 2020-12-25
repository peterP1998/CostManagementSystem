package service
import(
	"github.com/peterP1998/CostManagementSystem/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"io"
)
func SelectAllExpensesForUser(db *sql.DB,id int)([]models.Expense,error){
	res,err :=db.Query("select * from Expense where userid=?;",id)
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
func CreateExpense(db *sql.DB,id int,body io.Reader)(error){
    var expense models.Expense
	err := json.NewDecoder(body).Decode(&expense)
	if err!=nil{
		return err
	}
	_,err =db.Query("insert into Expense(description,value,userid) Values(?,?,?);",expense.Description,expense.Value,id)
	if err!=nil{
		return err
	}
	return nil
}