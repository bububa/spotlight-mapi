package conversion

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/conversion"
)

// App APP口令码数据回传
func App(ctx context.Context, clt *core.SDKClient, req *conversion.AppRequest) error {
	return clt.Post(ctx, "/app", req, nil, "")
}
