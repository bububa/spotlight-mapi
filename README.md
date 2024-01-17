# 小红书聚光 MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/spotlight-mapi.svg)](https://pkg.go.dev/github.com/bububa/spotlight-mapi)
[![Go](https://github.com/bububa/spotlight-mapi/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/spotlight-mapi/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/spotlight-mapi/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/spotlight-mapi/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/spotlight-mapi.svg)](https://github.com/bububa/spotlight-mapi)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/spotlight-mapi)](https://goreportcard.com/report/github.com/bububa/spotlight-mapi)
[![GitHub license](https://img.shields.io/github/license/bububa/spotlight-mapi.svg)](https://github.com/bububa/spotlight-mapi/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/spotlight-mapi.svg)](https://GitHub.com/bububa/spotlight-mapi/releases/)


## API

- 转化追踪(api/conversion)
  - 生成点击监测链接 [ ClickMonitorLink(ctx context.Context, req string) (string, error) ]
  - 转化回传 [ Conversion(ctx context.Context, clt *core.SDKClient, req *conversion.Request) error ]

