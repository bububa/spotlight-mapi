package conversion

import (
	"strconv"
	"time"

	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// AppRequest APP口令码数据回传 API Request
type AppRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID model.JSONUint64 `json:"advertiser_id,omitempty"`
	// Version 接口版本, 固定为"1.0"
	Version string `json:"version,omitempty"`
	// Timestamp 当前时间戳, 时间戳格式，毫秒
	Timestamp int64 `json:"timestamp,omitempty"`
	// OaidMd5 安卓广告标识符, 高版本，32位大写md5加密
	OaidMd5 string `json:"oaid_md5,omitempty"`
	// CaidMd5 苹果广告标识符(20230330版), 高版本，32位大写md5加密
	CaidMd5 string `json:"caid_md5,omitempty"`
	// TokenCode 口令码
	TokenCode string `json:"token_code,omitempty"`
	// EventType 转化事件类型id
	EventType enum.EventType `json:"event_type,omitempty"`
	// ConvTime 转化事件发生时间
	ConvTime int64 `json:"conv_time,omitempty"`
	// ReportSource 数据来源
	ReportSource string `json:"report_source,omitempty"`
	// ForTest 是否是联调环境
	ForTest bool `json:"for_test,omitempty"`
	// AccessToken access_token
	AccessToken string `json:"access_token,omitempty"`
	// Sign 签名
	Sign string `json:"sign,omitempty"`
}

// Encode implement PostRequest interface
func (r AppRequest) Encode() []byte {
	r.Timestamp = time.Now().UnixMilli()
	r.Sign = r.sign()
	return util.JSONMarshal(r)
}

func (r AppRequest) sign() string {
	raw := util.StringsJoin("advertiser_id", r.AdvertiserID.String(), "&timestamp", strconv.FormatInt(r.Timestamp, 10), "&version", r.Version)
	return util.Md5String(raw)
}
