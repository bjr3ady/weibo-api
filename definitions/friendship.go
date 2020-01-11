package definitions

//FriendsResult is the weibo querying friends result struct
type FriendsResult struct {
	Users []User `json:"users"`
	NextCursor int `json:"next_cursor"`
	PreviousCursor int `json:"previous_cursor"`
	TotalNumber int `json:"total_number"`
}

//Friendship is the weibo friendship definition struct
type Friendship struct {
	ID int `json:"id"`
	ScreenName string `json:"screen_name"`
	FollowedBy bool `json:"followed_by"`
	Following bool `json:"following"`
	NotificationsEnabled bool `json:"notifications_enabled"`
}

//FriendshipDetail is the weibo querying two friends friendship detail struct
type FriendshipDetail struct {
	Target Friendship `json:"target"`
	Source Friendship `json:"source"`
}