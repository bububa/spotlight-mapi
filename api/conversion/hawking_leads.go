package conversion

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/conversion"
)

// HawkingLeads 聚光落地页线索数据回传
func HawkingLeads(ctx context.Context, clt *core.SDKClient, req *conversion.HawkingLeadsRequest) error {
	return clt.PostHawkingLeads(ctx, req)
}
