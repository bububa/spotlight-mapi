package campaign

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/campaign"
)

// Create 创建计划
func Create(ctx context.Context, clt *core.SDKClient, req *campaign.CreateRequest, accessToken string) (uint64, error) {
	var resp campaign.CreateResponse
	if err := clt.Post(ctx, "/jg/campaign/create", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
