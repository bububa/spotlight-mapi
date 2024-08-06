package enum

// KeywordRecommendRequestType 推词类型
type KeywordRecommendRequestType string

const (
	// KeywordRecommendRequestType_NOTE 智能推词-笔记推词
	KeywordRecommendRequestType_NOTE KeywordRecommendRequestType = "note"
	// KeywordRecommendRequestType_INDUSTRY 行业推词
	KeywordRecommendRequestType_INDUSTRY KeywordRecommendRequestType = "industry"
	// KeywordRecommendRequestType_SEARCH 以词推词
	KeywordRecommendRequestType_SEARCH KeywordRecommendRequestType = "search"
	// KeywordRecommendRequestType_SESSION 上下游推词
	KeywordRecommendRequestType_SESSION KeywordRecommendRequestType = "session"
)
