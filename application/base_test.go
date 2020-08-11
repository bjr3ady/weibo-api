package application

import (
	"encoding/json"
	constant "github.com/bjr3ady/weibo-api/constant"
	"testing"
)

var TestToken = ""

func TestSimpleGetCountry(t *testing.T) {
	result, err := SimpleGet(constant.BaseURI + constant.CommonGetTimezone + "?access_token=" + TestToken + "&language=zh-cn")
	if err != nil || result == nil {
		t.Error(err)
		return
	}
	t.Log(result)
	countries := map[string]string{}
	err = json.Unmarshal(result.OriginalBytes, &countries)
	if err != nil {

	}
	t.Log(countries)
}

func TestShowUserInfo(t *testing.T) {
	result, err := SimpleGet(constant.BaseURI + constant.UserShow + "?access_token=" + TestToken + "&uid=1760323357")
	if err != nil || result == nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
