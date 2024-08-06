package target

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/unit"
	"github.com/bububa/spotlight-mapi/util"
)

// GetAvailableTargetInfoRequest 获取定向信息 API Request
type GetAvailableTargetInfoRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingTarget 营销目标 (第一期只支持4)
	// 旧计划 OLD = 0, 应用下载 APP = 1, 销售线索 CLUE = 2, 商品推广 ITEM = 3, 笔记推广 NOTE = 4, 私信营销 MESSAGE = 5, 直播推广 LIVE_BROCAST = 8, 客资收集 USER_DATA_COLLECT = 9,
	MarketingTarget int `json:"marketing_target,omitempty"`
}

// Encode implements PostRequest interface
func (r GetAvailableTargetInfoRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// GetAvailableTargetInfoResponse 获取定向信息 API Response
type GetAvailableTargetInfoResponse struct {
	Data *TargetInfo `json:"data,omitempty"`
	model.BaseResponse
}

// TargetInfo 定向信息
type TargetInfo struct {
	// IndustryInterestTarget 行业兴趣
	IndustryInterestTarget *unit.IndustryInterestTarget `json:"industry_interest_target,omitempty"`
	// ContentInterest 内容兴趣
	ContentInterest *model.CodeNamePair `json:"content_interest,omitempty"`
	// Children 子节点
	Children []model.CodeNamePair `json:"children,omitempty"`
	// ShoppingInterest 购物兴趣
	ShoppingInterest *model.CodeNamePair `json:"shopping_interest,omitempty"`
	// CrowdTarget 人群包
	CrowdTarget *unit.CrowdTarget `json:"crowd_target,omitempty"`
	// GenderTarget 性别
	GenderTargets []model.CodeNamePair `json:"gender_targets,omitempty"`
	// AgeTargets 年龄
	AgeTargets []model.CodeNamePair `json:"age_target,omitempty"`
	// AreaTargets 地域
	AreaTargets []model.CodeNamePair `json:"area_target,omitempty"`
	// DeviceTargets 设备
	DeviceTargets []model.CodeNamePair `json:"device_target,omitempty"`
}
