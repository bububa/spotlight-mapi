package enum

// OAuthScope 权限
type OAuthScope string

const (
	// OAuthScope_REPORT_SERVICE 获取广告账户报表信息） 需要在授权 URL 的 scope 参数上添加 report_service，进行授权后才能访问获取广告账户报表信息的接口。
	OAuthScope_REPORT_SERVICE OAuthScope = "report_service"
	// OAuthScope_AD_QUERY 获取广告计划、广告组、广告创意信息） 需要在授权 URL 的 scope 参数上添加 ad_query 进行授权后才能访问查询广告的接口。
	OAuthScope_AD_QUERY OAuthScope = "ad_query"
	// OAuthScope_AD_MANAGE （创建&修改广告计划、广告组、广告创意） 需要在授权 URL 的 scope 参数上添加 ad_manage 进行授权后才能访问搭建广告的接口。
	OAuthScope_AD_MANAGE OAuthScope = "ad_manage"
)
