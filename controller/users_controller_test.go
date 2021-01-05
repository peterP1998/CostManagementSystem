package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestDashboard(t *testing.T) {
	mux := http.NewServeMux()
    var userController UserController
	mux.HandleFunc("/dashboard", userController.GetCreateUserPage)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/dashboard", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}