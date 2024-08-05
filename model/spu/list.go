package spu

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 获取spu列表 API Request
type ListRequest struct {
	// Keyword 搜索词
	Keyword string `json:"keyword,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 请求的页码，默认为 1
	Page int64 `json:"page,omitempty"`
	// PageSize 每页行数，默认为 10
	PageSize int64 `json:"page_size,omitempty"`
	// CanBind 是否可以绑定
	CanBind bool `json:"can_bind,omitempty"`
}

// Encode implements PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 获取spu列表 API Response
type ListResponse struct {
	Data *ListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ListResult struct {
	// List 笔记信息
	List []Spu `json:"spu,omitempty"`
	// Page 当前页
	Page int64 `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int64 `json:"page_size,omitempty"`
	// Total 总数
	Total int64 `json:"total,omitempty"`
}

// Spu 笔记信息
type Spu struct {
	// SpuName 名称
	SpuName string `json:"spu_name,omitempty"`
	// TaxonomyCode spu类目code
	TaxonomyCode string `json:"taxonomy_code,omitempty"`
	// BrandID 品牌id
	BrandID string `json:"brand_id,omitempty"`
	// SpuAuditReason 审核拒绝原因
	SpuAuditReason string `json:"spu_audit_reason,omitempty"`
	// NickNameList spu昵称列表
	NickNameList []string `json:"nick_name_list,omitempty"`
	// SeriesList spu系列名称列表
	SeriesList []string `json:"series_list,omitempty"`
	// PicURLList spu图片链接列表
	PicURLList []string `json:"pic_url_list,omitempty"`
	// MainSpuID
	MainSpuID uint64 `json:"main_spu_id,omitempty"`
	// SpuID
	SpuID uint64 `json:"spu_id,omitempty"`
	// SpuStatus 状态1：可绑定2：审核中3：审核不通过
	SpuStatus enum.SpuStatus `json:"spu_status,omitempty"`
}
