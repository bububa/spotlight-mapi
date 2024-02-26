package realtime

import "github.com/bububa/spotlight-mapi/model/report"

type CampaignDTO struct {
	// Data 数据指标
	Data *report.Metric `json:"data,omitempty"`
	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
}

// BaseCampaignDTO 计划属性信息
type BaseCampaignDTO struct {
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignFilterState 计划状态;
	CampaignFilterState int `json:"campaign_filter_state,omitempty"`
	// CampaignCreateTime 计划创建时间: 格式yyyy-MM-dd HH:mm:ss
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// CampaignEnabled 计划启停状态：0：暂停，1：开启
	CampaignEnabled int `json:"campaign_enabled,omitempty"`
	// MarketingTarget 营销诉求:3：商品销量_日常推广，4：产品种草，8：直播推广_日常推广，9：客资收集，10：抢占赛道，14：直播推广_直播预热，15：商品销量_店铺拉新
	MarketingTarget int `json:"marketing_target,omitempty"`
	// Placement 广告类型:1：信息流，2：搜索，4：全站智投，7：视频内流
	Placement int `json:"placement,omitempty"`
	// OptimizeTarget 推广目标：0：点击量1：互动量3：表单提交量4：商品成单量5：私信咨询量6：直播间观看量11：商品访客量12：落地页访问量13：私信开口量14：有效观看量18：站外转化量20：TI人群规模21：行业商品成单23：直播预热量24：直播间成交25：直播间支付ROI
	OptimizeTarget int `json:"optimize_target,omitempty"`
	// PromotionTarget 投放标的:1：笔记，2：商品，7：外链落地页，9：落地页，18：直播间
	PromotionTarget int `json:"promotion_target,omitempty"`
	// BiddingStrategy 出价方式：2：手动出价3：自动出价
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// ConstraitType 成本控制方式: -1: 无，101: 自动控制，0: 点击成本控制，1: 互动成本控制，3: 表单提交成本控制，5: 私信咨询成本控制，11: 访客成本控制，13: 私信开口成本控制，14: 有效观播成本控制，17: ROI控制，23: 预热成本控制，50: 私信留资成本控制
	ConstraitType int `json:"constrait_type,omitempty"`
	// LimitDayBudget 预算类型: 预算类型：0：不限预算，1：指定预算
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 计划日预算
	OriginCampaignDayBudget int `json:"origin_campaign_day_budget,omitempty"`
	// BudgetState 预算状态，0: 计划预算不足，1 计划预算充足
	BudgetState int `json:"budget_state,omitempty"`
	// SmartSwitch 是否节假日预算上调，0: 关闭，1: 开启
	SmartSwitch int `json:"smart_switch,omitempty"`
	// PacingMode 投放速率，1: 匀速投放，2: 加速投放
	PacingMode int `json:"pacing_mode,omitempty"`
	// StartTime 计划开始时间：格式yyyy-MM-dd
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 计划结束时间：格式yyyy-MM-dd
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriod 时段: 默认168个1：表示一周每个小时用0和1表示，0表示不投，1表示投放，示例中表示1点不投，其他时间投
	TimePeriod string `json:"time_period,omitempty"`
	// TimePeroidType 推广时段类型, 0: 全时段，1:自定义时间段
	TimePeroidType int `json:"time_peroid_type,omitempty"`
	// BuildType 搭建方式，0：标准搭建，1：省心智投
	BuildType int `json:"build_type,omitempty"`
	// FeedFlag 是否搜索追投信息流：0: 否，1：是
	FeedFlag int `json:"feed_flag,omitempty"`
	// SearchFlag 是否信息流快投搜索：0: 否，1：是
	SearchFlag int `json:"search_flag,omitempty"`
	// MigrationStatus 专业号平台计划迁移状态: 0：非迁移计划，2：迁移计划
	MigrationStatus int `json:"migration_status,omitempty"`
}
