package realtime

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/realtime"
)

// Unit 单元层级实时数据
func Unit(ctx context.Context, clt *core.SDKClient, req *realtime.UnitRequest, accessToken string) (*realtime.UnitResponse, error) {
	resp := new(realtime.UnitResponse)
	if err := clt.Post(ctx, "/jg/data/report/realtime/unit", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
