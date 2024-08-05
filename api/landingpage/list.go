package landingpage

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/landingpage"
)

// List 落地页表单查询
func List(ctx context.Context, clt *core.SDKClient, req *landingpage.ListRequest, accessToken string) (*landingpage.ListResult, error) {
	var resp landingpage.ListResponse
	if err := clt.Post(ctx, "/jg/landing_page/list_landing_page", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
