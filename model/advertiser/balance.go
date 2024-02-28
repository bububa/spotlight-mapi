package advertiser

import (
	"strconv"

	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// Balance 账号余额
type Balance struct {
	// TotalBalance 账户余额，单位：分
	TotalBalance int64 `json:"total_balance,omitempty"`
	// CashBalance 现金余额，单位：分
	CashBalance int64 `json:"cash_balance,omitempty"`
	// ReturnBalance 常规返货余额，单位：分
	ReturnBalance int64 `json:"return_balance,omitempty"`
	// FreezeBalance 冻结余额，单位：分
	FreezeBalance int64 `json:"freeze_balance,omitempty"`
	// CompensateReturnBalance 赔付返货余额，单位：分
	CompensateReturnBalance int64 `json:"compensate_return_balance,omitempty"`
	// AvailableBalance 可用金额，单位：分
	AvailableBalance int64 `json:"available_balance,omitempty"`
	// TodaySpend 今日消费，单位：分
	TodaySpend int64 `json:"today_spend,omitempty"`
}

// 获取账号余额 API Request
type BalanceInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r BalanceInfoRequest) Encode() string {
	values := util.NewUrlValues()
	defer util.ReleaseUrlValues(values)
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	return ret
}

// BalanceInfoResponse 获取账号余额 API Response
type BalanceInfoResponse struct {
	model.BaseResponse
	Data *Balance `json:"data,omitempty"`
}
