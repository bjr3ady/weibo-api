package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetProvinces query weibo province data by country code
func GetProvinces(accessToken string, countryCode string) ([]definitions.Province, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommonGetProvince, "?access_token=", accessToken, "&country=", countryCode)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var provinceArr []map[string]string
	if err := json.Unmarshal(response.OriginalBytes, &provinceArr); err != nil {
		return nil, err
	}
	var provinces []definitions.Province
	for _, province := range provinceArr {
		for key, value := range province {
			provinces = append(provinces, definitions.Province{ProvinceCode: key, ProvinceName: value})
		}
	}
	return provinces, nil
}
