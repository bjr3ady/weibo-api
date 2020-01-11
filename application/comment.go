package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
	"strconv"
)

func handleComment(comment *definitions.Comment) error {
	if comment.ReplyComment != nil {
		if err := comment.ShowReplyComment(); err != nil {
			return err
		}
	}
	if err := comment.ShowStatus(); err != nil {
		return err
	}
	return nil
}

func handleCommentResult(comments *definitions.CommentResult) error {
	for i := range comments.Comments {
		if err := handleComment(&comments.Comments[i]); err != nil {
			return err
		}
	}
	return nil
}

//ShowComment query weibo specific status' comment
func ShowComment(accessToken string, statusID, count, startCommentID int) (*definitions.CommentResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommentShow, "?access_token=", accessToken, "&id=", strconv.Itoa(statusID), "&count=", strconv.Itoa(count+2), "&max_id=", strconv.Itoa(startCommentID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comments definitions.CommentResult
	if err := json.Unmarshal(response.OriginalBytes, &comments); err != nil {
		return nil, err
	}
	if err := handleCommentResult(&comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

//UserComments query the weibo comments sent by current user
func UserComments(accessToken string, count, startCommentID int) (*definitions.CommentResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommentByMe, "?access_token=", accessToken, "&count=", strconv.Itoa(count+2), "&max_id=", strconv.Itoa(startCommentID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comments definitions.CommentResult
	if err := json.Unmarshal(response.OriginalBytes, &comments); err != nil {
		return nil, err
	}
	if err := handleCommentResult(&comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

//UserRecievedComments query the weibo comments recieved by current user
func UserRecievedComments(accessToken string, count, startCommentID int) (*definitions.CommentResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommentToMe, "?access_token=", accessToken, "&count=", strconv.Itoa(count+2), "&max_id=", strconv.Itoa(startCommentID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comments definitions.CommentResult
	if err := json.Unmarshal(response.OriginalBytes, &comments); err != nil {
		return nil, err
	}
	if err := handleCommentResult(&comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

//GetUserCommentTimeline query the weibo current user's all comments
func GetUserCommentTimeline(accessToken string, count, startCommentID int) (*definitions.CommentResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommentTimeline, "?access_token=", accessToken, "&count=", strconv.Itoa(count+2), "&max_id=", strconv.Itoa(startCommentID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comments definitions.CommentResult
	if err := json.Unmarshal(response.OriginalBytes, &comments); err != nil {
		return nil, err
	}
	if err := handleCommentResult(&comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

//GetUserMentionedComments query the weibo current user's mentioned comments
func GetUserMentionedComments(accessToken string, count, startCommentID int) (*definitions.CommentResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.CommentMentions, "?access_token=", accessToken, "&count=", strconv.Itoa(count+2), "&max_id=", strconv.Itoa(startCommentID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comments definitions.CommentResult
	if err := json.Unmarshal(response.OriginalBytes, &comments); err != nil {
		return nil, err
	}
	if err := handleCommentResult(&comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

//PostComment post a comment to the specific status
func PostComment(accessToken, commentStr string, statusID int) (*definitions.Comment, error) {
	response, err := SimplePost(constant.BaseURI+constant.CommentCreate, map[string]string{"access_token": accessToken, "comment": commentStr, "id": strconv.Itoa(statusID)})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comment definitions.Comment
	if err := json.Unmarshal(response.OriginalBytes, &comment); err != nil {
		return nil, err
	}
	if err := handleComment(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

//RemoveComment remove a comment from the specific status
func RemoveComment(accessToken string, commentID int) (*definitions.Comment, error) {
	response, err := SimplePost(constant.BaseURI+constant.CommentDestroy, map[string]string{"access_token": accessToken, "cid": strconv.Itoa(commentID)})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comment definitions.Comment
	if err := json.Unmarshal(response.OriginalBytes, &comment); err != nil {
		return nil, err
	}
	if err := handleComment(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

//ReplyComment sent the weibo reply comment to a specific status' comment
func ReplyComment(accessToken, commentStr string, statusID, commentID int) (*definitions.Comment, error) {
	response, err := SimplePost(constant.BaseURI+constant.CommentReply, map[string]string{"access_token": accessToken, "comment": commentStr, "cid": strconv.Itoa(commentID), "id": strconv.Itoa(statusID)})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var comment definitions.Comment
	if err := json.Unmarshal(response.OriginalBytes, &comment); err != nil {
		return nil, err
	}
	if err := handleComment(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}
