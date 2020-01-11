package application

import (
	"testing"
)

func TestGetCountries(t *testing.T) {
	result, err := GetCountries(TestToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
