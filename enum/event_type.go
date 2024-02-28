package enum

// EventType 事件枚举
type EventType int

const (
	// EventType_UNKNOWN 未知事件
	EventType_UNKNOWN EventType = 0
	// EventType_ORDER 商品成交
	EventType_ORDER EventType = 1
	// EventType_PAGE_VIEW 中间页曝光
	EventType_PAGE_VIEW EventType = 2
	// EventType_PAGE_CLICK 中间页点击
	EventType_PAGE_CLICK EventType = 3
	// EventType_FORM_SUBMIT 表单提交; 用户在页面内完成表单填写并提交，可以作为优化目标
	EventType_FORM_SUBMIT EventType = 101
	// EventType_FORM_EFFICIENT 有效表单; 提交线索的用户进一步完成了一次有价值的行为，具体行为取决于广告主业务模式，如预约到店等
	EventType_FORM_EFFICIENT EventType = 102
	// EventType_PHONE_CALL 电话拨打; 用户点击给商家拨打电话，可以作为优化目标
	EventType_PHONE_CALL EventType = 103
	// EventType_PHONE_EFFECIENT 电话接通; 商家给用户拨打电话，可接通
	EventType_PHONE_EFFECIENT EventType = 104
	// EventType_WECHAT_COPY 微信复制; 用户完成个人或公众号微信复制行为，引导用户微信加粉，可以作为优化目标
	EventType_WECHAT_COPY EventType = 121
	// EventType_WECHAT_FOLLOW 微信加好友; 用户完成个人或公众号微信加粉
	EventType_WECHAT_FOLLOW EventType = 122
	// EventType_REGISTER 注册
	EventType_REGISTER EventType = 123
	// EventType_WECHAT_MINIAPP 微信小程序访问
	EventType_WECHAT_MINIAPP EventType = 124
	// EventType_IDCERT 身份认证; 用户完成身份认证，比如金融类可能会用身份证号、人脸识别做身份认证
	EventType_IDCERT EventType = 131
	// EventType_PURCHASE 付费; 用户完成商品购买行为，可以作为优化目标
	EventType_PURCHASE EventType = 141
	// EventType_NEGATIVE 负向样本; 不可直接作为广告的优化目标使用，但回传后将有利于优化整体广告效果
	EventType_NEGATIVE EventType = 200
	// EventType_APP_OPEN APP打开; APP端内口令码搜索
	EventType_APP_OPEN EventType = 401
	// EventType_APP_ENTER_SHOP APP进店; 平台客户内店铺详情页浏览量
	EventType_APP_ENTER_SHOP EventType = 402
	// EventType_APP_INTERACT APP互动; 店铺内产生互动行为（用户明确表达转化意愿）
	EventType_APP_INTERACT EventType = 403
	// EventType_APP_PAY APP支付; 线上订单付款成功
	EventType_APP_PAY EventType = 404
)
