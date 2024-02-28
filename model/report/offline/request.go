package offline

import (
	"github.com/bububa/spotlight-mapi/enum"
	"github.com/bububa/spotlight-mapi/util"
)

// Request 离线报表数据 API Request
type Request struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度：
	// "DAY"：分天
	// "HOUR"：分时
	// "SUMMARY"：汇总,默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大1000
	PageSize int64 `json:"page_size,omitempty"`
	// MarketingStrategy 营销诉求过滤条件：
	// 3：商品销售日常推广
	// 4：产品种草
	// 8：直播推广_日常推广
	// 9：客资收集
	// 10：抢占赛道
	// 14：直播推广_直播预热
	MarketingStrategy []int `json:"marketing_strategy,omitempty"`
	// BiddingStrategy 出价方式过滤条件：2：手动出价3：自动出价
	BiddingStrategy []int `json:"bidding_strategy,omitempty"`
	// OptimizeTarget 推广目标过滤条件：
	// 0：点击量
	// 1：互动量
	// 3：表单提交量
	// 4：商品成单量
	// 5：私信咨询量
	// 6：直播间观看量
	// 11：商品访客量
	// 12：落地页访问量
	// 13：私信开口量
	// 14：有效观看量
	// 18：站外转化量
	// 20：TI人群规模
	// 21：行业商品成单
	// 23：直播预热量
	// 24：直播间成交
	// 25：直播间支付ROI
	OptimizeTarget []int `json:"optimize_target,omitempty"`
	// Placement 广告类型过滤条件
	// 1：信息流推广
	// 2：搜索推广
	// 4：全站智投
	// 7：视频内流
	Placement []int `json:"placement,omitempty"`
	// PromotionTarget 推广标的类型过滤条件
	// 1：笔记
	// 2：商品
	// 9：落地页
	// 18：直播间
	// 商品、落地页、直播间笔记只能三选一
	PromotionTarget []int `json:"promotion_target,omitempty"`
	// BuildType 搭建方式过滤条件
	// 0：标准搭建
	// 1：省心智投
	BuildType []int `json:"build_type,omitempty"`
	// SplitColumns 细分条件(相当于group by)
	// marketingTarget：营销诉求
	// buildType：搭建方式
	// placement：广告类型
	// optimizeTarget：推广目标
	// biddingStrategy：出价方式
	// promotionTarget：推广标的类型
	// jumpType：创意跳转类型
	// itemId：商品
	// pageId：落地页
	// liveRedId：直播间笔记
	// keyword：关键词
	SplitColumns []string `json:"split_columns,omitempty"`
	// Programmatic 创意组合方式过滤条件
	// 0：自定义创意
	// 1：程序化创意
	Programmatic []int `json:"programmatic,omitempty"`
	// MarketingTarget 营销诉求过滤条件：3：商品销量_日常推广4：产品种草8：直播推广_日常推广9：客资收集10：抢占赛道14：直播推广_直播预热
	MarketingTarget []int `json:"marketing_target,omitempty"`
}

// Encode implement PostRequest interface
func (r Request) Encode() []byte {
	return util.JSONMarshal(r)
}
