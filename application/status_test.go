package application

import (
	"testing"
)

var TestStatusID = 4458123827521558

func TestGetHomeTimeline(t *testing.T) {
	result, err := GetHomeTimeline(TestToken, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserTimeline(t *testing.T) {
	result, err := GetUserTimeline(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetBilateralTimeline(t *testing.T) {
	result, err := GetBilateralTimeline(TestToken, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetRepostTimeline(t *testing.T) {
	result, err := GetReposts(TestToken, TestStatusID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestShowStatus(t *testing.T) {
	result, err := ShowStatus(TestToken, TestStatusID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestShowStatusCount(t *testing.T) {
	result, err := ShowStatusCommentsAndRepostsCount(TestToken, []int{TestStatusID})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestShareStatus(t *testing.T) {
	result, err := ShareStatus(TestToken, "test https://api.google.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}