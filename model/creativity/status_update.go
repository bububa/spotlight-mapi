package creativity

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// StatusUpdateRequest 修改创意状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativityIDs 创意ID, 上限20个
	CreativityIDs []uint64 `json:"creativity_ids,omitempty"`
	// ActionType 操作类型	必填1-开启2-暂停3-删除
	ActionType int `json:"action_type,omitempty"`
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 修改创意状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// CreativityIDs 创意ID
		CreativityIDs []uint64 `json:"creativity_ids,omitempty"`
	} `json:"data,omitempty"`
}
