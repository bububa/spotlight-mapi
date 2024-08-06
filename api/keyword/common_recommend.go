package keyword

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/keyword"
)

// CommonRecommend 定向推词-以词推词
func CommonRecommend(ctx context.Context, clt *core.SDKClient, req *keyword.CommonRecommendRequest, accessToken string) (*keyword.CommonRecommendResult, error) {
	var resp keyword.CommonRecommendResponse
	if err := clt.Post(ctx, "/jg/keyword/common/recommend", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
