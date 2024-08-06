package keyword

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/keyword"
)

// IndustryTaxonomy 行业类目
func IndustryTaxonomy(ctx context.Context, clt *core.SDKClient, req *keyword.IndustryTaxonomyRequest, accessToken string) (*keyword.IndustryTaxonomyResult, error) {
	var resp keyword.IndustryTaxonomyResponse
	if err := clt.Post(ctx, "/jg/keyword/industry/taxonomy", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
