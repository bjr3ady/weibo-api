package definitions

import (
	"encoding/json"
)

//EmotionResult is the weibo common emotion struct
type EmotionResult struct {
	Category string `json:"category"`
	Common bool `json:"common"`
	Hot bool `json:"hot"`
	Icon string `json:"icon"`
	Phrase string `json:"phrase"`
	PicID string `json:"picid"`
	Type string `json:"type"`
	URL string `json:"url"`
	Value string `json:"value"`
}

//Visible is the weibo status' visible struct
type Visible struct {
	Type   int `json:"type"`
	ListID int `json:"list_id"`
}

//Annotation is the weibo status' Annotation struct
type Annotation struct {
	MAPIRequest bool `json:"mapi_request"`
}

//CommentManageInfo is the weibo status' comment management information struct
type CommentManageInfo struct {
	CommentPermissionType int `json:"comment_permission_type"`
	ApprovalCommentType   int `json:"approval_comment_type"`
}

//NumberDisplayStrategy is the weibo status' number display strategy struct
type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
	DisplayTextMinNumber int    `json:"1000000"`
	DisplayText          string `json:"display_text"`
}

//AlchemyParams is the weibo status' alchemy params
type AlchemyParams struct {
	CommentGuideType int `json:"comment_guide_type"`
}

//PicURL is the weibo status' picture url struct
type PicURL struct {
	ThumbnailPic string `json:"thumbnail_pic"`
}

//Geo is the weibo status' geolocation struct
type Geo struct {
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	City         string `json:"city"`
	Province     string `json:"province"`
	CityName     string `json:"city_name"`
	ProvinceName string `json:"province_name"`
	Address      string `json:"address"`
	Pinyin       string `json:"pinyin"`
	More         string `json:"more"`
}

//Status is the weibo status struct
type Status struct {
	Visible                  Visible               `json:"visible"`
	CreatedAt                string                `json:"created_at"`
	ID                       int                   `json:"id"`
	IDStr                    string                `json:"idstr"`
	MID                      string                `json:"mid"`
	CanEdit                  bool                  `json:"can_edit"`
	ShowAdditionalIndication int                   `json:"show_additional_indication"`
	Text                     string                `json:"text"`
	TextLength               int                   `json:"textLength"`
	SourceAllowClick         int                   `json:"source_allowclick"`
	SourceType               int                   `json:"source_type"`
	Source                   string                `json:"source"`
	Favorited                bool                  `json:"favorited"`
	Truncated                bool                  `json:"truncated"`
	InReplyToStatusID        string                `json:"in_reply_to_status_id"`
	InReplyToUserID          string                `json:"in_reply_to_user_id"`
	InReplyToScreenName      string                `json:"in_reply_to_screen_name"`
	PicURLs                  []PicURL              `json:"pic_urls"`
	ThumbnailPic             string                `json:"thumbnail_pic"`
	BMiddlePic               string                `json:"bmiddle_pic"`
	OriginalPic              string                `json:"original_pic"`
	RetweetedStatus          interface{}           `json:"retweeted_status"`
	Geo                      Geo                   `json:"geo"`
	IsPaid                   bool                  `json:"is_paid"`
	MBlogVIPType             int                   `json:"mblog_vip_type"`
	Annotations              []Annotation          `json:"annotations"`
	RepostsCount             int                   `json:"reposts_count"`
	CommentsCount            int                   `json:"comments_count"`
	AttitudesCount           int                   `json:"attitudes_count"`
	PendingApprovalCount     int                   `json:"pending_approval_count"`
	IsLongText               bool                  `json:"isLongText"`
	RewardExhibitionType     int                   `json:"reward_exhibition_type"`
	RewardScheme             string                `json:"reward_scheme"`
	HideFlag                 int                   `json:"hide_flag"`
	MLevel                   int                   `json:"mlevel"`
	BizFeature               int                   `json:"biz_feature"`
	HasActionTypeCard        int                   `json:"hasActionTypeCard"`
	HotWeiboTags             []string              `json:"hot_weibo_tags"`
	TextTagTips              []string              `json:"text_tag_tips"`
	MBlogType                int                   `json:"mblogtype"`
	RID                      string                `json:"rid"`
	UserType                 int                   `json:"userType"`
	User                     interface{}           `json:"user"`
	MoreInfoType             int                   `json:"more_info_type"`
	NumberDisplayStrategy    NumberDisplayStrategy `json:"number_display_strategy"`
	CardID                   string                `json:"cardid"`
	PostitiveRecomFlag       int                   `json:"positive_recom_flag"`
	EnableCommentGuide       bool                  `json:"enable_comment_guide"`
	ContentAuth              int                   `json:"content_auth"`
	GifIDs                   string                `json:"gif_ids"`
	IsShowBulletin           int                   `json:"is_show_bulletin"`
	RepostType               int                   `json:"repost_type"`
	PicNum                   int                   `json:"pic_num"`
	AlchemyParams            AlchemyParams         `json:"alchemy_params"`
	CommentManageInfo        CommentManageInfo     `json:"comment_manage_info"`
}

//TimelineResult is the weibo timeline api return result struct
type TimelineResult struct {
	Statuses          []Status `json:"statuses"`
	HasVisible        bool     `json:"hasvisible"`
	PreviousCursor    int      `json:"previous_cursor"`
	NextCursor        int      `json:"next_cursor"`
	PreviousCursorStr string   `json:"previous_cursor_str"`
	NextCursorStr     string   `json:"next_cursor_str"`
	TotalNumber       int      `json:"total_number"`
	Interval          int      `json:"interval"`
	UVEBlank          int      `json:"uve_blank"`
	SinceID           int      `json:"since_id"`
	SinceIDStr        string   `json:"since_id_str"`
	MaxID             int      `json:"max_id"`
	MaxIDStr          string   `json:"max_id_str"`
	HasUnread         int      `json:"has_unread"`
}

//RepostResult is thw weibo repost
type RepostResult struct {
	Reposts        []Status `json:"reposts"`
	PreviousCursor int      `json:"previous_cursor"`
	NextCursor     int      `json:"next_cursor"`
	TotalNumber    int      `json:"total_number"`
}

//StautsCountResult is the weibo specific status' reposts and comments count querying result
type StautsCountResult struct {
	ID        int `json:"id"`
	Comments  int `json:"comments"`
	Reposts   int `json:"reposts"`
	Attitudes int `json:"attitudes"`
}

//ShowRetweeted unmarshal retweeted status from original bytes
func (sts *Status) ShowRetweeted() error {
	jsonBytes, err := json.Marshal(sts.RetweetedStatus)
	if err != nil {
		return err
	}
	var retweetedStatus Status
	if err := json.Unmarshal(jsonBytes, &retweetedStatus); err != nil {
		return err
	}
	sts.RetweetedStatus = &retweetedStatus
	return nil
}

//ShowUser unmarshal status' user from orignal bytes
func (sts *Status) ShowUser() error {
	jsonBytes, err := json.Marshal(sts.User)
	if err != nil {
		return err
	}
	var user User
	if err := json.Unmarshal(jsonBytes, &user); err != nil {
		return err
	}
	sts.User = &user
	return nil
}
