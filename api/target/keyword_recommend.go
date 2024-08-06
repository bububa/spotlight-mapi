package target

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/target"
)

// KeywordRecommend 获取推荐关键词信息
func KeywordRecommend(ctx context.Context, clt *core.SDKClient, req *target.KeywordRecommendRequest, accessToken string) ([]target.KeywordRecommend, error) {
	var resp target.KeywordRecommendResponse
	if err := clt.Post(ctx, "/jg/target/keyword/recommend", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
