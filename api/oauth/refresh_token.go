package oauth

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/oauth"
)

// RefreshToken 刷新Token
// 请求小红书服务器，刷新 access_token 和 refresh_token 及 token 过期时间。
func RefreshToken(ctx context.Context, clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.AccessToken, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp oauth.AccessTokenResponse
	if err := clt.Post(ctx, "/oauth2/refresh_token", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
