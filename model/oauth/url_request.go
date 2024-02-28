package oauth

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// URlRequest OAuth授权链接请求
type URLRequest struct {
	model.BaseRequest
	// Scope 可在各功能模块文档中获取对应的权限 code，参数值原始格式："report_service","xxxxxxx","……"，使用时需要 UrlEncode 一次"report_service"
	Scope []enum.OAuthScope `json:"scope,omitempty"`
	// RedirectURI 申请应用时开发者提供的回调地址，使用时需要 UrlEncode 一次
	RedirectURI string `json:"redirect_uri,omitempty"`
	// State 回调时会原样返回，可用于广告主区分不同投放渠道等用途，广告主可选择性使用
	State string `json:"state,omitempty"`
}

// Encode implement GetRequest interface
func (r URLRequest) Encode() string {
	values := util.NewUrlValues()
	defer util.ReleaseUrlValues(values)
	values.Set("app_id", r.AppID)
	values.Set("scope", string(util.JSONMarshal(r.Scope)))
	values.Set("redirect_uri", r.RedirectURI)
	if r.State != "" {
		values.Set("state", r.State)
	}
	ret := values.Encode()
	return ret
}
