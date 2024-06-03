package creativity

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CreateRequest 创建笔记创意 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativityName 创意名称
	CreativityName string `json:"creativity_name,omitempty"`
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// ClickURLs 点击链接, 唤端场景不传
	ClickURLs []string `json:"click_urls,omitempty"`
	// ExpoURLs 曝光链接，唤端场景不传
	ExpoURLs []string `json:"expo_urls,omitempty"`
	// MaskPerfer 是否开启封面优选。默认0，开启传1
	MaskPerfer int `json:"mask_perfer,omitempty"`
	// TitleMaskPerfer 是否开启标题优选，默认0，开启传1
	TitleMaskPerfer int `json:"title_mask_perfer,omitempty"`
	// ConversionType 组件类型，0: 无组件，2: 落地页组件，3: 私信组件，5: poi门店组件，8: 搜索组件，9: 小程序组件，10: 留资组件，11：唤端组件
	ConversionType int `json:"conversion_type,omitempty"`
	// JumpURL 落地页/外链url
	JumpURL string `json:"jump_url,omitempty"`
	// LandingPageType 落地页链接类型，1：站内落地页，2：外链
	LandingPageType int `json:"landing_page_type,omitempty"`
	// BarContent 按钮文案内容，包括：
	// 落地页组件类型：立即参与、立即购买、立即领取、立即预约；
	// 私信组件类型：立即咨询、立即参与、立即领取、立即预约；
	// 唤端组件类型：了解详情，购买同款，领取补贴，优惠下单。
	// 如果是搜索组件则是搜索词，1~20个字符
	BarContent string `json:"bar_content,omitempty"`
	// ConversionComponentTypes 组件位置，0：默认位置，1：置顶评论
	// 两者可以同时选择，默认位置具体为：
	// 落地页组件的是图片底表，
	// 私信组件的是互动胶囊，
	// 搜索组件的是图片或视频；
	// 底部唤端组件是互动胶囊
	ConversionComponentTypes []int `json:"conversion_component_types,omitempty"`
	// Comment 置顶评论文案
	// 当conversion_component_types包含置顶评论时，评论文案
	Comment string `json:"comment,omitempty"`
	// AppCompIcon 唤端下，商品主图的地址，必传
	AppCompIcon string `json:"app_comp_icon,omitempty"`
	// FallBackJumpURL 唤端下，兜底链接，必选
	FallBackJumpURL string `json:"fall_back_jump_url,omitempty"`
	// QualInfo 创意对应的资质信息,信息见：/api/open/jg/data/qual/info
	// 传了能加快审核速度
	QualInfo *QualInfo `json:"qual_info,omitempty"`
}

// QualInfo 创意对应的资质信息
type QualInfo struct {
	// ApplyID 申请id，对应行业资质
	ApplyID string `json:"apply_id,omitempty"`
	// ProductQualIDList 产品资质id
	ProductQualIDList []uint64 `json:"product_qual_id_list,omitempty"`
	// BrandQualIDList 品牌资质id
	BrandQualIDList []uint64 `json:"brand_qual_id_list,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建笔记创意 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// CreativityID 创意ID
		CreativityID uint64 `json:"creativity_id,omitempty"`
	} `json:"data,omitempty"`
}
