package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/peterP1998/CostManagementSystem/models"
	"testing"
)

func TestTokenFunctions(t *testing.T){
	claims:=testConfigureToken(t)
	token:=testCreateToken(t,claims)
	testParseToken(t,token)
}

func testParseToken(t *testing.T,token string){
	var accountService AccountService
	_,_,err:=accountService.ParseToken(token)
	assert.Equal(t, err, nil, "Error should be nill")
    _,_,err=accountService.ParseToken(token+"/")
	assert.NotEqual(t, err, nil, "Error should be nill")
}

func testConfigureToken(t *testing.T)(*models.Claims){
	var user models.User
	user.Name="test"
	user.Admin=false
	claims,_:=configureToken(user)
	assert.Equal(t, "test", claims.Username, "ConfigureToken not working")
	assert.Equal(t, false, claims.Admin, "ConfigureToken not working")
    return claims
}

func testCreateToken(t *testing.T,claims *models.Claims)(string){
	token,err:=createToken(claims)
	assert.Equal(t, err, nil, "Error should be nill")
	return token
}