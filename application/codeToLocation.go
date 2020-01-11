package application

import (
	"encoding/json"
	"strings"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetLocationByCode query the weibo location by code
func GetLocationByCode(accessToken string, codes []string) ([]definitions.Location, error) {
	codeStr := strings.Join(codes, ",")
	response, err := SimpleGet(constant.BaseURI, constant.CommonCodeToLocation, "?access_token=", accessToken, "&codes=", codeStr)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var locationArr []map[string]string
	if err := json.Unmarshal(response.OriginalBytes, &locationArr); err != nil {
		return nil, err
	}
	var locations []definitions.Location
	for _, location := range locationArr {
		for key, value := range location {
			locations = append(locations, definitions.Location{LocationCode: key, LocationName: value})
		}
	}
	return locations, nil
}