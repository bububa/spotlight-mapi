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
    - 获取accessToken [ AccessToken(ctx context.Context, clt *core.SDKClient, advertiserID uint64) (*oauth.AccessToken, error) ]
    - 获取账号余额接口 [ BalanceInfo(ctx context.Context, clt *core.SDKClient, req *advertiser.BalanceInfoRequest, accessToken string) (*advertiser.Balance, error) ]
    - 账户白名单 [ WhiteList(ctx context.Context, clt *core.SDKClient, req *advertiser.WhiteListRequest, accessToken string) (*advertiser.WhiteList, error) ]

- 广告投放
  - 广告计划(api/campaign)
    - 创建计划 [ Create(ctx context.Context, clt *core.SDKClient, req *campaign.CreateRequest, accessToken string) (uint64, error) ]
    - 编辑计划 [ Update(ctx context.Context, clt *core.SDKClient, req *campaign.UpdateRequest, accessToken string) (uint64, error) ]
    - 修改计划状态 [ StatusUpdate(ctx context.Context, clt *core.SDKClient, req *campaign.StatusUpdateRequest, accessToken string) ([]uint64, error) ]
    - 查询计划 [ List(ctx context.Context, clt *core.SDKClient, req *campaign.ListRequest, accessToken string) (*campaign.ListResult, error) ]
  - 广告单元(api/unit)
    - 创建单元 [ Create(ctx context.Context, clt *core.SDKClient, req *unit.CreateRequest, accessToken string) (uint64, error) ]
    - 编辑单元 [ Update(ctx context.Context, clt *core.SDKClient, req *unit.UpdateRequest, accessToken string) (uint64, error) ]
    - 修改单元状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, req *unit.UpdateStatusRequest, accessToken string) ([]uint64, error) ]
    - 获取单元列表接口 [ List(ctx context.Context, clt *core.SDKClient, req *unit.ListRequest, accessToken string) (*unit.ListResult, error) ] 
  - 广告创意(api/creativity)
    - 创建笔记创意 [ Create(ctx context.Context, clt *core.SDKClient, req *creativity.CreateRequest, accessToken string) (uint64, error) ]
    - 编辑创意 [ Update(ctx context.Context, clt *core.SDKClient, req *creativity.UpdateRequest, accessToken string) (uint64, error) ]
    - 修改创意状态 [ StatusUpdate(ctx context.Context, clt *core.SDKClient, req *creativity.StatusUpdateRequest, accessToken string) ([]uint64, error) ]
    - 创意查询 [ Search(ctx context.Context, clt *core.SDKClient, req *creativity.SearchRequest, accessToken string) (*creativity.SearchResult, error) ]

- 数据报表
  - 离线报表(api/report/offline)
    - 账户层级离线报表数据 [ Advertiser(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) ]
    - 计划层级离线报表数据 [ Campaign(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) ]
    - 单元层级离线报表数据 [ Unit(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) ]
    - 创意层级离线报表数据 [ Creativity(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) ]
    - 关键词层级离线报表数据 [ Keyword(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) ]
  - 实时报表(api/report/realtime)
    - 账户层级实时数据 [ Advertiser(ctx context.Context, clt *core.SDKClient, req *realtime.AdvertiserRequest, accessToken string) (*report.Metric, error) ]
    - 计划层级实时数据 [ Campaign(ctx context.Context, clt *core.SDKClient, req *realtime.CampaignRequest, accessToken string) (*realtime.CampaignResponse, error) ]
    - 单元层级实时数据 [ Unit(ctx context.Context, clt *core.SDKClient, req *realtime.UnitRequest, accessToken string) (*realtime.UnitResponse, error) ]
    - 创意层级实时数据 [ Creativity(ctx context.Context, clt *core.SDKClient, req *realtime.CreativityRequest, accessToken string) (*realtime.CreativityResponse, error) ]
    - 关键词层级实时数据 [ Keyword(ctx context.Context, clt *core.SDKClient, req *realtime.KeywordRequest, accessToken string) (*realtime.KeywordResponse, error) ]

- 素材管理
  - 获取直达链接 [ directlink.List(ctx context.Context, clt *core.SDKClient, req *directlink.ListRequest, accessToken string) (*directlink.ListResult, error) ]
  - 删除直达链接 [ directlink.Delete(ctx context.Context, clt *core.SDKClient, req *directlink.DeleteRequest, accessToken string) error ]
  - 创建直达链接 [ directlink.Create(ctx context.Context, clt *core.SDKClient, req *directlink.CreateRequest, accessToken string) ([]directlink.DirectLink, error) ]

- 转化追踪(api/conversion)
  - 生成点击监测链接 [ ClickMonitorLink(ctx context.Context, clt *core.SDKClient, req string) (string, error) ]
  - 外链落地页
      - 外链线索数据回传 [ Conversion(ctx context.Context, clt *core.SDKClient, req *conversion.Request, accessToken string) error ]
      - 线索转化数据回传 [ AuroraLeads(ctx context.Context, clt *core.SDKClient, req *conversion.AuroraLeadsRequest) error ]
  - 聚光落地页
      - 聚光落地页线索数据回传 [ HawkingLeads(ctx context.Context, clt *core.SDKClient, req *conversion.HawkingLeadsRequest) error ]
  - 口令码
      - APP口令码数据回传 [ App(ctx context.Context, clt *core.SDKClient, req *conversion.AppRequest) error ]

