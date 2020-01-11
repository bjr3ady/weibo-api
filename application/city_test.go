package application

import (
	"testing"
)

var TestProvince = "001011"

func TestGetCity(t *testing.T) {
	result, err := GetCities(TestToken, TestProvince)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}