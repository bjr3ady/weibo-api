package constant

const (
	//RateLimitStatus 获取当前登录用户的API访问频率限制情况 
	RateLimitStatus = "account/rate_limit_status.json"

	//GetUID OAuth授权之后，获取授权用户的UID 
	GetUID = "account/get_uid.json"

	//GetTokenInfo 查询用户access_token的授权相关信息，包括授权时间，过期时间和scope权限。 
	GetTokenInfo = "oauth2/get_token_info"
)