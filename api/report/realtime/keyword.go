package realtime

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/report/realtime"
)

// Keyword 关键词层级实时数据
func Keyword(ctx context.Context, clt *core.SDKClient, req *realtime.KeywordRequest, accessToken string) (*realtime.KeywordResponse, error) {
	resp := new(realtime.KeywordResponse)
	if err := clt.Post(ctx, "/jg/data/report/realtime/keyword", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
