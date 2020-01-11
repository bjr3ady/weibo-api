package definitions

//RemindResult is the weibo remind result strcut
type RemindResult struct {
	Status             int `json:"status"`
	Follower           int `json:"follower"`
	Comments           int `json:"cmt"`
	DM                 int `json:"dm"`
	MentionStatus      int `json:"mention_status"`
	MentionComments    int `json:"mention_cmt"`
	Group              int `json:"group"`
	PrivateGroup       int `json:"private_group"`
	Notice             int `json:"notice"`
	Invite             int `json:"invite"`
	Badge              int `json:"badge"`
	Photo              int `json:"photo"`
	AllMentionStatus   int `json:"all_mention_status"`
	AllMentionComments int `json:"all_mention_cmt"`
	AllComments        int `json:"all_cmt"`
	AllFollowers       int `json:"all_follower"`
	LikeUser           int `json:"likeuser"`
	LikeFriends        int `json:"addfriends"`
	MsgBox             int `json:"msgbox"`
}
