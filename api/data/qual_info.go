package data

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/data"
)

// QualInfo 获取资质列表
func QualInfo(ctx context.Context, clt *core.SDKClient, req *data.QualInfoRequest, accessToken string) (*data.QualInfoResult, error) {
	var resp data.QualInfoResponse
	if err := clt.Post(ctx, "/jg/data/qual/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
