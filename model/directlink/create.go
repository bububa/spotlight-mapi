package directlink

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CreateRequest 创建直达链接 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DirectLinkList 直达链接信息
	DirectLinkList []DirectLink `json:"direct_link_list,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建直达链接 API Response
type CreateResponse struct {
	model.BaseResponse
	Data []DirectLink `json:"data,omitempty"`
}
