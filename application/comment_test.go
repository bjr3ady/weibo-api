package application

import (
	"testing"
)

func TestGetComment(t *testing.T) {
	result, err := ShowComment(TestToken, 4431273584461063, 10, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserComments(t *testing.T) {
	result, err := UserComments(TestToken, 10, 4458201321348751)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserRecievedComments(t *testing.T) {
	result, err := UserRecievedComments(TestToken, 10, 4457774240791000)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserCommentTimeline(t *testing.T) {
	result, err := GetUserCommentTimeline(TestToken, 10, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestGetUserMentionedComments(t *testing.T) {
	result, err := GetUserMentionedComments(TestToken, 10, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestPostComment(t *testing.T) {
	result, err := PostComment(TestToken, "!", 4459592379039380)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestRemoveComment(t *testing.T) {
	result, err := RemoveComment(TestToken, 4459596301042681)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestReplayComment(t *testing.T) {
	result, err := ReplyComment(TestToken, "?", 4459592379039380, 4459596301042681)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}