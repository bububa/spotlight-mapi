package oauth

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/oauth"
)

// AccessToken 获取Token
// 利用授权码 auth_code，请求小红书服务器，获取 access_token 和 refresh_token 及当前账户的广告主 ID。
func AccessToken(ctx context.Context, clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.AccessToken, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp oauth.AccessTokenResponse
	if err := clt.Post(ctx, "/oauth2/access_token", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
