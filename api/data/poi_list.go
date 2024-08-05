package data

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/data"
)

// PoiList 门店信息列表
func PoiList(ctx context.Context, clt *core.SDKClient, req *data.PoiListRequest, accessToken string) (*data.PoiListResult, error) {
	var resp data.PoiListResponse
	if err := clt.Post(ctx, "/jg/data/poi/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
