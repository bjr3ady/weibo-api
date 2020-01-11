package application

import (
	"testing"
)

var TestUserID = 2658081315
var TestUserDomain = ""

func TestGetUserByID(t *testing.T) {
	result, err := GetUserByID(TestToken, TestUserID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserByDomain(t *testing.T) {
	result, err := GetUserByDomain(TestToken, TestUserDomain)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}