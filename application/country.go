package application

import (
	"errors"
	"encoding/json"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetCountries query all weibo supported countries
func GetCountries(accessToken string) ([]definitions.Country, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommonGetCountry, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var countryArr []map[string]string
	if err := json.Unmarshal(response.OriginalBytes, &countryArr); err != nil {
		return nil, err
	}
	var countries []definitions.Country
	for _, country := range countryArr {
		for key, value := range country {
			countries = append(countries, definitions.Country{CountryCode: key, CountryName: value})
		}
	}
	return countries, nil
}