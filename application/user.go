package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
	"strconv"
)

//GetUserByID query weibo user's information by id
func GetUserByID(accessToken string, uid int) (*definitions.User, error) {
	response, err := SimpleGet(constant.BaseURI, constant.UserShow, "?access_token=", accessToken, "&uid="+strconv.Itoa(uid))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var user definitions.User
	if err := json.Unmarshal(response.OriginalBytes, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUserByDomain query weibo user's information by domain
func GetUserByDomain(accessToken string, domain string) (*definitions.User, error) {
	response, err := SimpleGet(constant.BaseURI, constant.DomainShow, "?access_token=", accessToken, "&domain=", domain)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var user definitions.User
	if err := json.Unmarshal(response.OriginalBytes, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
