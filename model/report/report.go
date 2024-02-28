package report

import "github.com/bububa/spotlight-mapi/model"

// DataReportDTO 数据指标
type DataReportDTO struct {
	// Fee 推广消费金额
	Fee model.Float64 `json:"fee,omitempty"`
	// Impression 推广展现量
	Impression model.Int64 `json:"impression,omitempty"`
	// Click 推广点击量；视频流中「观看视频5s」记做一次点击
	Click model.Int64 `json:"click,omitempty"`
	// Ctr 点击率; 推广点击量/推广展现量*100%
	Ctr model.Float64 `json:"ctr,omitempty"`
	// Acp 平均点击成本; 推广消费/推广点击量
	Acp model.Float64 `json:"acp,omitempty"`
	// Cpm 平均千次曝光; 推广消费/推广展现量*1000
	Cpm model.Float64 `json:"cpm,omitempty"`
	// Like 点赞; 用户点击推广后1小时内，对笔记的点赞量
	Like model.Int64 `json:"like,omitempty"`
	// Comment 评论; 用户点击推广后1小时内，对笔记的评论量
	Comment model.Int64 `json:"comment,omitempty"`
	// Collect 收藏;用户点击推广后1小时内，对笔记的收藏量
	Collect model.Int64 `json:"collect,omitempty"`
	// Follow 关注;用户点击推广后，对笔记作者的关注量
	Follow model.Int64 `json:"follow,omitempty"`
	// Share 分享;用户点击推广后1小时内，对笔记进行站内或站外分享量, 支持2021.07.07之后的数据
	Share model.Int64 `json:"share,omitempty"`
	// Interaction 互动量; 点赞、收藏、关注、评论、分享5个互动行为指标之和
	Interaction model.Int64 `json:"interaction,omitempty"`
	// Cpi 平均互动成本; 消费/互动量
	Cpi model.Float64 `json:"cpi,omitempty"`
	// ActionButtonClick 行动按钮点击量; 不同类型创意包含的组件点击量，包含预告组件点击
	ActionButtonClick model.Int64 `json:"action_button_click,omitempty"`
	// ActionButtonCtr 行动按钮点击率; 行动按钮点击量/点击量
	ActionButtonCtr model.Float64 `json:"action_button_ctr,omitempty"`
	// Screenshot 截图; 用户点击推广后1小时内，对笔记进行截图的次数
	Screenshot model.Int64 `json:"screenshot,omitempty"`
	// PicSave 保存图片; 用户点击推广后1小时内，保存笔记中图片的次数
	PicSave model.Int64 `json:"pic_save,omitempty"`
	// ReservePv 预告组件点击; 用户点击推广1小时内，在账户内的预约人次。不去重
	ReservePv model.Int64 `json:"reserve_pv,omitempty"`
	// ClkLiveEntryPv 直播间观看次数; 用户点击推广后24小时内，进入直播间的总次数
	ClkLiveEntryPv model.Int64 `json:"clk_live_entry_pv,omitempty"`
	// ClkLiveEntryPvCost 直播间观看成本; 推广消费/直播间观看次数
	ClkLiveEntryPvCost model.Float64 `json:"clk_live_entry_pv_cost,omitempty"`
	// ClkLiveAvgViewTime 直播间人均停留时长; 用户点击推广后24小时内，进入直播间后的平均观看时长（单位：分钟）
	ClkLiveAvgViewTime model.Float64 `json:"clk_live_avg_view_time,omitempty"`
	// ClkLiveAllFollow 直播间新增粉丝量; 用户点击推广后24小时内，进入直播间后在直播间内关注主播的总次数
	ClkLiveAllFollow model.Int64 `json:"clk_live_all_follow,omitempty"`
	// ClkLive5sEntryPv 直播间有效观看次数; 用户点击推广后24小时内，进入直播间后停留时长大于等于5秒的总次数
	ClkLive5sEntryPv model.Int64 `json:"clk_live_5s_entry_pv,omitempty"`
	// ClkLive5sEntryUvCost 直播间有效观看成本; 推广消费/直播间有效观看人次
	ClkLive5sEntryUvCost model.Float64 `json:"clk_live_5s_entry_uv_cost,omitempty"`
	// ClkLiveComment 直播间评论次数; 用户点击推广后24小时内，进入直播间后在直播间内评论的总次数
	ClkLiveComment model.Int64 `json:"clk_live_comment,omitempty"`
	// SearchCmtClick 搜索组件点击量; 点击推广搜索组件的次数
	SearchCmtClick model.Int64 `json:"search_cmt_click,omitempty"`
	// SearchCmtClickCvr 搜索组件点击转化率; 搜索组件点击量/点击量
	SearchCmtClickCvr model.Float64 `json:"search_cmt_click_cvr,omitempty"`
	// SearchCmtAfterRead 搜后阅读量; 点击组件后的搜索场景阅读量
	SearchCmtAfterRead model.Int64 `json:"search_cmt_after_read,omitempty"`
	// SearchCmtAfterReadAvg 平均搜后阅读笔记篇数
	SearchCmtAfterReadAvg model.Float64 `json:"search_cmt_after_read_avg,omitempty"`
	// GoodsVisit 进店访问量; 用户点击推广后的7日内，对客户所绑定小红书店铺首页或商品详情页的访问总次数
	GoodsVisit model.Int64 `json:"goods_visit,omitempty"`
	// GoodsVisitPrice 进店访问成本; 推广消费/进店访问量
	GoodsVisitPrice model.Float64 `json:"goods_visit_price,omitempty"`
	// SellerVisit 商品访客量; 用户点击推广后的7日内，对客户所绑定小红书店铺首页或商品详情页的访问总人数
	SellerVisit model.Int64 `json:"seller_visit,omitempty"`
	// SellerVisitPrice 商品访客成本; 推广消费/商品访客量
	SellerVisitPrice model.Float64 `json:"seller_visit_price,omitempty"`
	// ShoppingCartAdd 商品加购量; 用户点击推广后的7日内，对客户所绑定小红书店铺全部商品加入购物车的总次数
	ShoppingCartAdd model.Int64 `json:"shopping_cart_add,omitempty"`
	// AddCartPrice 商品加购成本; 推广消费/商品加购量
	AddCartPrice model.Float64 `json:"add_cart_price,omitempty"`
	// PresaleOrderNum7d 7日预售订单量; 用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部预售商品，完成预售定金支付的总订单数
	PresaleOrderNum7d model.Int64 `json:"presale_order_num_7d,omitempty"`
	// PresaleOrderGmv7d 7日预售订单金额; 用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部预售商品，完成预售定金支付的订单所对应的商品总金额（含定金与应支付尾款）
	PresaleOrderGmv7d model.Float64 `json:"presale_order_gmv_7d,omitempty"`
	// GoodsOrder 7日下单订单量; 用户点击广告后，7天内在广告主绑定的店铺，对全部商品的创建订单数之和
	GoodsOrder model.Int64 `json:"goods_order,omitempty"`
	// GoodsOrderPrice 7日下单订单成本; 推广消费/7日下单订单量
	GoodsOrderPrice model.Float64 `json:"goods_order_price,omitempty"`
	// Rgmv 7日下单金额; 用户点击广告后，7天内在广告主绑定的店铺，对全部商品的支付金额之和（含退款金额）
	Rgmv model.Float64 `json:"rgmv,omitempty"`
	// Roi 7日下单ROI; rgmv/消费
	Roi model.Float64 `json:"roi,omitempty"`
	// SuccessGoodsOrder 7日支付订单量; 用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部商品的支付总订单数量
	SuccessGoodsOrder model.Int64 `json:"success_goods_order,omitempty"`
	// ClickOrderCvr 7日支付转化率; 7日支付订单量/推广点击数*100%
	ClickOrderCvr model.Float64 `json:"click_order_cvr,omitempty"`
	// PurchaseOrderPrice7d 7日支付订单成本; 推广消费/7日支付订单量
	PurchaseOrderPrice7d model.Float64 `json:"purchase_order_price_7d,omitempty"`
	// PurchaseOrderGmv7d 7日支付金额; 用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部商品的支付总订单金额
	PurchaseOrderGmv7d model.Float64 `json:"purchase_order_gmv_7d,omitempty"`
	// PurchaseOrderRoi7d 7日支付ROI; 7日支付金额/推广消费
	PurchaseOrderRoi7d model.Float64 `json:"purchase_order_roi_7d,omitempty"`
	// ClkLiveRoomOrderNum 直播间支付订单量; 用户点击推广后的24小时内，在直播间内全部商品的支付总订单量
	ClkLiveRoomOrderNum model.Int64 `json:"clk_live_room_order_num,omitempty"`
	// LiveAverateOrderCost 直播间支付订单成本; 推广消费/直播间支付订单量
	LiveAverateOrderCost model.Float64 `json:"live_averate_order_cost,omitempty"`
	// ClkLiveRoomRgmv 直播间支付金额; 用户点击推广后的24小时内，在直播间内全部商品的支付总订单金额
	ClkLiveRoomRgmv model.Float64 `json:"clk_live_room_rgmv,omitempty"`
	// ClkLiveRoomRoi 直播间支付ROI; 直播间支付金额/推广消费
	ClkLiveRoomRoi model.Float64 `json:"clk_live_room_roi,omitempty"`
	// Leads 表单提交; 用户点击广告后成功提交表单的次数
	Leads model.Int64 `json:"leads,omitempty"`
	// LeadsCpl 表单成本; 消费/表单数
	LeadsCpl model.Float64 `json:"leads_cpl,omitempty"`
	// LandingPageVisit 落地页访问量（行为时间）; 推广落地页访问量，时间记录在行为时间上
	LandingPageVisit model.Int64 `json:"landing_page_visit,omitempty"`
	// LeadsButtonImpression 表单按钮曝光量（行为时间）; 推广表单按钮曝光量，时间记录在行为时间上
	LeadsButtonImpression model.Int64 `json:"leads_button_impression,omitempty"`
	// ValidLeads 有效表单; 用户提交表单后7天内，有过进一步的沟通行为（需要您回传此事件）；转化记录在计费时间上
	ValidLeads model.Int64 `json:"valid_leads,omitempty"`
	// ValidLeadsCpl 有效表单成本; 消费/有效表单
	ValidLeadsCpl model.Float64 `json:"valid_leads_cpl,omitempty"`
	// LeadsCvr 表单转化率; 表单数/点击数
	LeadsCvr model.Float64 `json:"leads_cvr,omitempty"`
	// PhoneCallCnt 电话拨打; 用户点击推广后拨打电话的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	PhoneCallCnt model.Int64 `json:"phone_call_cnt,omitempty"`
	// PhoneCallSuccCnt 电话接通; 用户拨打电话后7天内，能拨通电话的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	PhoneCallSuccCnt model.Int64 `json:"phone_call_succ_cnt,omitempty"`
	// WechatCopyCnt 微信复制; 用户点击推广后，用户点击推广后成功复制微信号的次数，（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	WechatCopyCnt model.Int64 `json:"wechat_copy_cnt,omitempty"`
	// WechatCopySuccCnt 微信加为好友; 用户复制微信后7天内，成功复制微信且成功加为好友（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	WechatCopySuccCnt model.Int64 `json:"wechat_copy_succ_cnt,omitempty"`
	// IdentityCertCnt 身份认证; 用户点击推广后，用户点击推广后成功完成身份认证的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	IdentityCertCnt model.Int64 `json:"identity_cert_cnt,omitempty"`
	// CommodityBuyCnt 商品购买; 用户点击推广后，用户点击推广后成功完成付费的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	CommodityBuyCnt model.Int64 `json:"commodity_buy_cnt,omitempty"`
	// MessageUser 私信咨询人数; 用户点击推广后24小时内，通过推广产生私信对话的用户数（包括点击私信组件、专业号主页咨询等）
	MessageUser model.Int64 `json:"message_user,omitempty"`
	// Message 私信咨询条数; 用户点击推广后24小时内，通过推广产生私信对话的条数
	Message model.Int64 `json:"message,omitempty"`
	// MessageConsult 私信咨询数; 用户点击推广后24小时内，至少产生过一次咨询记一个私信咨询数（包括点击私信组件、专业号主页咨询等）
	MessageConsult model.Int64 `json:"message_consult,omitempty"`
	// InitiativeMessage 私信开口数; 用户点击推广24小时内，至少产生过一次开口行为记一个私信开口数
	InitiativeMessage model.Int64 `json:"initiative_message,omitempty"`
	// MessageConsultCpl 私信咨询成本; 消耗/私信咨询数
	MessageConsultCpl model.Float64 `json:"message_consult_cpl,omitempty"`
	// InitiativeMessageCpl 私信开口成本; 消耗/私信开口数
	InitiativeMessageCpl model.Float64 `json:"initiative_message_cpl,omitempty"`
	// MsgLeadsNum 私信留资数; 用户点击推广后7日内，成功留资转化的次数（需要在私信回复工具上进行客资标注）
	MsgLeadsNum model.Int64 `json:"msg_leads_num,omitempty"`
	// MsgLeadsCost 私信留资成本; 消耗/私信留资数
	MsgLeadsCost model.Float64 `json:"msg_leads_cost,omitempty"`
	// ExternalGoodsVisit24h 行业商品点击量; 用户点击广告组件进行跳转的行为次数
	ExternalGoodsVisit24h model.Int64 `json:"external_goods_visit_24h,omitempty"`
	// ExternalGoodsVisitRate24h 行业商品点击转化率; 行业商品点击量 / 点击量
	ExternalGoodsVisitRate24h model.Float64 `json:"external_goods_visit_rate_24h,omitempty"`
	// ExternalGoodsOrder7 行业商品成交订单量（7日); 用户点击广告后，7天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，7天内对全部商品的创建且成功支付订单数之和。
	ExternalGoodsOrder7 model.Int64 `json:"external_goods_order_7,omitempty"`
	// ExternalRgmv7 行业商品GMV（7日）; 用户点击广告后，7天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，7天内对全部商品的成交金额之和。
	ExternalRgmv7 model.Float64 `json:"external_rgmv_7,omitempty"`
	// ExternalGoodsOrderPrice7 行业商品成交订单成本（7日）; 消费 / 行业商品成交订单量(7日)
	ExternalGoodsOrderPrice7 model.Float64 `json:"external_goods_order_price_7,omitempty"`
	// ExternalGoodsOrderRate7 行业商品成交订单转化率（7日）; 行业商品成交订单量(7日) / 点击量
	ExternalGoodsOrderRate7 model.Float64 `json:"external_goods_order_rate_7,omitempty"`
	// ExternalRoi7 行业商品ROI（7日）; 行业商品GMV(7日) / 消费
	ExternalRoi7 model.Float64 `json:"external_roi_7,omitempty"`
	// ExternalGoodsOrder15 行业商品成交订单量（15日）; 用户点击广告后，15天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，15天内对全部商品的创建且成功支付订单数之和。
	ExternalGoodsOrder15 model.Int64 `json:"external_goods_order_15,omitempty"`
	// ExternalRgmv15 行业商品GMV（15日）; 用户点击广告后，15天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，15天内对全部商品的成交金额之和。
	ExternalRgmv15 model.Float64 `json:"external_rgmv_15,omitempty"`
	// ExternalGoodsOrderPrice15 行业商品成交订单成本（15日）; 消费 / 行业商品成交订单量(15日)
	ExternalGoodsOrderPrice15 model.Float64 `json:"external_goods_order_price_15,omitempty"`
	// ExternalGoodsOrderRate15 行业商品成交订单转化率（15日）; 行业商品成交订单量(15日) / 点击量
	ExternalGoodsOrderRate15 model.Float64 `json:"external_goods_order_rate_15,omitempty"`
	// ExternalRoi15 行业商品ROI（15日）; 行业商品GMV(15日) / 消费
	ExternalRoi15 model.Float64 `json:"external_roi_15,omitempty"`
	// ExternalGoodsOrder30 行业商品成交订单量（30日）; 用户点击广告后，30天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，30天内对全部商品的创建且成功支付订单数之和。
	ExternalGoodsOrder30 model.Int64 `json:"external_goods_order_30,omitempty"`
	// ExternalRgmv30 行业商品GMV（30日）; 用户点击广告后，30天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，30天内对全部商品的成交金额之和。
	ExternalRgmv30 model.Float64 `json:"external_rgmv_30,omitempty"`
	// ExternalGoodsOrderPrice30 行业商品成交订单成本（30日）; 消费 / 行业商品成交订单量(30日)
	ExternalGoodsOrderPrice30 model.Float64 `json:"external_goods_order_price_30,omitempty"`
	// ExternalGoodsOrderRate30 行业商品成交订单转化率（30日）; 行业商品成交订单量(30日) / 点击量
	ExternalGoodsOrderRate30 model.Float64 `json:"external_goods_order_rate_30,omitempty"`
	// ExternalRoi30 行业商品ROI（30日）; 行业商品GMV(30日) / 消费
	ExternalRoi30 model.Float64 `json:"external_roi_30,omitempty"`
	// ExternalLeads 外链转化数; 外链回传的转化数据，例如：表单数、下载数、付费数
	ExternalLeads model.Int64 `json:"external_leads,omitempty"`
	// ExternalLeadsCpl 平均外链转化成本; 消费/转化数
	ExternalLeadsCpl model.Float64 `json:"external_leads_cpl,omitempty"`
	// WordAvgLocation 平均位次
	WordAvgLocation model.Int `json:"word_avg_location,omitempty"`
	// WordImpressionRankFirst 首位曝光排名
	WordImpressionRankFirst model.Int `json:"word_impression_rank_first,omitempty"`
	// WorkImpressionRateFirst 首位曝光占比
	WordImpressionRateFirst model.Float64 `json:"word_impression_rate_first,omitempty"`
	// WordImpressionRankThird 前三位曝光排名
	WordImpressionRankThird model.Int `json:"word_impression_rank_third,omitempty"`
	// WorkImpressionRateThird 前三位曝光占比
	WordImpressionRateThird model.Float64 `json:"word_impression_rate_third,omitempty"`
	// WordClickRankFirst 首位点击排名
	WordClickRankFirst model.Int `json:"word_click_rank_first,omitempty"`
	// WorkClickRateFirst 首位点击占比
	WordClickRateFirst model.Float64 `json:"word_click_rate_first,omitempty"`
	// WordClickRankThird 前三位点击排名
	WordClickRankThird model.Int `json:"word_click_rank_third,omitempty"`
	// WorkClickRateThird 前三位点击占比
	WordClickRateThird model.Float64 `json:"word_click_rate_third,omitempty"`
}
