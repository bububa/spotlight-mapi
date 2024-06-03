package campaign

// Campaign 计划
type Campaign struct {
	// CampaignID 	计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 	计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignFilterState 计划状态	1-有效，2-暂停，3-已删除，4-计划预算不足，5-现金余额不足，6-所有未删除状态，7-账户日预算不足
	CampaignFilterState int `json:"campaign_filter_state,omitempty"`
	// CampaignCreateTime 计划创建时间
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// CampaignEnable 计划是否可用
	CampaignEnable int `json:"campaign_enable,omitempty"`
	// MarketingTarget 营销目标	3-商品销量4-产品种草8-直播推广9-客资收集10-抢占关键词13-种草直达14-直播预热15-店铺拉新 16-应用推广
	MarketingTarget int `json:"marketing_target,omitempty"`
	// Placement 广告类型	1-信息流2-搜索推广4-全站智投7-视频内流
	Placement int `json:"placement,omitempty"`
	// OptimizeTarget 优化目标
	OptimizeTarget int `json:"optimize_target,omitempty"`
	// PromotionTarget 投放标的
	PromotionTarget int `json:"promotion_target,omitempty"`
	// BiddingStrategy 出价策略
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// ConstraintType 成本控制类型
	ConstraintType int `json:"constraint_type,omitempty"`
	// ConstraintValue 成本值
	ConstraintValue int `json:"constraint_value,omitempty"`
	// LimitDayBudget 预算类型	0-不限预算，1-指定预算
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// CampaignDayBudget 指定预算
	CampaignDayBudget int64 `json:"campaign_day_budget,omitempty"`
	// BudgetState 推广计划日预算是否充足，	0-不足，1-充足
	BudgetState int `json:"budget_state,omitempty"`
	// SmartSwitch 智能开关
	SmartSwitch int `json:"smart_switch,omitempty"`
	// Platform 创建来源
	Platform int `json:"platform,omitempty"`
	// PacingMode 投放速率	1-匀速2-加速
	PacingMode int `json:"pacing_mode,omitempty"`
	// StartTime 推广开始时间
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 推广结束时间
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriod 推广时间的bitmap
	TimePeriod string `json:"time_period,omitempty"`
	// TimePeroidType 推广时间段类型	0-全时段1-自定义时间段
	TimePeroidType int `json:"time_peroid_type,omitempty"`
	// FeedFlag 是否开启搜索追投,	0-关1-开
	FeedFlat int `json:"feed_flag,omitempty"`
	// BuildType 构建类型
	BuildType int `json:"build_type,omitempty"`
	// CreativityState 总预算达成时间	创意聚合状态，ark电商推广场景使用
	CreativityState int `json:"creativity_state,omitempty"`
	// EventAssetID 事件资产
	EventAssetID uint64 `json:"event_asset_id,omitempty"`
	// AssetEvent 资产事件
	AssetEvent int `json:"asset_event,omitempty"`
	// AssetEventID 资产事件ID
	AssetEventID uint64 `json:"asset_event_id,omitempty"`
	// PageCategory 页面类目
	PageCategory int `json:"page_category,omitempty"`
	// SearchFlag 搜索快投开关
	SearchFlat int `json:"search_flag,omitempty"`
	// SearchBidRatio 定向拓展
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// DeeplinkID deeplink ID
	DeeplinkID uint64 `json:"deeplink_id,omitempty"`
	// UniversalLinkID universal link ID
	UniversalLinkID uint64 `json:"universal_link_id,omitempty"`
	// DetectURLLink detect url link
	DetectURLLink string `json:"detect_url_link,omitempty"`
}
