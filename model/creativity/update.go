package creativity

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateRequest 编辑创意 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativityID 创意ID
	CreativityID uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称
	CreativityName string `json:"creativity_name,omitempty"`
	// ClickURLs 点击链接, 唤端场景不传
	ClickURLs []string `json:"click_urls,omitempty"`
	// ExpoURLs 曝光链接，唤端场景不传
	ExpoURLs []string `json:"expo_urls,omitempty"`
	// MaskPerfer 是否开启封面优选。默认0，开启传1
	MaskPerfer *bool `json:"mask_perfer,omitempty"`
	// TitleMaskPerfer 是否开启标题优选，默认0，开启传1
	TitleMaskPerfer *bool `json:"title_mask_perfer,omitempty"`
	// JumpURL 落地页/外链url
	JumpURL string `json:"jump_url,omitempty"`
	// BarContent 按钮文案内容，包括：
	// 落地页组件类型：立即参与、立即购买、立即领取、立即预约；
	// 私信组件类型：立即咨询、立即参与、立即领取、立即预约；
	// 唤端组件类型：了解详情，购买同款，领取补贴，优惠下单。
	// 如果是搜索组件则是搜索词，1~20个字符
	BarContent string `json:"bar_content,omitempty"`
	// ItemID 商品id
	ItemID string `json:"item_id,omitempty"`
	// H5Infos 前链h5信息
	H5Infos interface{} `json:"h5_infos,omitempty"`
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
	// H5MaterialInfo 程序化创意素材
	H5MaterialInfo interface{} `json:"h5_material_info,omitempty"`
	// PoiID poiId
	PoiID string `json:"poi_id,omitempty"`
	// PoiJumpType poi组件跳转类型
	PoiJumpType string `json:"poi_jump_type,omitempty"`
	// MonitorCompany
	MonitorCompany string `json:"monitor_company,omitempty"`
	// MonitorParams 监测参数配置
	MonitorParams string `json:"monitor_params,omitempty"`
	// AdBizItemID 广告侧绑定商品id
	AdBizItemID string `json:"ad_biz_item_id,omitempty"`
	// AppCompIcon 唤端下，商品主图的地址
	AppCompIcon string `json:"app_comp_icon,omitempty"`
	// FallBackJumpURL 唤端下，兜底链接
	FallBackJumpURL string `json:"fall_back_jump_url,omitempty"`
}

// Encode implements PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 编辑创意 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *struct {
		// CreativityID 创意ID
		CreativityID uint64 `json:"creativity_id,omitempty"`
	} `json:"data,omitempty"`
}
