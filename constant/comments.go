package constant

const (
	//CommentShow 根据微博ID返回某条微博的评论列表 
	CommentShow = "comments/show.json"

	//CommentByMe 获取当前登录用户所发出的评论列表
	CommentByMe = "comments/by_me.json"

	//CommentToMe 获取当前登录用户所接收到的评论列表
	CommentToMe = "comments/to_me.json"

	//CommentTimeline 获取当前登录用户的最新评论包括接收到的与发出的 
	CommentTimeline = "comments/timeline.json"

	//CommentMentions 获取最新的提到当前登录用户的评论，即@我的评论 
	CommentMentions = "comments/mentions.json"

	//CommentShowBatch 根据评论ID批量返回评论信息 
	CommentShowBatch = "comments/show_batch.json"

	//CommentCreate 对一条微博进行评论
	CommentCreate = "comments/create.json"

	//CommentDestroy 删除一条评论
	CommentDestroy = "comments/destroy.json"

	//CommentDestroyBatch 根据评论ID批量删除评论
	CommentDestroyBatch = "comments/destroy_batch.json"

	//CommentReply 回复一条评论
	CommentReply = "comments/reply.json"
)