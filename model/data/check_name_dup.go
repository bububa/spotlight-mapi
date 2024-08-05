package data

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CheckNameDupRequest 计划单元名称重复性校验 API Reqeust
type CheckNameDupRequest struct {
	// Name 计划名称/单元名称集合
	// 单次查询上限100
	Name []string `json:"name,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 查询类型：1：计划，2：单元
	Type int `json:"type,omitempty"`
}

// Encode implements PostRequest interface
func (r CheckNameDupRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CheckNameDupData 计划单元名称重复性校验 API Response
type CheckNameDupResponse struct {
	// Data 计划单元名称重复性校验
	Data *CheckNameDupResult `json:"data,omitempty"`
	model.BaseResponse
}

type CheckNameDupResult struct {
	// CheckResult 查询结果
	// key为查询的名称，value为名称是否重复，true：重复，false：不重复
	CheckResult map[string]bool `json:"check_result,omitempty"`
	// Type 查询类型：1：计划，2：单元
	Type int `json:"type,omitempty"`
}
