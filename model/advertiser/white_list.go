package advertiser

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// WhiteListRequest 账户白名单 API Request
type WhiteListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implments PostRequest interface
func (r WhiteListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// WhiteListResponse 账户白名单 API Response
type WhiteListResponse struct {
	model.BaseResponse
	Data *WhiteList `json:"data,omitempty"`
}

// WhiteList 账户白名单
type WhiteList struct {
	// InNoteForceBindSpuWhiteList 是否在笔记强绑spu白名单
	InNoteForceBindSpuWhiteList bool `json:"in_note_force_bind_spu_white_list,omitempty"`
}
