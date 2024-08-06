package keyword

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// IndustryTaxonomyAttributeRequest 行业类目属性 API Request
type IndustryTaxonomyAttributeRequest struct {
	// TaxomonyID 行业类目id,通过/api/open/jg/keyword/industry/taxonomy获取
	TaxomonyID string `json:"taxonomy_id,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r IndustryTaxonomyAttributeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// IndustryTaxonomyAttributeResponse 行业类目属性 API Response
type IndustryTaxonomyAttributeResponse struct {
	Data *IndustryTaxonomyAttributeResult `json:"data,omitempty"`
	model.BaseResponse
}

type IndustryTaxonomyAttributeResult struct {
	// List 行业类目属性
	List []*TaxonomyAttribute `json:"taxonomy_attribute_dtos,omitempty"`
}

// TaxonomyAttribute 行业类目属性
type TaxonomyAttribute struct {
	// TaxonomyID 行业类目id
	TaxonomyID string `json:"taxonomy_id,omitempty"`
	// TaxonomyAttributeName 行业类目名称
	TaxonomyAttributeName string `json:"taxonomy_attribute_name,omitempty"`
	// TaxonomyLevel 层级
	TaxonomyLevel int `json:"taxonomy_level,omitempty"`
}
