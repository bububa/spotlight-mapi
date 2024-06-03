package campaign

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateRequest 编辑计划 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称（长度不超过50个字符）
	CampaignName string `json:"campaign_name,omitempty"`
	// LimitDayBudget 预算类型，0：不限预算，1：指定预算
	LimitDayBudget *int `json:"limit_day_budget"`
	// CampaignDayBudget 计划日预算，单位分，范围 [10000~99999900)，>=10000, < 99999900
	// 指定预算时设置
	CampaignDayBudget int64 `json:"campaign_day_budget,omitempty"`
	// SmartSwitch 节假日预算上浮, 0:关闭，1:开启
	// 不限预算不支持节假日上浮(不传或传0都可）
	SmartSwitch *int `json:"smart_switch,omitempty"`
	// TimeType 推广时时间类型,，0:长期投放，1:自定义设置开始结束时间
	TimeType *int `json:"time_type"`
	// StartTime 推广开始时间，格式 yyyy-MM-dd
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 推广结束时间，格式 yyyy-MM-dd
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeroidType 推广时段类型, 0: 全时段，1:自定义时间段
	TimePeroidType *int `json:"time_peroid_type"`
	// TimePeriod 高级设置-自定义时间段（全时段可不传）
	TimePeriod *TimePeriod `json:"time_period,omitempty"`
	// PacingMode 投放速率, 1:匀速，2:加速
	// 不限预算可不填，默认是加速，自动出价-成本自动控制，固定匀速
	PacingMode int `json:"pacing_mode,omitempty"`
	// FeedFlag 搜索追投
	// 是否开启搜索追投：（搜索推广-普通投放下，互动成本控制和手动出价下支持搜索追投功能）0：关闭1：开启
	FeedFlag *int `json:"feed_flag,omitempty"`
	// BiddingStrategy 出价方式，2: 手动出价，3: 自动出价-成本自动控制，7: 自动出价-成本手动控制
	// 自动出价-成本手动控制时设置（即bidding_strategy=7）唤端场景下传7
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// SearchFlag 搜索快投
	// 0：开启后关闭，1：开启
	SearchFlat *int `json:"search_flag,omitempty"`
	// TargetExtensionSwitch 搜索快投定向拓展
	TargetExtensionSwitch *int `json:"target_extension_switch,omitempty"`
	// SearchBidRatio 搜索快投-出价系数
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// DeeplinkID deeplink链接id
	// 唤端场景下需要
	DeeplinkID uint64 `json:"deeplink_id,omitempty"`
	// UniversalLinkID ulk的链接id
	// 唤端场景下需要
	UniversalLinkID uint64 `json:"universal_link_id,omitempty"`
	// DetectURLLink 监测链接
	// 唤端场景下需要关注。optimize_target如果是35、36、37、38，则必填
	DetectURLLink string `json:"detect_url_link,omitempty"`
}

// Encode implements PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 编辑计划 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data struct {
		// CampaignID 	计划id
		CampaignID uint64 `json:"campaign_id,omitempty"`
	} `json:"data,omitempty"`
}
