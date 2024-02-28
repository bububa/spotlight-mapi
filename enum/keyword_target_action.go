package enum

// KeywordTargetAction 关键词定向行为
type KeywordTargetAction int

const (
	// KeywordTargetAction_SEARCH 搜索
	KeywordTargetAction_SEARCH KeywordTargetAction = 1
	// KeywordTargetAction_INTERACTIVE 互动
	KeywordTargetAction_INTERACTIVE KeywordTargetAction = 2
	// KeywordTargetAction_READ 阅读
	KeywordTargetAction_READ KeywordTargetAction = 3
)
