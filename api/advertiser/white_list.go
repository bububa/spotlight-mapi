package advertiser

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/advertiser"
)

// WHteList 账户白名单
func WhiteList(ctx context.Context, clt *core.SDKClient, req *advertiser.WhiteListRequest, accessToken string) (*advertiser.WhiteList, error) {
	var resp advertiser.WhiteListResponse
	if err := clt.Post(ctx, "/jg/white/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
