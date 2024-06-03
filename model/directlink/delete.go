package directlink

import "github.com/bububa/spotlight-mapi/util"

// DeleteRequest 删除直达链接 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 直达链接ID
	ID uint64 `json:"id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
