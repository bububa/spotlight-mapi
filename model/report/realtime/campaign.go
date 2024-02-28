package realtime

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// CampaignRequest 计划层级实时数据 API Request
type CampaignRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大1000
	PageSize int64 `json:"page_size,omitempty"`
	// MarketTargetingList 营销诉求筛选，3：商品销量_日常推广告，4：产品种草，8：直播推广_日常推广，9：客资收集，10：抢占赛道，14：直播推广_直播预热，15：商品销量_店铺拉新
	MarketTargetingList []int `json:"market_targeting_list,omitempty"`
	// CampaignFilterState 计划状态过滤，1：有效，2：暂停，4：计划预算不足，7：账户日预算不足，5：现金余额不足，8：处于暂停时段，3：已删除，6：所有未删除状态
	CampaignFilterState []int `json:"campaign_filter_state,omitempty"`
	// CampaignCreateBeginTime 计划创建时间范围的开始
	CampaignCreateBeginTime string `json:"campaign_create_begin_time,omitempty"`
	// CampaignCreateEndTime 计划创建时间范围的结束
	CampaignCreateEndTime string `json:"campaign_create_end_time,omitempty"`
	// PlacementLit 广告类型：1：信息流广告，2：搜索广告，4：全站智投，7：视频流广告
	PlacementList []int `json:"placement_list,omitempty"`
	// LimitDayBudgetList 预算类型：0：不限预算，1：指定预算
	LimitDayBudgetList []int `json:"limit_day_budget_list,omitempty"`
	// OptimizeTargetList 推广目标：0：点击量，1：互动量，16：种草值，11：商品访客量，12：落地页访问量，3：表单提交量，4：商品成单量，5：私信咨询量，6：观看量，13：私信开口量，14：有效观看量，17：ROI，站外转化量，20：TI人群规模
	OptimizeTargetList []int `json:"optimize_target_list,omitempty"`
	// BuildTypeList 搭建方式：0：标准投放，1：省心智投
	BuildTypeList []int `json:"build_type_list,omitempty"`
	// BiddingStrategyList 出价方式：2：手动出价，101: 自动出价
	BiddingStrategyList []int `json:"bidding_strategy_list,omitempty"`
	// ConstraintTypeList 成本控制方式：-1: 无，101: 自动控制，0: 点击成本控制，1: 互动成本控制，3: 表单提交成本控制，5: 私信咨询成本控制，11: 访客成本控制，13: 私信开口成本控制，14: 有效观播成本控制，17: ROI控制，23: 预热成本控制，50: 私信留资成本控制
	ConstraintTypeList []int `json:"constraint_type_list,omitempty"`
	// PromotionTargetList 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	PromotionTargetList []int `json:"promotion_target_list,omitempty"`
	// CombineAuditStatus 创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	CombineAuditStatus int `json:"combine_audit_status,omitempty"`
	// MigrateStatusList 计划迁移状态，0：非迁移计划，2：迁移计划
	MigrateStatusList []int `json:"migrate_status_list,omitempty"`
	// Name 搜索计划名称
	Name string `json:"name,omitempty"`
	// ID 搜索计划id
	ID uint64 `json:"id,omitempty"`
}

// Encode implement PostRequest interface
func (r CampaignRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CampaignResponse 计划层级实时数据 API Response
type CampaignResponse struct {
	model.BaseResponse
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// CampaignDTOs 计划数据list
	CampaignDTOs []CampaignDTO `json:"campaign_dtos,omitempty"`
	// TotalData 汇总数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
}

// CampaignDTO 计划数据
type CampaignDTO struct {
	// Data 数据指标
	Data *report.DataReportDTO `json:"data,omitempty"`
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
