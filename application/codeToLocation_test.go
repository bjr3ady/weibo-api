package application

import (
	"testing"
)

var TestLocationCode = "100"

func TestGetCodeToLocation(t *testing.T) {
	result, err := GetLocationByCode(TestToken, []string{TestLocationCode})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}