package application

import (
	"testing"
)

func TestGetEmotions(t *testing.T) {
	result, err := GetEmotions(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}