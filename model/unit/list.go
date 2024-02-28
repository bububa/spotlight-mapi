package unit

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 获取单元列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitIDs 单元id,不超过10个
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// Status 投放状态：1：投放中2：暂停
	Status int `json:"status,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// StartDate 创建单元开始时间
	StartDate string `json:"start_date,omitempty"`
	// EndDate 创建单元结束时间
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码数，默认1
	Page int `json:"page,omitempty"`
	// PageSize 请求的每页行数，默认20
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 获取单元列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// UnitInfos 单元信息
	UnitInfos []Unit `json:"unit_infos,omitempty"`
}
