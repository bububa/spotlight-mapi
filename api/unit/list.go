package unit

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/unit"
)

// List 获取单元列表接口
func List(ctx context.Context, clt *core.SDKClient, req *unit.ListRequest, accessToken string) (*unit.ListResult, error) {
	var resp unit.ListResponse
	if err := clt.Post(ctx, "/jg/unit/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
