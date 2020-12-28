package service
import (
	"github.com/peterP1998/CostManagementSystem/models"
	"strings"
	"strconv"
	"net/http"
)
func SelectGroupById(groupId int)(models.Group,error){
	var group models.Group
	err :=models.DB.QueryRow("select * from Groupp where id=?;",groupId).Scan(&group.ID, &group.GroupName, &group.MoneyByNow, &group.TargetMoney)
	return group,err
}
func SplitUrlGroup(r *http.Request)(int,error){
	p := strings.Split(r.URL.Path, "/group/")
	groupId,err:=strconv.Atoi(p[len(p)-1])
	return groupId,err
}
func CreateGroupDB(targetmoney int, groupname string)(error){
	_,err :=models.DB.Query("insert into Groupp(groupname,moneybynow,targetmoney) Values(?,?,?);",groupname,0,targetmoney)
	if err != nil {
		return err
	}
	return nil
}
func DeleteGroup(groupId int,w http.ResponseWriter){
    _,err :=models.DB.Query("delete from Groupp where id=?;",groupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func DonateToGroup(groupId int,money int,w http.ResponseWriter){
	_,err :=models.DB.Query("UPDATE Groupp SET moneybynow='?' WHERE id=?",money,groupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}