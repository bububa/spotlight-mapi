package target

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/target"
)

// KeywordMatch 获取关键词匹配词库信息
func KeywordMatch(ctx context.Context, clt *core.SDKClient, req *target.KeywordMatchRequest, accessToken string) (*target.KeywordMatchResult, error) {
	var resp target.KeywordMatchResponse
	if err := clt.Post(ctx, "/jg/target/keyword/match", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
