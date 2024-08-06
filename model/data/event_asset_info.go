package data

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// EventAssetInfoRequest 资产事件获取 API Request
type EventAssetInfoRequest struct {
	// Page
	Page *model.PageDTO `json:"page,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r EventAssetInfoRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// EventAssetInfoResponse 资产事件获取 API Response
type EventAssetInfoResponse struct {
	Data *EventAssetInfoResult `json:"data,omitempty"`
	model.BaseResponse
}

// EventAssetInfoResult
type EventAssetInfoResult struct {
	// Page
	Page *model.PageRespDTO `json:"page,omitempty"`
	// EventAssetDtos 资质信息
	EventAssetDtos []EventAsset `json:"event_asset_dtos,omitempty"`
}

// EventAsset 资质信息
type EventAsset struct {
	// EventAssetName 资产名称
	EventAssetName string `json:"event_asset_name,omitempty"`
	// EventList 资产对应的事件列表
	EventList []Event `json:"event_list,omitempty"`
	// EventAssetID 资产id
	EventAssetID uint64 `json:"event_asset_id,omitempty"`
	// Status 状态
	Status int `json:"status,omitempty"`
}

// Event 资产事件
type Event struct {
	// EventID 资产事件id
	EventID uint64 `json:"event_id,omitempty"`
	// EventType 资产事件
	EventType enum.EventType `json:"event_type,omitempty"`
	// EventStatus 状态
	EventStatus int `json:"event_status,omitempty"`
}
