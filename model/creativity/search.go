package creativity

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// SearchRequest 创意查询 API Request
type SearchRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 单元 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeIDs 创意ID列表
	// 一次查询不超过 20 个，当有创意id时会忽略其他字段
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Status 创意状态
	// 1 已删除2 所有未删除状态3 暂停4 已被单元暂停5 已被计划暂停8 有效9 商品状态异常10 单元未开始11 单元已结束12 单元处于暂停时段
	Status int `json:"status,omitempty"`
	// StartTime 开始时间	start_time 和 end_time同时传或同时不传，过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应 modifier_time 信息(创意更新时间)
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"end_time,omitempty"`
	// Page 分页
	Page *model.PageDTO `json:"page,omitempty"`
}

// Encode implement PostRequest interface
func (r SearchRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SearchResponse 创意查询 API Response
type SearchResponse struct {
	model.BaseResponse
	Data *SearchResult `json:"data,omitempty"`
}

type SearchResult struct {
	// Page 分页
	Page *model.PageRespDTO `json:"page,omitempty"`
	// Creativities 创意列表
	Creativities []Creativity `json:"creativity_dtos,omitempty"`
}
