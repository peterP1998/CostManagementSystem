package utils
import (
	"net/http"
)
func InternalServerError(err error,w http.ResponseWriter){
	if err!=nil{
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
}

func UserNotFound(err error,w http.ResponseWriter){
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}