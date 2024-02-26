package enum

// DeveloperRoleType 应用角色类型
type DeveloperRoleType int

const (

	// DeveloperRoleType_BRAND 品牌开发者
	DeveloperRoleType_BRAND DeveloperRoleType = 1
	// DeveloperRoleType_AGENCY 代理商开发者，
	DeveloperRoleType_AGENCY DeveloperRoleType = 2
	// DeveloperRoleType_SERVICE_PROVIDER 服务商开发者
	DeveloperRoleType_SERVICE_PROVIDER DeveloperRoleType = 3
)
