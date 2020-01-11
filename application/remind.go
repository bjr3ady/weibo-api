package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetUnReadCount query the weibo current user's un-read remind
func GetUnReadCount(accessToken string) (*definitions.RemindResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.UnreadCount, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var remind definitions.RemindResult
	if err := json.Unmarshal(response.OriginalBytes, &remind); err != nil {
		return nil, err
	}
	return &remind, nil
}
