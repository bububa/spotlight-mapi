package creativity

// Creativity 创意
type Creativity struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeID 创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// CreativityEnable 创意开启状态
	CreativityEnable int `json:"creativity_enable,omitempty"`
	// CreativityStatus 创意状态
	CreativityStatus int `json:"creativity_status,omitempty"`
	// CreativityCreateTime 创意创建时间
	CreativityCreateTime string `json:"creativity_create_time,omitempty"`
	// MaterialType 笔记类型
	MaterialType int `json:"material_type,omitempty"`
	// ConversionType 组件类型，
	// 0: 无组件，1：商品组件2: 落地页组件，3: 私信组件，4: 直播组件5: poi门店组件6: 外链商品7: 直播间8: 搜索组件9: 小程序组件10：留资组件11：唤端组件（新功能，后续支持）
	ConversionType int `json:"conversion_type,omitempty"`
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// NoteType 笔记类型
	NoteType int `json:"note_type,omitempty"`
	// MaskPerfer 是否需要替换
	MaskPerfer bool `json:"mask_perfer,omitempty"`
	// TitleMaskPerfer 是否需要替换
	TitleMaskPerfer bool `json:"title_mask_perfer,omitempty"`
	// AuditStatus 审核状态
	AuditStatus int `json:"audit_status,omitempty"`
	// AuditComment 审核备注（驳回原因）
	AuditComment string `json:"audit_comment,omitempty"`
	// PageID 落地页id
	PageID string `json:"page_id,omitempty"`
	// ClickURLs 点击URL
	ClickURLs []string `json:"click_urls,omitempty"`
	// ExpoURLs 曝光URL
	ExpoURLs []string `json:"expo_urls,omitempty"`
	// JumpURL 落地页/跳转URL
	JumpURL string `json:"jump_url,omitempty"`
	// BarContent 按钮文案内容
	// 落地页组件类型：立即参与、立即购买、立即领取、立即预约；私信组件类型：立即咨询、立即参与、立即领取、立即预约；
	BarContent string `json:"bar_content,omitempty"`
	// H5ItemInfo 商品信息
	H5ItemInfo interface{} `json:"h5_item_info,omitempty"`
	// H5PageInfo 落地页信息
	H5PageInfo interface{} `json:"h5_page_info,omitempty"`
	// H5MaterialInfo 程序化创意素材
	H5MaterialInfo interface{} `json:"h5_material_info,omitempty"`
	// Image 创意图片
	Image string `json:"image,omitempty"`
	// ItemInvalidReason 商品状态异常原因
	ItemInvalidReason string `json:"item_invalid_reason,omitempty"`
	// ConversionComponentTypes 组件类型
	ConversionComponentTypes []int `json:"conversion_component_types,omitempty"`
	// Comment 备注
	Comment string `json:"comment,omitempty"`
	// Programmatic 是否是程序化创意	0-非程序化创意1-程序化创意
	Programmatic int `json:"programmatic,omitempty"`
	// CreativityExtraInfo 创意extrainfo
	CreativityExtraInfo string `json:"creativity_extra_info,omitempty"`
	// IntoShopParam 店铺小程序转化组件填写参数
	IntoShopParam string `json:"into_shop_param,omitempty"`
	// BootScreenInfo 开屏创意
	BootScreenInfo interface{} `json:"boot_screen_info,omitempty"`
	// PoiID poiId
	PoiID string `json:"poiId,omitempty"`
	// PoiJumpType poi跳转类型
	PoiJumpType string `json:"poi_jump_type,omitempty"`
	// MonitorCompany 监控公司
	MonitorCompany string `json:"monitor_company,omitempty"`
	// MonitorParams 监控参数
	MonitorParams string `json:"monitor_params,omitempty"`
	// AdBizItemID 广告业务ID
	AdBizItemID string `json:"ad_biz_item_id,omitempty"`
	// Title 直播标题	直播推广-标的为直播间时 推广标的为直播时选
	Title string `json:"title,omitempty"`
	// GoodsSellingPoint 商品卖点
	GoodsSellingPoint string `json:"goods_selling_point,omitempty"`
	// DataPostURL 数据回传url
	DataPostURL string `json:"data_post_url,omitempty"`
	// KosMsgType kos私信承接方	0-私信tok1-私信tob
	KosMsgType int `json:"kos_msg_type,omitempty"`
	// QualInfo 资质信息
	QualInfo *QualInfo `json:"qual_info,omitempty"`
	// MiniProgramPath 小程序组件path
	MiniProgramPath string `json:"mini_program_path,omitempty"`
}
