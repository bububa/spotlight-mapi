package directlink

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/directlink"
)

// Delete 删除直达链接
func Delete(ctx context.Context, clt *core.SDKClient, req *directlink.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/direct-link/delete", req, nil, accessToken)
}
