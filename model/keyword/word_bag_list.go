package keyword

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// WordBagListRequest 词包推荐 API Request
type WordBagListRequest struct {
	// Name 搜索名称
	Name string `json:"name,omitempty"`
	// Category 状态枚举通用：通用型词包其他：通过接口/api/open/jg/keyword/industry/taxonomy获取，一级类目taxonomy_id
	Category string `json:"category,omitempty"`
	// StartTime 开始时间，示例："2024-02-01"
	StartTIme string `json:"start_time,omitempty"`
	// EndTime 结束时间，示例："2024-02-01"
	EndTime string `json:"end_time,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageNum 页数，默认10
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页码页，默认1
	PageSize int64 `json:"page_size,omitempty"`
}

// Encode implements PostRequest interface
func (r WordBagListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// WordBagListResponse 词包推荐 API Response
type WordBagListResponse struct {
	Data *WordBagListResult `json:"data,omitempty"`
	model.BaseResponse
}

type WordBagListResult struct {
	// Page 页信息
	Page *model.Page `json:"page,omitempty"`
	// List 词包信息
	List []WordBag `json:"word_tag_dto_list,omitempty"`
}
