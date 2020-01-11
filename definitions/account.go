package definitions

//APIRateLimitStatus is the single api rate limit status
type APIRateLimitStatus struct {
	API           string `json:"api"`
	Limit         int    `json:"limit"`
	LimitTimeUnit string `json:"limit_time_unit"`
	RemainingHits int    `json:"remaining_hits"`
}

//RateLimitStatus is the weibo account rate limit status struct
type RateLimitStatus struct {
	APIRateLimits      []APIRateLimitStatus `json:"api_rate_limits"`
	IPLimit            int                  `json:"ip_limit"`
	LimitTimeUnit      string               `json:"limit_time_unit"`
	RemainingIPHits    int                  `json:"remaining_ip_hits"`
	RemainingUserHits  int                  `json:"remaining_user_hits"`
	ResetTime          string               `json:"reset_time"`
	ResetTimeInSeconds int                  `json:"reset_time_in_seconds"`
	UserLimit          int                  `json:"user_limit"`
}

//UIDResult is the weibo account get authorized user id result struct
type UIDResult struct {
	UID int `json:"uid"`
}

//TokenInfo is the weibo Oauth2 token info
type TokenInfo struct {
	UID       int    `json:"uid"`
	AppKey    string `json:"appkey"`
	Scope     string `json:"scope"`
	CreatedAt string `json:"created_at"`
	ExpiredIn int    `json:"expire_in"`
}
