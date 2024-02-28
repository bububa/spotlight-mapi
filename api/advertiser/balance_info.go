package advertiser

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/advertiser"
)

// BalanceInfo 获取账号余额接口
func BalanceInfo(ctx context.Context, clt *core.SDKClient, req *advertiser.BalanceInfoRequest, accessToken string) (*advertiser.Balance, error) {
	var resp advertiser.BalanceInfoResponse
	if err := clt.Get(ctx, "/jg/account/balance/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
