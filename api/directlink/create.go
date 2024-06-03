package directlink

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/directlink"
)

// Create 创建直达链接
func Create(ctx context.Context, clt *core.SDKClient, req *directlink.CreateRequest, accessToken string) ([]directlink.DirectLink, error) {
	var resp directlink.CreateResponse
	if err := clt.Post(ctx, "/jg/direct-link/create", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
