package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetRateLimitStatus query the weibo rate limit status
func GetRateLimitStatus(accessToken string) (*definitions.RateLimitStatus, error) {
	response, err := SimpleGet(constant.BaseURI, constant.RateLimitStatus, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var rateLimitSts definitions.RateLimitStatus
	if err := json.Unmarshal(response.OriginalBytes, &rateLimitSts); err != nil {
		return nil, err
	}
	return &rateLimitSts, nil
}

//GetUID query the weibo current authorized user id
func GetUID(accessToken string) (*definitions.UIDResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.GetUID, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var uidResult definitions.UIDResult
	if err := json.Unmarshal(response.OriginalBytes, &uidResult); err != nil {
		return nil, err
	}
	return &uidResult, nil
}

//GetTokenInfo query the weibo Oauth2 token info
func GetTokenInfo(accessToken string) (*definitions.TokenInfo, error) {
	response, err := SimplePost(constant.BaseURI + constant.GetTokenInfo, map[string]string{"access_token": accessToken})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var tokenInfo definitions.TokenInfo
	if err := json.Unmarshal(response.OriginalBytes, &tokenInfo); err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}
