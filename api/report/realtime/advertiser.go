package realtime

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/model/report/realtime"
)

// Advertiser 账户层级实时数据
func Advertiser(ctx context.Context, clt *core.SDKClient, req *realtime.AdvertiserRequest, accessToken string) (*report.Metric, error) {
	var resp realtime.AdvertiserResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/account", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
