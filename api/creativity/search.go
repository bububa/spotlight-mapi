package creativity

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/creativity"
)

// Search 创意查询
func Search(ctx context.Context, clt *core.SDKClient, req *creativity.SearchRequest, accessToken string) (*creativity.SearchResult, error) {
	var resp creativity.SearchResponse
	if err := clt.Post(ctx, "/jg/creativity/search", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
