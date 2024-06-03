package unit

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/unit"
)

// Update 编辑单元
func Update(ctx context.Context, clt *core.SDKClient, req *unit.UpdateRequest, accessToken string) (uint64, error) {
	var resp unit.UpdateResponse
	if err := clt.Post(ctx, "/jg/unit/update", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.UnitID, nil
}
