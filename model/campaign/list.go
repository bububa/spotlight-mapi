package campaign

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 查询计划列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignIDs 计划ID列表
	// 至少传一个限制单次查询计划数量，最多传20
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// StartTime 开始时间
	// 示例：2023-09-20，和expire_time配套填写，创建时间查询范围-开始时间
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 结束时间
	// 示例：2023-09-21，start_time配套填写，创建时间查询范围-结束时间
	ExpireTime string `json:"expire_time,omitempty"`
	// Status 创意状态
	// 1-有效 2-暂停 3-已删除 4-计划预算不足，5-现金余额不足，7-账户日预算不足 8-处于暂停阶段
	Status int `json:"status,omitempty"`
	// Page
	Page *model.PageDTO `json:"page,omitempty"`
}

// Encode implements PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 查询计划列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// Page 页面和总数
	Page *model.PageRespDTO `json:"page,omitempty"`
	// Campaigns 计划查询结果
	Campaigns []Campaign `json:"base_campaign_dtos,omitempty"`
}
