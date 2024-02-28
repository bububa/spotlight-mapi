package oauth

// Advertiser 广告主
type Advertiser struct {
	// ID 广告主ID
	ID uint64 `json:"advertiser_id,omitempty"`
	// Name 广告主名称
	Name string `json:"advertiser_name,omitempty"`
}
