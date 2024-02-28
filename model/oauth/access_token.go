package oauth

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// AccessToken 授权信息
type AccessToken struct {
	// AccessToken 用于验证权限的 token
	AccessToken string `json:"access_token,omitempty"`
	// AccessTokenExpiresIn access_token 剩余有效时间，单位：秒
	AccessTokenExpiresIn int64 `json:"access_token_expires_in,omitempty"`
	// RefreshToken 用于获取新的 access_token 和 refresh_token，并且刷新过期时间
	RefreshToken string `json:"refresh_token,omitempty"`
	// RefreshTokenExpiresIn refresh_token 剩余有效时间，单位：秒
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in,omitempty"`
	// UserID 授权账号的user_id
	UserID string `json:"user_id,omitempty"`
	// ApprovalRoleType 授权账号类型，4：品牌，601：代理商
	ApprovalRoleType enum.ApprovalRoleType `json:"approval_role_type,omitempty"`
	// RoleType 应用角色类型，1：品牌开发者，2：代理商开发者，3：服务商开发者
	RoleType enum.DeveloperRoleType `json:"role_type,omitempty"`
	// PlatformType 平台类型，1：聚光，2：蒲公英
	PlatformType enum.PlatformType `json:"platform_type,omitempty"`
	// ApprovalAdvertisers 授权广告主账号，品牌开发者与服务商开发者时有值
	ApprovalAdvertisers []Advertiser `json:"approval_advertisers,omitempty"`
}

// AccessTokenRequest 获取Token API Request
type AccessTokenRequest struct {
	model.BaseRequest
	// AuthCode 授权时返回的 auth_code
	AuthCode string `json:"auth_code,omitempty"`
}

// Encode implement PostRequest interface
func (r AccessTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type AccessTokenResponse struct {
	model.BaseResponse
	Data *AccessToken `json:"data,omitempty"`
}

// RefreshTokenRequest 刷新Token API Request
type RefreshTokenRequest struct {
	model.BaseRequest
	// RefreshToken 最近一次小红书返回的 refresh_token
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Encode implement PostRequest interface
func (r RefreshTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
