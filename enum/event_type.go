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
	// EventType_WECHAT_FOLLOW 微信加好友
	EventType_WECHAT_FOLLOW EventType = 122
	// EventType_REGISTER 注册
	EventType_REGISTER EventType = 123
	// EventType_WECHAT_MINIAPP 微信小程序访问
	EventType_WECHAT_MINIAPP EventType = 124
)
