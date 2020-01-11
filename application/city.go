package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetCities query all weibo cities by province code
func GetCities(accessToken string, provinceCode string) ([]definitions.City, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommonGetCity, "?access_token=", accessToken, "&province=", provinceCode)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var cityArr []map[string]string
	if err := json.Unmarshal(response.OriginalBytes, &cityArr); err != nil {
		return nil, err
	}
	var cities []definitions.City
	for _, city := range cityArr {
		for key, value := range city {
			cities = append(cities, definitions.City{CityCode: key, CityName: value})
		}
	}
	return cities, nil
}
