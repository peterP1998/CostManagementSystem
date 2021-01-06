package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestParseToken(t *testing.T){
	var accountService AccountService
	token:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InBlcGkiLCJleHAiOjE2MDk5Njg1MDksImFkbWluIjp0cnVlfQ.xLcyc_3aE09756PgdK1kKatPtpj_ZtuD1g_eUV9_BLM"
	_,_,err:=accountService.ParseToken(token)
	assert.Equal(t, err, nil, "Error should be nill")
    _,_,err=accountService.ParseToken(token+"/")
	assert.NotEqual(t, err, nil, "Error should be nill")
}
