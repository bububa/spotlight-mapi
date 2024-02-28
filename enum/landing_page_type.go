package enum

// LandingPageType 落地页链接类型
type LandingPageType int

const (
	// LandingPageTYpe_UNSPECIFIED 默认值，无实际意义
	LandingPageType_UNSPECIFIED LandingPageType = 0
	// LandingPageType_FORM 表单
	LandingPageType_FORM LandingPageType = 1
	// LandingPageType_LINK 外跳链接
	LandingPageType_LINK LandingPageType = 2
)
