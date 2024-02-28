package realtime

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// UnitRequest 单元层级实时数据 API Request
type UnitRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id"`
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
	// UnitFilterState 单元状态过滤，10：有效，4：暂停，2：未开始，3： 已结束，5：处于暂停时段，6：已被计划暂停，8：计划预算不足，11：账户日预算不足，7：现金余额不足，1：已删除，9：所有未删除状态
	UnitFilterState []int `json:"unit_filter_state,omitempty"`
	// UnitCreateBeginTime 单元创建时间范围的开始
	UnitCreateBeginTime string `json:"unit_create_begin_time,omitempty"`
	// UnitCreateEndTime 单元创建时间范围的结束
	UnitCreateEndTime string `json:"unit_create_end_time,omitempty"`
	// PlacementLit 广告类型：1：信息流广告，2：搜索广告，4：全站智投，7：视频流广告
	PlacementList []int `json:"placement_list,omitempty"`
	// BiddingStrategyList 出价方式：2：手动出价，101: 自动出价
	BiddingStrategyList []int `json:"bidding_strategy_list,omitempty"`
	// PromotionTargetList 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	PromotionTargetList []int `json:"promotion_target_list,omitempty"`
	// CombineAuditStatus 创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	CombineAuditStatus int `json:"combine_audit_status,omitempty"`
	// Name 搜索单元名称
	Name string `json:"name,omitempty"`
	// ID 单元ID
	ID uint64 `json:"id,omitempty"`
}

// Encode implement PostRequest interface
func (r UnitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UnitRequest 单元层级实时数据 API Response
type UnitResponse struct {
	model.BaseResponse
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// UnitDTOs 单元数据list
	UnitDTOs []UnitDTO `json:"unit_dtos,omitempty"`
	// TotalData 汇总数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
}

// UnitDTO 单元数据
type UnitDTO struct {
	// BaseUnitDTO 单元属性信息
	BaseUnitDTO *BaseUnitDTO `json:"base_unit_dto,omitempty"`
	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
	// Data 数据指标
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseUnitDTO 单元属性信息
type BaseUnitDTO struct {
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// UnitCreateTime 单元创建时间：格式 yyyy-MM-dd HH:mm:ss
	UnitCreateTime string `json:"unit_create_time,omitempty"`
	// UnitEnable 单元启停状态：0：暂停，1：开启
	UnitEnable int `json:"unit_enable,omitempty"`
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// EventBid 出价
	EventBid int64 `json:"event_bid,omitempty"`
}
