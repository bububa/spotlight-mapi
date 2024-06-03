package campaign

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CreateRequest 创建计划 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingTarget 营销目标（4：笔记种草，9：客资收集，16：应用推广）
	MarketingTarget int `json:"marketing_target,omitempty"`
	// CampaignName 计划名称（长度不超过50个字符）
	CampaignName string `json:"campaign_name,omitempty"`
	// Placement 广告类型（1：信息流，2：搜索推广））
	Placement int `json:"placement,omitempty"`
	// PromotionTarget 推广标的类型，（1：笔记）
	PromotionTarget int `json:"promotion_target,omitempty"`
	// Enable 计划创建后默认开启状态，1-开启，0-不开启，不传默认开启
	Enable *int `json:"enable,omitempty"`
	// TimeType 推广时时间类型,，0:长期投放，1:自定义设置开始结束时间
	TimeType int `json:"time_type"`
	// StartTime 推广开始时间，格式 yyyy-MM-dd
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 推广结束时间，格式 yyyy-MM-dd
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeroidType 推广时段类型, 0: 全时段，1:自定义时间段
	TimePeroidType int `json:"time_peroid_type"`
	// TimePeriod 高级设置-自定义时间段（全时段可不传）
	TimePeriod *TimePeriod `json:"time_period,omitempty"`
	// BiddingStrategy 出价方式，2: 手动出价，3: 自动出价-成本自动控制，7: 自动出价-成本手动控制
	// 自动出价-成本手动控制时设置（即bidding_strategy=7）唤端场景下传7
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// LimitDayBudget 预算类型，0：不限预算，1：指定预算
	LimitDayBudget int `json:"limit_day_budget"`
	// CampaignDayBudget 计划日预算，单位分，范围 [10000~99999900)，>=10000, < 99999900
	// 指定预算时设置
	CampaignDayBudget int64 `json:"campaign_day_budget,omitempty"`
	// OptimizeTarget 推广目标，0: 点击量，1: 互动量，35:APP打开，36：APP进店37：APP互动 38：APP成交43：APP打开按钮点击量
	// 唤端只需要关注35、36、37、38和43
	OptimizeTarget int `json:"optimize_target"`
	// ConstraintType 成本控制类型，0: 点击成本控制，1: 互动成本控制35:APP打开，36：APP进店37：APP互动 38：APP成交43：APP打开按钮点击量
	// 唤端只需要关注35、36、37、38和43
	ConstraintType *int `json:"constraint_type,omitempty"`
	// SmartSwitch 节假日预算上浮, 0:关闭，1:开启
	// 不限预算不支持节假日上浮(不传或传0都可）
	SmartSwitch *int `json:"smart_switch,omitempty"`
	// PacingMode 投放速率, 1:匀速，2:加速
	// 不限预算可不填，默认是加速，自动出价-成本自动控制，固定匀速
	PacingMode int `json:"pacing_mode,omitempty"`
	// FeedFlag 搜索追投
	// 是否开启搜索追投：（搜索推广-普通投放下，互动成本控制和手动出价下支持搜索追投功能）0：关闭1：开启
	FeedFlag *int `json:"feed_flag,omitempty"`
	// BuildType 搭建类型
	// 0：普通搭建 （标准搭建）1：智能搭建 （省心智投）默认为0（普通搭建），唤端下支持智能搭建
	BuildType *int `json:"build_type,omitempty"`
	// EventAssetID 资产id
	// 默认0见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_asset_id.唤端下需要关注
	EventAssetID uint64 `json:"event_asset_id,omitempty"`
	// AssetEvent 资产事件类型，这里和推广目标一一对应。401：APP打开（唤起）402：APP进店（唤起）403：APP互动（唤起）404：APP成交-订单数（唤起）
	// 默认0。见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_type.唤端下需要关注
	AssetEvent int `json:"asset_event,omitempty"`
	// AssetEventID 资产事件id
	// 默认0见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_id.唤端下需要关注
	AssetEventID uint64 `json:"asset_event_id,omitempty"`
	// PageCategory 落地页类型
	// 默认01：聚光落地页2：自研落地页3：原生落地页
	PageCategory int `json:"page_category,omitempty"`
	// SearchFlag 搜索快投
	// 0：开启后关闭，1：开启
	SearchFlat *int `json:"search_flag,omitempty"`
	// TargetExtensionSwitch 搜索快投定向拓展
	TargetExtensionSwitch int `json:"target_extension_switch,omitempty"`
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

// TimePeriod 自定义时间段
type TimePeriod struct {
	// Mon 星期一
	// 默认24个1：111111111111111111111111示例：101111111111111111111111每个小时用0和1表示，0表示不投，1表示投放，示例中表示1点不投，其他时间投以下星期同理
	Mon string `json:"mon,omitempty"`
	// Tues 星期二
	Tues string `json:"tues,omitempty"`
	// Wed 	星期三
	Wed string `json:"wed,omitempty"`
	// Thu 星期四
	Thu string `json:"thu,omitempty"`
	// Fri 星期五
	Fri string `json:"fri,omitempty"`
	// Sat 星期六
	Sat string `json:"sat,omitempty"`
	// Sun 星期日
	Sun string `json:"sun,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建计划 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// CampaignID 	计划id
		CampaignID uint64 `json:"campaign_id,omitempty"`
	} `json:"data,omitempty"`
}
