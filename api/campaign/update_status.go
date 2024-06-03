package campaign

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/campaign"
)

// StatusUpdate 修改计划状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *campaign.StatusUpdateRequest, accessToken string) ([]uint64, error) {
	var resp campaign.StatusUpdateResponse
	if err := clt.Post(ctx, "/jg/campaign/status/update", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CampaignIDs, nil
}
