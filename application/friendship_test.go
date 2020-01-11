package application

import (
	"testing"
)

var TargetUserID = 2658081315

func TestGetFollowing(t *testing.T) {
	result, err := GetUserFriends(TestToken, TestUserID, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetFollowers(t *testing.T) {
	result, err := GetUserFollowers(TestToken, TestUserID, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetFriendship(t *testing.T) {
	result, err := GetFriendship(TestToken, TestUserID, TargetUserID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}