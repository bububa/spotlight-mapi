package target

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// KeywordRecommendRequest 获取推荐关键词信息 API Request
type KeywordRecommendRequest struct {
	// Keyword 搜索词
	Keyword string `json:"keyword,omitempty"`
	// NoteIDs 笔记Id列表
	NoteIDs []string `json:"note_ids,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r KeywordRecommendRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// KeywordRecommendResponse 获取推荐关键词信息 API Response
type KeywordRecommendResponse struct {
	model.BaseResponse
	Data []KeywordRecommend `json:"data,omitempty"`
}

type KeywordRecommend struct {
	// Word 搜索词
	Word string `json:"target_word,omitempty"`
	// RecommendReason 搜索词
	RecommendReason string `json:"recommend_reason,omitempty"`
	// CoverNum 覆盖人数
	CoverNum int64 `json:"cover_num,omitempty"`
}
