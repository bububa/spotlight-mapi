package creativity

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/creativity"
)

// Update 编辑创意
func Update(ctx context.Context, clt *core.SDKClient, req *creativity.UpdateRequest, accessToken string) (uint64, error) {
	var resp creativity.UpdateResponse
	if err := clt.Post(ctx, "/jg/creativity/update", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.CreativityID, nil
}
