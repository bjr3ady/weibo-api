package constant

const (
	//HomeTimeline 获取当前登录用户及其所关注用户的最新微博
	HomeTimeline = "statuses/home_timeline.json" 

	//UserTimeline 获取某个用户最新发表的微博列表
	UserTimeline = "statuses/user_timeline.json"

	//RepostTimeline 获取指定微博的转发微博列表
	RepostTimeline = "statuses/repost_timeline.json"

	//BilateralTimeline 获取双向关注用户的最新微博
	BilateralTimeline = "statuses/bilateral_timeline.json"

	//Mentions 获取最新的提到登录用户的微博列表，即@我的微博 
	Mentions = "statuses/memtions.json"

	//StatusShow 根据微博ID获取单条微博内容
	StatusShow = "statuses/show.json"

	//Count 批量获取指定微博的转发数评论数
	Count = "statuses/count.json"

	//Go 根据ID跳转到单条微博页
	Go = "statuses/go"

	//Emotions 获取微博官方表情的详细信息
	Emotions = "emotions.json"

	//Share 第三方分享一条链接到微博
	Share = "statuses/share.json"
)