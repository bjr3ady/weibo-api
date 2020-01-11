package application

import (
	"testing"
)

var TestCountryCode = "001"

func TestGetProvinces(t *testing.T) {
	result, err := GetProvinces(TestToken, TestCountryCode)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}