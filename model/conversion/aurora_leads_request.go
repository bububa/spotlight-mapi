package conversion

import (
	"strconv"
	"time"

	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// AuroraLeadsRequest 回传数据接口 API Request
type AuroraLeadsRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID model.JSONUint64 `json:"advertiser_id,omitempty"`
	// Version 接口版本, 固定为"1.0"
	Version string `json:"version,omitempty"`
	// Timestamp 当前时间戳, 时间戳格式，毫秒
	Timestamp int64 `json:"timestamp,omitempty"`
	// Method 转发方法, 固定为"aurora.leads"
	Method string ` json:"method,omitempty"`
	// Token access_token
	Token string `json:"token,omitempty"`
	// EventType 转化事件类型 id
	EventType enum.EventType `json:"event_type,omitempty"`
	// ConvTime 转化发生时间, 时间戳格式，毫秒
	ConvTime int64 `json:"conv_time,omitempty"`
	// ClickID 事件对应的click_id，同页面跳转click_id
	ClickID string `json:"click_id,omitempty"`
	// AccessToken access_token
	AccessToken string `json:"access_token,omitempty"`
	// Sign 签名
	Sign string `json:"sign,omitempty"`
}

// Encode implement PostRequest
func (r AuroraLeadsRequest) Encode() []byte {
	r.Version = "1.0"
	r.Method = "aurora.leads"
	r.Timestamp = time.Now().UnixMilli()
	r.Sign = r.sign()
	return util.JSONMarshal(r)
}

func (r AuroraLeadsRequest) sign() string {
	raw := util.StringsJoin("advertiser_id", r.AdvertiserID.String(), "&method", r.Method, "&timestamp", strconv.FormatInt(r.Timestamp, 10), "&version", r.Version)
	return util.Md5String(raw)
}
