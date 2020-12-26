package service
import (
	"github.com/peterP1998/CostManagementSystem/models"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
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
func CreateGroupDB(w http.ResponseWriter, r *http.Request){
	var group models.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_,err =models.DB.Query("insert into Groupp(groupname,moneybynow,targetmoney) Values(?,?,?);",group.GroupName,0,group.TargetMoney)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func DeleteGroup(groupId int,w http.ResponseWriter){
    _,err :=models.DB.Query("delete from Groupp where id=?;",groupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func DonateToGroup(groupId int,userid int,w http.ResponseWriter){
	_,err :=models.DB.Query("insert into user_group(user_id,group_id) Values(?,?);",groupId,userid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}