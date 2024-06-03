package directlink

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/directlink"
)

// List 获取直达链接
func List(ctx context.Context, clt *core.SDKClient, req *directlink.ListRequest, accessToken string) (*directlink.ListResult, error) {
	var resp directlink.ListResponse
	if err := clt.Get(ctx, "/jg/direct-link/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
