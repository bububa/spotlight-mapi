package landingpage

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 落地页表单查询 API Request
type ListRequest struct {
	// Keyword 搜索词
	Keyword string `json:"keyword,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 请求的页码，默认为 1
	Page int64 `json:"page,omitempty"`
	// PageSize 每页行数，默认为 10
	PageSize int64 `json:"page_size,omitempty"`
	// Status 固定传2(审核通过)
	Status int `json:"status,omitempty"`
}

// Encode implements PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 落地页表单查询 API Response
type ListResponse struct {
	// Data
	Data *ListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ListResult struct {
	// List 落地页信息
	List []LandingPage `json:"list,omitempty"`
	// Page 当前页
	Page int64 `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int64 `json:"page_size,omitempty"`
	// Total 总数
	Total int64 `json:"total,omitempty"`
}

// LandingPage 落地页
type LandingPage struct {
	// AuditComment 审核结果备注
	AuditComment string `json:"audit_comment,omitempty"`
	// Content 页面内容json str
	Content string `json:"content,omitempty"`
	// DraftPreviewPath 草稿页面内容json str，用于发布后再编辑
	DraftPreviewPath string `json:"draft_preview_path,omitempty"`
	// LeadsApi leads_api
	LeadsApi string `json:"leads_api,omitempty"`
	// ObjectID object_id
	ObjectID string `json:"object_id,omitempty"`
	// OnlineDomain 在线域名
	OnlineDomain string `json:"online_domain,omitempty"`
	// PageName 页面名称
	PageName string `json:"page_name,omitempty"`
	// WebDomain 域名
	WebDomain string `json:"web_domain,omitempty"`
	// WebPath path
	WebPath string `json:"web_path,omitempty"`
	// UnitLandingPageDesc 表单落地页描述
	UnitLandingPageDesc []string `json:"unit_landing_page_desc,omitempty"`
	// ID 页面ID
	ID uint64 `json:"id,omitempty"`
	// Status 落地页状态
	// 未审核/草稿 DRAFT(0), 审核中 IN_AUDIT(1), 审核通过 PASS(2), 审核未通过 REJECT(3), 已过期 EXPIRED(4);
	Status enum.LandingPageStatus `json:"status,omitempty"`
	// DataUpdated 数据是否已更新
	DataUpdated bool `json:"data_updated,omitempty"`
}
