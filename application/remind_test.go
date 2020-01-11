package application

import (
	"testing"
)

func TestGetRemind(t *testing.T) {
	result, err := GetUnReadCount(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}