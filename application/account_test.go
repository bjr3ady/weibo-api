package application

import (
	"testing"
)

func TestGetRateLimitStatus(t *testing.T) {
	result, err := GetRateLimitStatus(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUID(t *testing.T) {
	result, err := GetUID(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetTokenInfo(t *testing.T) {
	result, err := GetTokenInfo(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}