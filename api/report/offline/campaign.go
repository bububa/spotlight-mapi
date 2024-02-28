package offline

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/offline"
)

// Campaign 广告计划层级离线数据
func Campaign(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	var resp offline.Response
	if err := clt.Post(ctx, "/jg/data/report/offline/campaign", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
