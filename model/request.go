package model

import (
	"io"
)

type Request interface {
	// SetApp set appid and access token to request
	SetApp(appID string, accessToken string)
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
	AppID string `json:"app_id,omitempty"`
	// AccessToken 访问令牌
	AccessToken string `json:"access_token,omitempty"`
}

func (r *BaseRequest) SetApp(appID string, accessToken string) {
	r.AppID = appID
	r.AccessToken = accessToken
}
