package directlink

import (
	"strconv"

	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 获取直达链接 API Request
type ListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 直达链接id
	ID uint64 `json:"id,omitempty"`
	// PageNum 页码，从1开始
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，最大100
	PageSize int64 `json:"page_size,omitempty"`
	// RemarkName 备注名称，传了会模糊匹配
	RemarkName string `json:"remark_name,omitempty"`
	// TypeStr 传了会按照类型查找，	1-deeplink，2-ulk
	TypeStr string `json:"type_str,omitempty"`
	// StatusStr 传了会按照状态搜索，如果想查多个状态的，用英文逗号分隔，比如：1,2	1-审核中，2-审核通过，3-审核拒绝
	StatusStr string `json:"status_str,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.NewUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("id", strconv.FormatUint(r.ID, 10))
	if r.PageNum > 0 {
		values.Set("page_num", strconv.FormatInt(r.PageNum, 10))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.FormatInt(r.PageSize, 10))
	}
	if r.RemarkName != "" {
		values.Set("remark_name", r.RemarkName)
	}
	if r.TypeStr != "" {
		values.Set("type_str", r.TypeStr)
	}
	if r.StatusStr != "" {
		values.Set("status_str", r.StatusStr)
	}
	ret := values.Encode()
	defer util.ReleaseUrlValues(values)
	return ret
}

// ListResponse 获取直达链接 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// Total 直达链接总数
	Total int64 `json:"total,omitempty"`
	// DirectLinkList 直达链接
	DirectLinkList []DirectLink `json:"direct_link_list,omitempty"`
}

// DirectLink 直达链接
type DirectLink struct {
	// ID 	直达链接的id
	ID uint64 `json:"id,omitempty"`
	// URL url内容
	URL string `json:"url,omitempty"`
	// Type 类型，1-deeplink， 2-ulk
	Type int `json:"type,omitempty"`
	// Status 状态，1-审核中，2-审核通过，3-审核拒绝
	Status int `json:"status,omitempty"`
	// RemarkName 备注名称
	RemarkName string `json:"remark_name,omitempty"`
}
