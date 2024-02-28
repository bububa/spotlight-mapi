package conversion

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/util"
)

// Request 转化回传请求
type Request struct {
	// AppID
	AppID string `json:"app_id,omitempty"`
	// AccessToken
	AccessToken string `json:"access_token,omitempty"`
	// EventType 事件类型
	EventType enum.EventType `json:"event_type,omitempty"`
	// Timestamp 事件发生时间, 时间戳格式，毫秒
	Timestamp int64 `json:"timestamp,omitempty"`
	// ClickID 事件对应的track_id，同页面跳转track_id，进行回传，最长不超过30位
	ClickID string `json:"click_id,omitempty"`
	// Platform 电商平台
	// 京东电商平台，淘宝电商平台，客户名称+微信小程序（自定义命名，eg 小红书微信小程序）
	// - 注意：当电商平台 = 淘宝时，订单id需要回传【子订单号】，举例：假设1个父订单内含3个子订单，则 -> 回传3个不同的「子订单号」，每个子订单对应的购买件数=1，各自对应3条gmv数据。
	// - 当电商平台 = 京东时，订单id需要回传京东侧的【id】字段（订单+sku维度的唯一标识），实际业务逻辑与淘宝的【子订单号】一致。
	Platform string `json:"platform,omitempty"`
	// Context 转化回传上下文
	Context *Context `json:"context,omitempty"`
}

// Context 转化回传上下文
type Context struct {
	// Ad 广告上下文
	Ad Ad `json:"ad,omitempty"`
	// Product 商品信息
	Product *Product `json:"product,omitempty"`
}

// Ad 广告上下文
type Ad struct {
	// CallbackParam 广告唯一标识
	// 从监测链接的__CALLBACK_PARAM__获取
	// 1.callback不回传时，可以传空，不能传0
	// 2.callback在联调场景必传
	CallbackParam string `json:"callback_param"`
}

// Product 商品信息
type Product struct {
	// OrderID 电商订单id（子订单)
	OrderID string `json:"order_id,omitempty"`
	// OrderCount 购买件数
	OrderCount int `json:"order_count,omitempty"`
	// PayAmount 成交金额
	PayAmount int64 `json:"pay_amount,omitempty"`
	// ProductID 商品id
	ProductID string `json:"product_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductPrice 商品价格
	ProductPrice int64 `json:"product_price,omitempty"`
	// ProductCategory 商品类目
	ProductCategory string `json:"product_category,omitempty"`
	// ImageURL 商品大图
	ImageURL string `json:"image_url,omitempty"`
	// ShopName 店铺名称
	ShopName string `json:"shop_name,omitempty"`
}

// Encode implement PostRequest interface
func (r Request) Encode() []byte {
	return util.JSONMarshal(r)
}
