package realtime

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// CreativityRequest 创意层级实时数据 API Request
type CreativityRequest struct {
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
	// PlacementLit 广告类型：1：信息流广告，2：搜索广告，4：全站智投，7：视频流广告
	PlacementList []int `json:"placement_list,omitempty"`
	// CreativeityFilterState 创意状态过滤，8：有效，3：暂停，9：商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被计划暂停，13：计划预算不足，16：账户日预算不足，14：现金余额不足，1：已删除，2：所有未删除状态，
	CreativityFilterState int `json:"creativity_filter_state,omitempty"`
	// CreativityCreateBeginTime 创意创建开始时间
	CreativityCreateBeginTime string `json:"creativity_create_begin_time,omitempty"`
	// CreativityCreateEndTime 创意创建结束时间
	CreativityCreateEndTime string `json:"creativity_create_end_time,omitempty"`
	// ConversionType 创意类型：30：商品，20：落地页，4：直播间笔记，7：直播间，0：笔记（无组件），1：笔记（商品组件），2：笔记（落地页组件）
	ConversionType int `json:"conversion_type,omitempty"`
	// ProgrammmaticList 创意组合方式，0：自定义创意，1：程序化创意
	ProgrammaticList []int `json:"programmatic_list,omitempty"`
	// CreativityAuditStatus 创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	CreativityAuditStatus []int `json:"creativity_audit_status,omitempty"`
	// Name 搜索创意名称
	Name string `json:"name,omitempty"`
	// ID 创意ID
	ID uint64 `json:"id,omitempty"`
}

// Encode implement PostRequest interface
func (r CreativityRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreativityResponse 创意层级实时数据 API Response
type CreativityResponse struct {
	model.BaseResponse
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// CreativityDTOs 创意数据list
	CreativityDTOs []CreativityDTO `json:"creativity_dtos,omitempty"`
	// TotalData 汇总数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
}

// CreativityDTO 创意数据
type CreativityDTO struct {
	// BaseCreativtyDTO 创意属性信息
	BaseCreativtyDTO *BaseCreativityDTO `json:"base_creativty_dto,omitempty"`
	// BaseUnitDTO 单元属性信息
	BaseUnitDTO *BaseUnitDTO `json:"base_unit_dto,omitempty"`
	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
	// Data 数据指标
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseCreativityDTO 创意属性信息
type BaseCreativityDTO struct {
	// CreativityID 创意ID
	CreativityID uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称
	CreativityName string `json:"creativity_name,omitempty"`
	// CreativityFilterState 创意状态8：有效，3：暂停，9: 商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被计划暂停，13：计划预算不足，16：账户日预算不足，14：现金余额不足，1：已删除
	CreativityFilterState int `json:"creativity_filter_state,omitempty"`
	// CreativityCreateTime 创意创建时间：格式 yyyy-MM-dd HH:mm:ss
	CreativityCreateTime string `json:"creativity_create_time,omitempty"`
	// CreativityEnable 创意启停状态：0：暂停，1：开启
	CreativeEnable int `json:"creative_enable,omitempty"`
	// AuditStatus 审核状态1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密)
	AuditStatus int `json:"audit_status,omitempty"`
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// Programmatic 创意组合类型：0：自定义创意，1：程序化创意
	Programmatic int `json:"programmatic,omitempty"`
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// CreativityType 创意类型：
	// 0：笔记-无组件
	// 1：笔记-商品组件-购买同款商品
	// 2：笔记-商品组件-进店看看
	// 3：笔记-商品组件-小程序购买同款商品
	// 4：笔记-商品组件-小程序购买同款商品
	// 5：笔记-落地页组件-表单
	// 6：笔记-落地页组件-外跳链接
	// 7：笔记-私信组件
	// 8：笔记-直播间组件
	// 9：笔记-poi门店组件
	// 10：笔记-外链商品
	// 11：直播间
	// 12：搜索组件
	// 13：小程序组件
	// 14：留资组件
	CreativityType int `json:"creativity_type,omitempty"`
}
