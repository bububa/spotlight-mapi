package conversion

import (
	"context"
	"net/url"
)

// ClickMonitorLink 生成点击监测链接
func ClickMonitorLink(ctx context.Context, req string) (string, error) {
	link, err := url.ParseRequestURI(req)
	if err != nil {
		return "", err
	}
	query := link.Query()
	query.Set("click_id", "__CLICK_ID__")
	query.Set("callback_param", "__CALLBACK_PARAM__")
	query.Set("advertiser_id", "__ADVERTISER_ID__")
	query.Set("campaign_id", "__CAMPAIGN_ID__")
	query.Set("unit_id", "__UNIT_ID__")
	query.Set("creativity_id", "__CREATIVITY_ID__")
	link.RawQuery = query.Encode()
	return link.String(), nil
}
