package conversion

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/conversion"
)

// Conversion 转化回传API
// 回传您通过小红书获得的转化信息，我们可以提供如下服务：
// 1.展现回传事件的统计信息，例如回传”商品成单“时间，我们会by计划/创意等维度呈现订单量、订单成本、转化率等指标
// 2.数据进入您的广告模型，优化您的广告效果。
func Conversion(ctx context.Context, clt *core.SDKClient, req *conversion.Request) error {
	return clt.Post(ctx, "/conversion", req, nil, "")
}
