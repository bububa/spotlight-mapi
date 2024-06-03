package unit

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateStatusRequest 修改单元状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitIDs 单元id，最多20个
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// Status 单元状态
	// 状态枚举1：开启2：暂停3：删除
	Status int `json:"status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateStatusResponse 修改单元状态 API Response
type UpdateStatusResponse struct {
	model.BaseResponse
	Data struct {
		// UnitIDs 单元id
		UnitIDs []uint64 `json:"unit_ids,omitempty"`
	} `json:"data,omitempty"`
}
