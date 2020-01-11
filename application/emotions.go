package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetEmotions query the weibo common emotions
func GetEmotions(accessToken string) ([]definitions.EmotionResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.Emotions, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var emotionsResult []definitions.EmotionResult
	if err := json.Unmarshal(response.OriginalBytes, &emotionsResult); err != nil {
		return nil, err
	}
	return emotionsResult, nil
}
