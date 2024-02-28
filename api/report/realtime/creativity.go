package realtime

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/realtime"
)

// Creativity 创意层级实时数据
func Creativity(ctx context.Context, clt *core.SDKClient, req *realtime.CreativityRequest, accessToken string) (*realtime.CreativityResponse, error) {
	resp := new(realtime.CreativityResponse)
	if err := clt.Post(ctx, "/jg/data/report/realtime/creativity", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
