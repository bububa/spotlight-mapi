package unit

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/unit"
)

// UpdateStatus 修改单元状态
func UpdateStatus(ctx context.Context, clt *core.SDKClient, req *unit.UpdateStatusRequest, accessToken string) ([]uint64, error) {
	var resp unit.UpdateStatusResponse
	if err := clt.Post(ctx, "/jg/unit/update/status", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.UnitIDs, nil
}
