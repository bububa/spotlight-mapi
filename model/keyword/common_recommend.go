package keyword

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CommonRecommendRequest 定向推词-以词推词 API Request
type CommonRecommendRequest struct {
	// RequestType 推词类型
	// note：智能推词-笔记推词industry：行业推词search：以词推词session：上下游推词
	RequestType enum.KeywordRecommendRequestType `json:"request_type,omitempty"`
	// Keyword 以词推词(search)、上下游推词(session)必填
	Keyword string `json:"keyword,omitempty"`
	// TaxonomyID 行业id,通过/api/open/jg/keyword/industry/taxonomy接口获取，行业推词下必传
	TaxonomyID string `json:"taxonomy_id,omitempty"`
	// AttributeList 行业属性列表，通过/api/open/jg/keyword/industry/taxonomy/attribute接口获取
	AttributeList string `json:"attribute_list,omitempty"`
	// AttributeNameList 行业属性名称列表，通过/api/open/jg/keyword/industry/taxonomy/attribute接口获取
	AttributeNameList string `json:"attribute_name_list,omitempty"`
	// RecommendReasonFilter 智能推词-笔记推词：高点击行业推词：高点击以词推词：高点击上下游推词：上游、下游
	RecommendReasionFilter []string `json:"recommend_reason_filter,omitempty"`
	// ItemIDs 	笔记ids，智能推词-笔记推词必传
	ItemIDs []string `json:"item_ids,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionTarget 推广目标
	// 智能推词-笔记推词必传
	// 1: 笔记
	// 2: 商品
	// 7: 自由链接
	// 9: 落地页
	// 18: 直播间
	PromotionTarget int `json:"promotion_target,omitempty"`
	// Rank 排序 1：pv降序 2：pv升序 3：竞争指数降序 4：竞争指数升序
	Rank int `json:"rank,omitempty"`
}

// Encode implements PostRequest interface
func (r CommonRecommendRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CommonRecommendResponse 定向推词-以词推词 API Response
type CommonRecommendResponse struct {
	Data *CommonRecommendResult `json:"data,omitempty"`
	model.BaseResponse
}

type CommonRecommendResult struct {
	// WordList 推荐词列表
	WordList []Word `json:"word_list,omitempty"`
	// BagMonthPV 月pv
	BagMonthPV int64 `json:"bag_month_pv,omitempty"`
	// WordNum 推荐词数量
	WordNum int `json:"word_num,omitempty"`
}
