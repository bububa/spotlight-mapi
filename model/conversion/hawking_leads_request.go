package conversion

import "github.com/bububa/spotlight-mapi/util"

// HawkingLeadsRequest 聚光落地页线索数据回传 API Request
type HawkingLeadsRequest struct {
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaignId,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unitId,omitempty"`
	// CreativityID 创意ID
	CreativityID uint64 `json:"creativityId,omitempty"`
	// LeadsID
	LeadsID string `json:"leadsId,omitempty"`
}

// Encode implement PostRequest interface
func (r HawkingLeadsRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
