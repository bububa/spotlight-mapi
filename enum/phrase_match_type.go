package enum

// PhraseMatchType 匹配方式
type PhraseMatchType int

const (
	// PraseMatchType_EXACT 精确匹配
	PhraseMatchType_EXACT PhraseMatchType = 0
	// PhraseMatchType_PHRASE 短语匹配
	PhraseMatchType_PHRASE PhraseMatchType = 1
	// PhraseMatchType_SMART 智能匹配
	PhraseMatchType_SMART PhraseMatchType = 2
)
