package definitions

type FavoriteTag struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type FavoriteItem struct {
	Status        Status        `json:"status"`
	Tags          []FavoriteTag `json:"tags"`
	FavoritedTime string        `json:"favorited_time"`
}

type FavoriteIDItem struct {
	Status        string        `json:"status"`
	Tags          []FavoriteTag `json:"tags"`
	FavoritedTime string        `json:"favorited_time"`
}

//Favorites is the weibo favorites struct
type Favorites struct {
	Favorites   []FavoriteItem `json:"favorites"`
	TotalNumber int            `json:"total_number"`
}

type FavoriteIDs struct {
	Favorites   []FavoriteIDItem `json:"favorites"`
	TotalNumber int              `json:"total_number"`
}
