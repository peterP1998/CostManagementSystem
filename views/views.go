package views
import (
	"html/template"
	"net/http"
)
func CreateView(w http.ResponseWriter,pathToFile string,data interface{})(error){
	t, err := template.ParseFiles(pathToFile)
	if err!=nil{
       return err
	}
	t.Execute(w, data)
	return nil
}