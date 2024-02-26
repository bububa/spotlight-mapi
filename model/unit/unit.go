package unit

import (
	"github.com/bububa/spotlight-mapi/enum"
)

// Unit 广告单元
type Unit struct {
	// ID 单元id
	ID uint64 `json:"id,omitempty"`
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// Name 单元名称
	Name string `json:"name,omitempty"`
	// Enable 投放状态：0：暂停1：投放中
	Enable int `json:"enable,omitempty"`
	// EventBid 出价,单位(分)
	EventBid int64 `json:"event_bid,omitempty"`
	// TargetType 定向类型1-通投,2-智能定向,3-高级定向
	TargetType enum.UnitTargetType `json:"target_type,omitempty"`
	// ItemIDs 商品ID
	ItemIDs []string `json:"item_ids,omitempty"`
	// NoteIDs 笔记ID
	NoteIDs []string `json:"note_ids,omitempty"`
	// LiveUserID 直播用户ID
	LiveUserID string `json:"live_user_id,omitempty"`
	// PageID 落地页ID
	PageID string `json:"page_id,omitempty"`
	// LandingPageURL 	落地页Url
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// UnitExternalPageURL 外链Url
	UnitExternalPageURL string `json:"unit_external_page_url,omitempty"`
	// LandingPageType 落地页链接类型:1-表单,2-外跳链接，0-默认值，无实际意义
	LandingPageType enum.LandingPageType `json:"landing_page_type,omitempty"`
	// TargetPosition 抢占资源1-首位,3-第三位，0-不限位置
	TargetPosition int `json:"target_position,omitempty"`
	// TargetGoal 抢占目标1-点击抢占市场份额，0-默认值，无实际意义
	TargetGoal int `json:"target_goal,omitempty"`
	// WordTagName 词包名称
	WordTagName string `json:"word_tag_name,omitempty"`
	// ProportionGoal 占比目标
	ProportionGoal float64 `json:"proportion_goal,omitempty"`
	// BusinessTreeName 推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉
	BusinessTreeName string `json:"business_tree_name,omitempty"`
	// UnitLandingPageSpec 落地页表单描述
	UnitLandingPageSpec []string `json:"unit_landing_page_spec,omitempty"`
	// KeywordTargetPeriod 关键词定向行为周期，单位天，枚举包括 3，7，15，30
	KeywordTargetPeriod int `json:"keyword_target_period,omitempty"`
	// KeywordTargetAction 关键词定向行为1:搜索,2:互动,3:阅读
	KeywordTargetAction []enum.KeywordTargetAction `json:"keyword_target_action,omitempty"`
	// SubstitutedUserID 代投账号b的userId
	SubstitutedUserID string `json:"substituted_user_id,omitempty"`
	// CreateTime 单元创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 单元修改时间
	UpdateTime string `json:"update_time,omitempty"`
	// ItemNoteInfo 单元标的信息
	ItemNoteInfo []ItemNoteConfig `json:"item_note_info,omitempty"`
	// SpuNoteInfo spu&笔记标的信息
	SpuNoteInfo []SpuNoteConfig `json:"spu_note_info,omitempty"`
	// TargetConfig 定向信息
	TargetConfig *TargetConfig `json:"target_config,omitempty"`
	// KeywordGenType 单元选词方式： -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能
	KeywordGenType enum.KeywordGenType `json:"keyword_gen_type,omitempty"`
	// KeywordWithBid 选词出价
	KeywordWithBid []KeywordWithBid `json:"keyword_with_bid,omitempty"`
}

// ItemNoteConfig 单元标的信息
type ItemNoteConfig struct {
	// ItemID 商品id
	ItemID string `json:"item_id,omitempty"`
	// NoteIDs 绑定的笔记id
	NoteIDs []string `json:"note_ids,omitempty"`
}

// SpuNoteConfig spu&笔记标的信息
type SpuNoteConfig struct {
	// SpuID spuid
	SpuID string `json:"spu_id,omitempty"`
	// NoteIDs 绑定的笔记id
	NoteIDs []string `json:"note_ids,omitempty"`
}

// KeywordWithBid 选词出价
type KeywordWithBid struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Bid 关键词出价
	Bid int64 `json:"bid,omitempty"`
	// KeywordSource 关键词来源
	// 2：旧通道来源
	// 11：笔记id通道
	// 14：笔记行业通道
	// 15：笔记类目通道
	// 16：笔记图模型通道
	// 18：笔记id-v2通道
	// 20：以词推词核心term通道
	// 21：以词推词queryRewrite通道
	// 25：以词推词ES通道
	// 31：商品id通道
	// 32：商品spu类目通道
	// 35：商品类目通道
	// 36：商品行业通道
	// 50：表单默认通道
	// 51：表单行业通道
	// 57：表单pageId通道
	// 102：在线向量召回通道
	// 103：高相关性通道
	// 110：智投话题通道
	// 113：基础词包通道
	// 120：智投默认通道
	// 131：智投swing通道
	// 211：抢赛道通道
	KeywordSource int `json:"keyword_source,omitempty"`
	// PhraseMatchType 匹配方式
	// 0:精确匹配
	// 1:短语匹配
	// 2:智能匹配
	PhraseMatchType enum.PhraseMatchType `json:"phrase_match_type,omitempty"`
	// FeedBid 搜索追投出价
	FeedBid int64 `json:"feed_bid,omitempty"`
}
