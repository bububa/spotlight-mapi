package data

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// PoiListRequest 门店信息列表 API Request
type PoiListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r PoiListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PoiListResponse 门店信息列表 API Response
type PoiListResponse struct {
	Data *PoiListResult `json:"data,omitempty"`
	model.BaseResponse
}

type PoiListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageDTO `json:"page_info,omitempty"`
	// PoiInfoList 门店列表
	PoiInfoList []PoiInfo `json:"poi_info_list,omitempty"`
}

// PoiInfo 门店信息
type PoiInfo struct {
	// PoiID 门店id
	PoiID string `json:"poi_id,omitempty"`
	// PoiName 门店名称
	PoiName string `json:"poi_name,omitempty"`
}
