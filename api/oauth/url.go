package oauth

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/oauth"
	"github.com/bububa/spotlight-mapi/util"
)

// URL 生成oauth授权链接
func URL(ctx context.Context, clt *core.SDKClient, req *oauth.URLRequest) string {
	req.AppID = clt.AppID()
	return util.StringsJoin(core.OAUTH_URL, "?", req.Encode())
}
