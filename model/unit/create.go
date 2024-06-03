package unit

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CreateRequest 创建单元
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// EventBid 出价/目标成本，单位分，自动控制不需要传
	EventBid int64 `json:"event_bid,omitempty"`
	// NoteIDs 标的-笔记id(帮助定向推荐，创意里的笔记会单独填写)非种草人群规模、深度种草人群规模下必填
	NoteIDs []string `json:"note_ids,omitempty"`
	// PromotionTarget 推广标的（笔记1）
	PromotionTarget int `json:"promotion_target,omitempty"`
	// TargetType 定向类型，1-通投,2-智能定向, 3-高级定向
	TargetType enum.UnitTargetType `json:"target_type,omitempty"`
	// TargetConfig 定向配置
	TargetConfig *TargetConfig `json:"target_config,omitempty"`
	// KeywordTargetPeriod 关键词时间周期，单位天，枚举包括 3，7，15，30
	KeywordTargetPeriod int `json:"keyword_target_period,omitempty"`
	// KeywordTargetAction 关键词目标行为
	// 1: 搜索，2: 互动，3: 阅读
	KeywordTargetAction []enum.KeywordTargetAction `json:"keyword_target_action,omitempty"`
	// BusinessTreeName 推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉具体参看/api/open/jg/keyword/industry/taxonomy接口返回词。
	BusinessTreeName string `json:"business_tree_name,omitempty"`
	// SpuNoteInfo spu&笔记标的信息推广目标是种草人群规模、深度种草人群规模下必传。spu只能一个，至少绑定5篇笔记，最多50(api限制)只能选择合作笔记与我的笔记，笔记违规不能选，spu审核不通过也可以选，选择的笔记一定是绑定在该spu上。
	SpuNoteInfo []SpuNoteConfig `json:"spu_note_info,omitempty"`
	// KeywordWithBid 关键词，通过(关键词词包)/api/open/jg/keyword/word/bag/list (智能推词、行业推词、以词推词)/api/open/jg/keyword/common/recommend 接口获取搜索推广-标准投放下必传
	KeywordWithBid []KeywordWithBid `json:"keyword_with_bid,omitempty"`
	// SubstitutedUserID 代投账号b的userId，代投笔记与其他笔记互斥,，需要校验是否有代投账号权限
	SubstitutedUserID string `json:"substituted_user_id,omitempty"`
	// KeywordGenType 单元选词方式：
	// -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能
	// 关键词定向且在白名单中才支持
	KeywordGenType int `json:"keyword_gen_type,omitempty"`
	// PageID 落地页ID，聚光落地页下必填
	PageID string `json:"page_id,omitempty"`
	// LandingPageURL 落地页Url，自研落地页下必填
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// UnitExternalPageURL 外链Url，标的是外链落地页时必填
	UnitExternalPageURL string `json:"unit_external_page_url,omitempty"`
	// UnitLandingPageDesc 落地页表单描述
	UnitLandingPageDesc string `json:"unit_landing_page_desc,omitempty"`
	// TargetTemplateID 定向包id
	TargetTemplateID string `json:"target_template_id,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建单元 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// UnitID 单元ID
		UnitID uint64 `json:"unit_id,omitempty"`
	} `json:"data,omitempty"`
}
