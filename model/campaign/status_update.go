package campaign

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// StatusUpdateRequest 修改计划状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignIDs 计划ID列表
	// 至少传一个限制单次变更计划数量，最多传20
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// ActionType 操作类型
	// 1：开启2：暂停3：删除
	ActionType int `json:"action_type,omitempty"`
}

// Encode implements PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 修改计划状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// CampaignIDs 计划ID列表
		CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	} `json:"data,omitempty"`
}
