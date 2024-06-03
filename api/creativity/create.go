package creativity

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/creativity"
)

// Create 创建笔记创意
func Create(ctx context.Context, clt *core.SDKClient, req *creativity.CreateRequest, accessToken string) (uint64, error) {
	var resp creativity.CreateResponse
	if err := clt.Post(ctx, "/jg/creativity/create", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.CreativityID, nil
}
