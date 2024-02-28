package model

import (
	"io"

	"github.com/bububa/spotlight-mapi/util"
)

type Request interface {
}

// PostRequest post request interface
type PostRequest interface {
	Request
	// Encode encode request to bytes
	Encode() []byte
}

// GetRequest get request interface
type GetRequest interface {
	Request
	// Encode encode request to string
	Encode() string
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Key field key
	Key string
	// Value field value
	Value string
	// Reader upload file reader
	Reader io.Reader
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	Request
	// Encode encode request to UploadFields
	Encode() []UploadField
}

type BaseRequest struct {
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// Secret 应用secret
	Secret string `json:"secret,omitempty"`
}

// CommonRequest 通用 API Request
type CommonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Method 请求方法
	Method string `json:"method,omitempty"`
}

// Encode implement PostRequest interface
func (r CommonRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
