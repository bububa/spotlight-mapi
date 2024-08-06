package unit

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
)

// TargetConfig 定向信息
type TargetConfig struct {
	// TargetGender 定向性别"0"：男"1"：女"all"：全部
	TargetGender enum.TargetGender `json:"target_gender,omitempty"`
	// TargetAge 定向年龄"18-22""23-27""28-32""32-100""all"
	TargetAge string `json:"target_age,omitempty"`
	// TargetDevice 定向设备"ios""android""all"
	TargetDevice enum.TargetDevice `json:"target_device,omitempty"`
	// TargetCity 城市定向通过/get_available_target_info接口获取例如"北京#天津#衡阳"城市名称以#分割,all 代表全部省市
	TargetCity string `json:"target_city,omitempty"`
	// IndustryInterestTarget 行业兴趣
	IndustryInterestTarget *IndustryInterestTarget `json:"industry_interest_target,omitempty"`
	// CrowdTarget 人群包
	CrowdTarget *CrowdTarget `json:"crowd_target,omitempty"`
	// Keywords 关键词定向["口红","化妆"]
	Keywords []string `json:"keywords,omitempty"`
	// Recommend 推荐定向
	Recommend *RecommendTarget `json:"recommend_target,omitempty"`
	// DendelionCrowd 蒲公英人群
	DendelionCrowd *DendelionCrowd `json:"dendelion_crowd,omitempty"`
	// InterestKeywords 关键词兴趣定向，["口红","化妆"]
	InterestKeywords []string `json:"interest_keywords,omitempty"`
	// KeywordTargetPeriod 关键词时间周期，单位天，枚举包括 3，7，15，30	关键词行为定向选择时必填
	KeywordTargetPeriod int `json:"keyword_target_period,omitempty"`
	// KeywordTargetAction 关键词行为类型，1: 搜索，2: 互动，3: 阅读	关键词行为定向选择时必填
	KeywordTargetAction []enum.KeywordTargetAction `json:"keyword_target_action,omitempty"`
	// IntelligentExpension 智能扩量0：关闭1：开启
	IntelligentExpension int `json:"intelligent_expension,omitempty"`
	// HaveReverseBloggerFanTarget 是否选择反选博主粉丝人群
	HaveReverseBloggerFanTarget bool `json:"have_reverse_blogger_fan_target,omitempty"`
	// HaveReverseBloggerPurchasedTarget 是否选择反选博主购买人群
	HaveReverseBloggerPurchasedTarget bool `json:"have_reverse_blogger_purchased_target,omitempty"`
	// HaveBrandRecognitionGroup 是否选择本品牌认知人群
	HaveBrandRecognitionGroup bool `json:"have_brand_recognition_group,omitempty"`
	// HaveCategoryInterestGroup 是否选择行业种草人群
	HaveCategoryInterestGroup bool `json:"have_category_interest_group,omitempty"`
	// LiveStreamingFanTarget 直播人群
	LiveStreamingFanTarget *LiveStreamingFanTarget `json:"live_streaming_fan_target,omitempty"`
	// LiveRoomActiveTarget 直播间活跃人群
	LiveRoomActiveTarget *LiveRoomActiveTarget `json:"live_room_active_target,omitempty"`
	// EcBehaviorTarget 电商行为人群
	EcBehaviorTarget []EcBehaviorTarget `json:"ec_behavior_target,omitempty"`
}

// IndustryInterestTarget 行业兴趣
type IndustryInterestTarget struct {
	// ContentInterest 内容兴趣
	ContentInterest []model.CodeNamePair `json:"content_interest,omitempty"`
	// ShoppingInterest 购物兴趣
	ShoppingInterest []model.CodeNamePair `json:"shopping_interest,omitempty"`
}

// CrowdTarget 人群包列表
type CrowdTarget struct {
	// CrowdPkg 人群包列表
	CrowdPkg []CrowdPackage `json:"crowd_pkg,omitempty"`
}

// CrowdPackage 人群包
type CrowdPackage struct {
	// Value 人群包id
	Value string `json:"value,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// Desc 人群包描述
	Desc string `json:"desc,omitempty"`
	// Type 人群包类型
	// 人群包的类型,特色人群:outside/common,节促人群:timeliness没有type的时候表示是客户自己上传的人群包
	// type: ti, TI潜力转化人群（优化目标是ti人群才可使用，本期不涉及）type: expose, 直播间活跃人群（直播间营销目标才可使用,本期不涉及）
	Type string `json:"type,omitempty"`
	// GroupID 人群包真实ID
	GroupID string `json:"group_id,omitempty"`
	// SyncStatus 同步状态,0:未同步,1:已同步
	// 只用已同步(1)状态的才可用
	SyncStatus int `json:"sync_status,omitempty"`
	// Status 删除可用状态,2:成功,3计算失败:,4:停止更新
	// 只用成功(2)状态的才可用
	Status int `json:"status,omitempty"`
	// DmpPermission
	DmpPermission bool `json:"dmp_permission,omitempty"`
}

// RecommendTarget 推荐定向
type RecommendTarget struct {
	// HighPotential 高潜词包
	HighPotential []string `json:"high_potential,omitempty"`
	// InterestHighPotential 关键词兴趣-高潜词包
	InterestHighPotential []string `json:"interest_high_potential,omitempty"`
}

// DendelionCrowd 蒲公英人群
type DendelionCrowd struct {
	NormalDendelionCrowdList     []NormalDendelionCrowd     `json:"normal_dandelion_crowd_list,omitempty"`
	CustomizedDendelionCrowdList []CustomizedDendelionCrowd `json:"customized_dandelion_crowd_List,omitempty"`
}

type NormalDendelionCrowd struct {
	// ActionType 行为类型:imp,read
	ActionType enum.CrowdActionType `json:"action_type,omitempty"`
	// BrandUserID 品牌用户Id
	BrandUserID string `json:"brand_user_id,omitempty"`
	// TimePeriod 时间周期，30，90
	TimePeroid int `json:"time_period,omitempty"`
}

type CustomizedDendelionCrowd struct {
	// ActionType 行为类型:imp,read
	ActionType enum.CrowdActionType `json:"action_type,omitempty"`
	// CrowdName 蒲公英定制人群包名称
	CrowdName string `json:"crowd_name,omitempty"`
	// NotIDList 笔记ids
	NoteIDList []string `json:"note_id_list,omitempty"`
	// Channels 流量渠道：ad/nature
	Channels []string `json:"channels,omitempty"`
	// TimePeriod 时间周期，30，90
	TimePeroid int `json:"time_period,omitempty"`
	// CrowdID 蒲公英定制人群包Id
	CrowdID uint64 `json:"crowd_id,omitempty"`
}

// LiveStreamingFanTarget 直播人群
type LiveStreamingFanTarget struct{}

// LiveRoomActiveTarget 直播间活跃人群
type LiveRoomActiveTarget struct{}

// EcBehaviorTarget 电商行为人群
type EcBehaviorTarget struct{}
