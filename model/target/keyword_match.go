package target

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// KeywordMatchRequest 获取关键词匹配词库信息 API Request
type KeywordMatchRequest struct {
	// Keywords 关键词列表
	// 不能超过150个
	Keywords []string `json:"keywords,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r KeywordMatchRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// KeywordMatchResponse 获取关键词匹配词库信息 API Response
type KeywordMatchResponse struct {
	Data *KeywordMatchResult `json:"data,omitempty"`
	model.BaseResponse
}

type KeywordMatchResult struct {
	// MatchInfos 匹配信息列表
	MatchInfos []KeywordMatchInfo `json:"match_infos,omitempty"`
	// MatchDistinctCount 匹配个数
	MatchDistinctCount int64 `json:"match_distinct_count,omitempty"`
}

// KeywordMatchInfo 匹配信息
type KeywordMatchInfo struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// InThesaurus 是否匹配词库
	InThesaurus bool `json:"in_thesaurus,omitempty"`
}
