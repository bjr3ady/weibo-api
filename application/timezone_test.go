package application

import (
	"testing"
)

func TestGetTimezones(t *testing.T) {
	timezones, err := GetTimezones(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(timezones)
}