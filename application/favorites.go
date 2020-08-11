package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
)

//GetFavorites query weibo user's favorites
func GetFavorites(accessToken string, page, count int) (*definitions.Favorites, error) {
	response, err := SimpleGet(constant.BaseURI, constant.Favorites, "?access_token=", accessToken, "&count=", strconv.Itoa(count+2), "&page=", strconv.Itoa((page)))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var favorites definitions.Favorites
	if err := json.Unmarshal(response.OriginalBytes, &favorites); err != nil {
		return nil, err
	}
	return &favorites, nil
}

//GetFavoriteIDs query weibo user's favorites' ids
func GetFavoriteIDs(accessToken string, page, count int) (*definitions.FavoriteIDs, error) {
	response, err := SimpleGet(constant.BaseURI, constant.FavoriteIDs, "?access_token=", accessToken, "&count=", strconv.Itoa(count), "&page=", strconv.Itoa((page)))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var favoriteIDs definitions.FavoriteIDs
	if err := json.Unmarshal(response.OriginalBytes, &favoriteIDs); err != nil {
		return nil, err
	}
	return &favoriteIDs, nil
}

//RemoveFavorite remove status from favorite
func RemoveFavorite(accessToken, stsID string) (*definitions.Status, error) {
	response, err := SimplePost(fmt.Sprintf("%s%s", constant.BaseURI, constant.RemoveFavorite), map[string]string{
		"access_token": accessToken,
		"id":           stsID,
	})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var status definitions.Status
	if err := json.Unmarshal(response.OriginalBytes, &status); err != nil {
		return nil, err
	}
	return &status, nil
}
