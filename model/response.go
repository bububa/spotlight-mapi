package model

import (
	"strconv"

	"github.com/bububa/spotlight-mapi/util"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	// Success 是否成功
	Success bool `json:"success,omitempty"`
	// ErrorCode 返回码
	ErrorCode int `json:"errorCode"`
	// ErrorMsg 返回信息
	ErrorMsg string `json:"errorMsg"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return !r.Success
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	return util.StringsJoin(strconv.Itoa(r.ErrorCode), ":", r.ErrorMsg)
}
