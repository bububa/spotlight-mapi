package enum

// LandingPageStatus 落地页状态
type LandingPageStatus int

const (
	// LandingPageStatus_DRAFT 未审核/草稿
	LandingPageStatus_DRAFT LandingPageStatus = 0
	// LandingPageStatus_IN_AUDIT 审核中
	LandingPageStatus_IN_AUDIT LandingPageStatus = 1
	// LandingPageStatus_PASS 审核通过
	LandingPageStatus_PASS LandingPageStatus = 2
	// LandingPageStatus_REJECT 审核未通过
	LandingPageStatus_REJECT LandingPageStatus = 3
	// LandingPageStatus_EXPIRED 已过期
	LandingPageStatus_EXPIRED LandingPageStatus = 4
)
