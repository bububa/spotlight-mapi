package realtime

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/realtime"
)

// Campaign 计划层级实时数据
func Campaign(ctx context.Context, clt *core.SDKClient, req *realtime.CampaignRequest, accessToken string) (*realtime.CampaignResponse, error) {
	resp := new(realtime.CampaignResponse)
	if err := clt.Post(ctx, "/jg/data/report/realtime/campaign", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
