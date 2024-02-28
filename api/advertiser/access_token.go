package advertiser

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/oauth"
)

// AccessToken 获取accessToken
func AccessToken(ctx context.Context, clt *core.SDKClient, advertiserID uint64) (*oauth.AccessToken, error) {
	req := model.CommonRequest{
		AdvertiserID: advertiserID,
		Method:       "oauth.getAccessToken",
	}
	var ret oauth.AccessTokenResponse
	if err := clt.Post(ctx, "/common", req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
