package data

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// QualInfoRequest 获取资质列表 API Request
type QualInfoRequest struct {
	// Page
	Page *model.PageDTO `json:"page,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r QualInfoRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// QualInfoResponse 获取资质列表 API Response
type QualInfoResponse struct {
	// Data
	Data *QualInfoResult `json:"data,omitempty"`
	model.BaseResponse
}

type QualInfoResult struct {
	// QualInfos 资质信息
	QualInfos []QualInfo `json:"qual_infos,omitempty"`
	// ProductQualInfos 行业资质信息
	ProductQualInfos []ProductQualInfo `json:"product_qual_infos,omitempty"`
	// BrandQualInfos 产品资质信息
	BrandQualInfo []BrandQualInfo `json:"brand_qual_infos,omitempty"`
}

// QualInfo 资质信息
type QualInfo struct {
	// ApplyID 资质id
	ApplyID string `json:"apply_id,omitempty"`
	// TradeTypeFirstCode 一级行业编码
	TradeTypeFirstCode string `json:"trade_type_first_code,omitempty"`
	// TradeTypeFirstName 一级行业名称
	TradeTypeFirstName string `json:"trade_type_first_name,omitempty"`
	// TradeTypeSecondCode 二级行业编码
	TradeTypeSecondCode string `json:"trade_type_second_code,omitempty"`
	// TradeTypeSecondName 二级行业名称
	TradeTypeSecondName string `json:"trade_type_second_name,omitempty"`
}

// ProductQualInfo 行业资质信息
type ProductQualInfo struct {
	// ProductName 产品资质名称
	ProductName string `json:"product_name,omitempty"`
	// ProductQualID 产品资质id
	ProductQualID int `json:"product_qual_id,omitempty"`
}

// BrandQualInfo 产品资质信息
type BrandQualInfo struct {
	// QualTypeName 	资质名称
	QualTypeName string `json:"qual_type_name,omitempty"`
	// UserRemark 备注
	UserRemark string `json:"user_remark,omitempty"`
	// BrandQualID 品牌资质id
	BrandQualID int `json:"brand_qual_id,omitempty"`
}
