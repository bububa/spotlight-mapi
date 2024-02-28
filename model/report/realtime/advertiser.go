package realtime

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// AdvertiserRequest 账户层级实时数据 API Request
type AdvertiserRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
}

// Encode implement PostRequest interface
func (r AdvertiserRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserResponse 账户层级实时数据 API Response
type AdvertiserResponse struct {
	model.BaseResponse
	Data *report.DataReportDTO `json:"data,omitempty"`
}
