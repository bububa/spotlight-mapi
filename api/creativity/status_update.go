package creativity

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/creativity"
)

// StatusUpdate 修改创意状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *creativity.StatusUpdateRequest, accessToken string) ([]uint64, error) {
	var resp creativity.StatusUpdateResponse
	if err := clt.Post(ctx, "/jg/creativity/status/update", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CreativityIDs, nil
}
