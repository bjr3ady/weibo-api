package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
	"strconv"
)

//GetUserFriends query the weibo user's friends
func GetUserFriends(accessToken string, uid, startCursorID int) (*definitions.FriendsResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.Friends, "?access_token=", accessToken, "&uid=", strconv.Itoa(uid), "&cursor=", strconv.Itoa(startCursorID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var friendsResult definitions.FriendsResult
	if err := json.Unmarshal(response.OriginalBytes, &friendsResult); err != nil {
		return nil, err
	}
	for i := range friendsResult.Users {
		user := friendsResult.Users[i]
		if err := user.Status.ShowRetweeted(); err != nil {
			return nil, err
		}
	}
	return &friendsResult, nil
}

//GetUserFollowers query the weibo user's followers
func GetUserFollowers(accessToken string, uid, startCursorID int) (*definitions.FriendsResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.Followers, "?access_token=", accessToken, "&uid=", strconv.Itoa(uid), "&cursor=", strconv.Itoa(startCursorID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var friendsResult definitions.FriendsResult
	if err := json.Unmarshal(response.OriginalBytes, &friendsResult); err != nil {
		return nil, err
	}
	for i := range friendsResult.Users {
		user := friendsResult.Users[i]
		if err := user.Status.ShowRetweeted(); err != nil {
			return nil, err
		}
	}
	return &friendsResult, nil
}

//GetFriendship query the weibo two user's friendship detail
func GetFriendship(accessToken string, sourceID, targetID int) (*definitions.FriendshipDetail, error) {
	response, err := SimpleGet(constant.BaseURI, constant.Friendship, "?access_token=", accessToken, "&source_id=", strconv.Itoa(sourceID), "&target_id=", strconv.Itoa(targetID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var friendship definitions.FriendshipDetail
	if err := json.Unmarshal(response.OriginalBytes, &friendship); err != nil {
		return nil, err
	}
	return &friendship, nil
}
