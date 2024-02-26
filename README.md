# 小红书聚光 MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/spotlight-mapi.svg)](https://pkg.go.dev/github.com/bububa/spotlight-mapi)
[![Go](https://github.com/bububa/spotlight-mapi/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/spotlight-mapi/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/spotlight-mapi/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/spotlight-mapi/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/spotlight-mapi.svg)](https://github.com/bububa/spotlight-mapi)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/spotlight-mapi)](https://goreportcard.com/report/github.com/bububa/spotlight-mapi)
[![GitHub license](https://img.shields.io/github/license/bububa/spotlight-mapi.svg)](https://github.com/bububa/spotlight-mapi/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/spotlight-mapi.svg)](https://GitHub.com/bububa/spotlight-mapi/releases/)


## API

- Oauth2.0授权(api/oauth)
  - 获取授权链接 [ URL(ctx context.Context, clt *core.SDKClient, req *oauth.URLRequest) string ]
  - 获取token [ AccessToken(ctx context.Context, clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.AccessToken, error) ]
  - 刷新token [ RefreshToken(ctx context.Context, clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.AccessToken, error) ]

- 账户服务
  - 广告主(api/advertiser)
    - 获取账号余额接口 [ BalanceInfo(ctx context.Context, clt *core.SDKClient, req *advertiser.BalanceInfoRequest, accessToken string) (*advertiser.Balance, error) ]

- 广告投放
  - 广告单元(api/unit)
    - 获取单元列表接口 [ List(ctx context.Context, clt *core.SDKClient, req *unit.ListRequest, accessToken string) (*unit.ListResult, error) ] 

- 数据报表
  - 实时报表(api/report/realtime)
    - 账户层级实时数据 [ Advertiser(ctx context.Context, clt *core.SDKClient, req *realtime.AdvertiserRequest, accessToken string) (*report.Metric, error) ]

- 转化追踪(api/conversion)
  - 生成点击监测链接 [ ClickMonitorLink(ctx context.Context, clt *core.SDKClient, req string) (string, error) ]
  - 转化回传 [ Conversion(ctx context.Context, clt *core.SDKClient, req *conversion.Request, accessToken string) error ]

