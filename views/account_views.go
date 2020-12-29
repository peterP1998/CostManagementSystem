package views
import (
	"html/template"
	"net/http"
)
func CreateView(w http.ResponseWriter,pathToFile string,data interface{}){
	t, err := template.ParseFiles(pathToFile)
	if err!=nil{

	}
	t.Execute(w, data)
}