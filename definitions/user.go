package definitions

//Insecurity is the weibo user's information insecurity field
type Insecurity struct {
	SexualContent bool `json:"sexual_content"`
}

//User is the weibo user information struct
type User struct {
	ID int `json:"id"`
	IDStr string `json:"idstr"`
	Class int `json:"class"`
	ScreenName string `json:"screen_name"`
	Name string `json:"name"`
	Province string `json:"province"`
	City string `json:"city"`
	Location string `json:"location"`
	Description string `json:"description"`
	URL string `json:"url"`
	ProfileImageURL string `json:"profile_image_url"`
	CoverImage string `json:"cover_image"`
	CoverImagePhone string `json:"cover_image_phone"`
	ProfileURL string `json:"profile_url"`
	Domain string `json:"domain"`
	Weihao string `json:"weihao"`
	Gender string `json:"gender"`
	FollowersCount int `json:"followers_count"`
	FriendsCount int `json:"friends_count"`
	PageFriendsCount int `json:"pagefirends_count"`
	StatusesCount int `json:"statuses_count"`
	VideoStatusCount int `json:"video_status_count"`
	FavouritesCount int `json:"favourites_count"`
	CreatedAt string `json:"created_at"`
	Following bool `json:"following"`
	AllowAllActMsg bool `json:"allow_all_act_msg"`
	GeoEnabled bool `json:"geo_enabled"`
	Verified bool `json:"verified"`
	VerifiedType int `json:"verified_type"`
	Remark string `json:"remark"`
	Insecurity Insecurity `json:"insecurity"`
	Status Status `json:"status"`
	PType int `json:"ptype"`
	AllowAllComment bool `json:"allow_all_comment"`
	AvatarLarge string `json:"avatar_large"`
	AvatarHD string `json:"avatar_hd"`
	VerifiedReason string `json:"verified_reason"`
	VerifiedTrade string `json:"verified_trade"`
	VerifiedReasonURL string `json:"verified_reason_url"`
	VerifiedSource string `json:"verified_source"`
	VerifiedSourceURL string `json:"verified_source_url"`
	FollowMe bool `json:"follow_me"`
	Like bool `json:"like"`
	LieMe bool `json:"like_me"`
	OnlineStatus int `json:"oneline_status"`
	BiFollowersCount int `json:"bi_followers_count"`
	Lang string `json:"lang"`
	Star int `json:"star"`
	MBType int `json:"mbtype"`
	MBRank int `json:"mbrank"`
	BlockWord int `json:"block_word"`
	BlockApp int `json:"block_app"`
	CreditScore int `json:"credit_score"`
	UserAbility int `json:"user_ability"`
	CardID string `json:"cardid"`
	AvatarGJID string `json:"avatargj_id"`
	URank int `json:"urank"`
	StoryReadState int `json:"story_read_state"`
	VClubMember int `json:"vclub_member"`
	IsTeenager int `json:"is_teenager"`
	IsGuardian int `json:"is_guardian"`
	IsTeenagerList int `json:"is_teenager_list"`
	SpecialFollow bool `json:"special_follow"`
}

//UserRank is the weibo user rank result struct
type UserRank struct {
	UID int `json:"uid"`
	Rank int `json:"rank"`
}