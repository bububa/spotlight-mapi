package keyword

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// IndustryTaxonomyRequest 行业类目 API Request
type IndustryTaxonomyRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r IndustryTaxonomyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// IndustryTaxonomyResponse 行业类目 API Response
type IndustryTaxonomyResponse struct {
	Data *IndustryTaxonomyResult `json:"data,omitempty"`
	model.BaseResponse
}

type IndustryTaxonomyResult struct {
	// AdsIndustryTaxonomyDictDto 行业细节
	List *Taxonomy `json:"ads_industry_taxonomy_dict_dto,omitempty"`
	// AllIndustryTaxonomys 所有可选的行业类目
	AllIndustryTaxonomys string `json:"all_industry_taxonomys,omitempty"`
}

// Taxonomy 行业细节
type Taxonomy struct {
	// TaxonomyID 行业id
	TaxonomyID string `json:"taxonomy_id,omitempty"`
	// TaxonomyName 行业名称
	TaxonomyName string `json:"taxonomy_name,omitempty"`
	// FullPathName 全路径名
	FullPathName string `json:"full_path_name,omitempty"`
	// Children 子层级
	Children []Taxonomy `json:"children,omitempty"`
	// TaxonomyLevel 行业层级
	TaxonomyLevel int `json:"taxonomy_level,omitempty"`
}
