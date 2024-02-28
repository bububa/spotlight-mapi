package offline

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/offline"
)

// Advertiser 账户层级离线数据
func Advertiser(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	var resp offline.Response
	if err := clt.Post(ctx, "/jg/data/report/offline/account", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
