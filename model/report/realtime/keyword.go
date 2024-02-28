package realtime

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// KeywordRequest 关键词层级实时数据 API Request
type KeywordRequest struct {
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
	// KeywordFilterState 关键词状态过滤2：删除3：暂停4：已被单元暂停5：已被计划暂停6：现金余额不足7：计划预算不足8：有效9：计划未开始10：计划已结束11：计划处于暂停阶段12：账户日预算不足
	KeywordFilterState []int `json:"keyword_filter_state,omitempty"`
	// UserBidStrategy 出价策略：0：未使用出价策略1：已使用出价策略
	UserBidStrategy int `json:"user_bid_strategy,omitempty"`
	// KeywordName 搜索关键词名词
	KeywordName string `json:"keyword_name,omitempty"`
	// CampaignName 搜索计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitName 搜索单元名称
	UnitName string `json:"unit_name,omitempty"`
}

// Encode implement PostRequest interface
func (r KeywordRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// KeywordRequest 关键词层级实时数据 API Response
type KeywordResponse struct {
	model.BaseResponse
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// KeywordDTOs 关键词数据list
	KeywordDTOs []KeywordDTO `json:"keyword_dtos,omitempty"`
	// TotalData 汇总数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
}

// KeywordDTO 关键词数据
type KeywordDTO struct {
	// BaseKeywordDTO 创意属性信息
	BaseKeywordDTO *BaseKeywordDTO `json:"base_keyword_dto,omitempty"`
	// SubKeywordDTOs  子关键词信息
	SubKeywordDTOs []KeywordDTO `json:"sub_keyword_dtos,omitempty"`
	// BaseUnitDTO 单元属性信息
	BaseUnitDTO *BaseUnitDTO `json:"base_unit_dto,omitempty"`
	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
	// Data 数据指标
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseKeywordDTO 创意属性信息
type BaseKeywordDTO struct {
	// KeywordID 关键词ID
	KeywordID uint64 `json:"keyword_id,omitempty"`
	// Keyword 关键词名称
	Keyword string `json:"keyword,omitempty"`
	// UserBidStrategy 出价策略：0：未使用出价策略1：已使用出价策略
	UserBidStrategy int `json:"user_bid_strategy,omitempty"`
	// KeywordEnable 关键词状态：0：未上线1：已上线
	KeywordEnable int `json:"keyword_enable,omitempty"`
	// KeywordFilterState 关键词状态过滤2：删除3：暂停4：已被单元暂停5：已被计划暂停6：现金余额不足7：计划预算不足8：有效9：计划未开始10：计划已结束11：计划处于暂停阶段12：账户日预算不足
	KeywordFilterState int `json:"keyword_filter_state,omitempty"`
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
