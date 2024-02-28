package enum

// KeywordGenType 单元选词方式
type KeywordGenType int

const (
	// KeywordGenType_UNSPECIFIED 无意义默认值
	KeywordGenType_UNSPECIFIED KeywordGenType = -1
	// KeywordGenType_MANUAL 手动选词
	KeywordGenType_MANUAL KeywordGenType = 0
	// KeywordGenType_INTELLIGENT 智能拓词
	KeywordGenType_INTELLIGENT KeywordGenType = 1
	// KeywordGenType_MANUAL_INTELLIGENT 手动+智能
	KeywordGenType_MANUAL_INTELLIGENT KeywordGenType = 2
)
