package unit

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/unit"
)

// Create 创建单元
func Create(ctx context.Context, clt *core.SDKClient, req *unit.CreateRequest, accessToken string) (uint64, error) {
	var resp unit.CreateResponse
	if err := clt.Post(ctx, "/jg/unit/create", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.UnitID, nil
}
