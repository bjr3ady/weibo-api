package definitions

import (
	"encoding/json"
)

//Comment is the weibo status' comment struct
type Comment struct {
	CreatedAt    string      `json:"created_at"`
	ID           int         `json:"id"`
	Text         string      `json:"text"`
	Source       string      `json:"source"`
	User         User        `json:"user"`
	MID          string      `json:"mid"`
	IDStr        string      `json:"idstr"`
	Status       Status      `json:"status"`
	ReplyComment interface{} `json:"reply_comment"`
}

//ShowReplyComment unmarshal replay comment bytes
func (comment *Comment) ShowReplyComment() error {
	jsonBytes, err := json.Marshal(comment.ReplyComment)
	if err != nil {
		return err
	}
	var replyComment Comment
	if err := json.Unmarshal(jsonBytes, &replyComment); err != nil {
		return err
	}
	comment.ReplyComment = &replyComment
	return nil
}

//ShowStatus bind comment related status object
func (comment *Comment) ShowStatus() error {
	if err := comment.Status.ShowRetweeted(); err != nil {
		return err
	}
	if err := comment.Status.ShowUser(); err != nil {
		return err
	}
	return nil
}

//CommentResult is the weibo status' comments collection struct
type CommentResult struct {
	Comments       []Comment `json:"comments"`
	PreviousCursor int       `json:"previous_cursor"`
	NextCursor     int       `json:"next_cursor"`
	TotalNumber    int       `json:"total_number"`
}
