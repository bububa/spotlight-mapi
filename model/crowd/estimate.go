package crowd

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/unit"
	"github.com/bububa/spotlight-mapi/util"
)

// EstimateRequest 人群预估 API Request
type EstimateRequest struct {
	// TargetConfig 定向配置
	TargetConfig *unit.TargetConfig `json:"target_config,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingTarget 营销目标3：商品销量4：笔记种草8：直播推广9：客资收集10：抢占赛道13：种草直达14：直播预热
	MarketingTarget int `json:"marketing_target,omitempty"`
	// Placement 广告类型1：信息流2：搜索推广4：全站智投7：视频内流
	PlaceMent int `json:"placement,omitempty"`
	// OptimizeTarget 推广目标：0：点击量1：互动量3：表单提交量4：商品成单量5：私信咨询量6：直播间观看量11：商品访客量12：落地页访问量13：私信开口量14：有效观看量18：站外转化量20：TI人群规模21：行业商品成单23：直播预热量24：直播间成交25：直播间支付ROI
	OptimizeTarget int `json:"optimize_target,omitempty"`
	// TargetType 定向类型，1-通投,2-智能定向, 3-高级定向
	TargetType int `json:"target_type,omitempty"`
}

// Encode implements PostRequest interface
func (r EstimateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// EstimateResponse 人群预估 API Response
type EstimateResponse struct {
	Data *EstimateResult `json:"data,omitempty"`
	model.BaseResponse
}

type EstimateResult struct {
	// CrowdNum 受众人数,精确到百万
	CrowdNum string `json:"crowd_num,omitempty"`
	// RawCrowdNum 受众人数，纯数值
	RawCrowdNum int64 `json:"raw_crowd_num,omitempty"`
	// CrowdScope 受众范围1-偏窄,2-合适,3-偏广
	CrowdScope int `json:"crowd_scope,omitempty"`
}
