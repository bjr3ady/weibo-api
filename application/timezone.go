package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
	"strconv"
)

//GetTimezones query all weibo supported timezones
func GetTimezones(accessToken string) ([]definitions.Timezone, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommonGetTimezone, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	timezoneMap := map[string]string{}
	err = json.Unmarshal(response.OriginalBytes, &timezoneMap)
	if err != nil {
		return nil, err
	}
	var timezones []definitions.Timezone
	for key, zone := range timezoneMap {
		intKey, _ := strconv.Atoi(key)
		timezones = append(timezones, definitions.Timezone{ZoneCode: intKey, ZoneValue: zone})
	}
	return timezones, nil
}
