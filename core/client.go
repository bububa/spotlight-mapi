package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bububa/spotlight-mapi/core/internal/debug"
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	limiter RateLimiter
	client  *http.Client
	tracer  *Otel
	secret  string
	appID   uint64
	debug   bool
	sandbox bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID uint64, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
		client: defaultHttpClient(),
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.AppID())
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// SetRateLimiter 设置限流
func (c *SDKClient) SetRateLimiter(limiter RateLimiter) {
	c.limiter = limiter
}

func (c *SDKClient) AppID() uint64 {
	return c.appID
}

func (c *SDKClient) Secret() string {
	return c.secret
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		appID:   c.appID,
		secret:  c.secret,
		debug:   c.debug,
		sandbox: c.sandbox,
		client:  c.client,
	}
}

// Post post api
func (c *SDKClient) Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	var reqUrl string
	if strings.HasPrefix(gw, "https://") {
		reqUrl = gw
	} else {
		reqUrl = util.StringsJoin(BASE_URL, gw)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	var reqUrl string
	if strings.HasPrefix(gw, "https://") {
		reqUrl = gw
	} else {
		reqUrl = util.StringsJoin(BASE_URL, gw)
	}
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	return c.WithSpan(ctx, httpReq, resp, nil, c.fetch)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(ctx context.Context, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	buf := util.NewBufferPool()
	defer util.ReleaseBufferPool(buf)
	mw := multipart.NewWriter(buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			mp[v.Key] = util.StringsJoin("@", v.Value)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	var reqUrl string
	if strings.HasPrefix(gw, "https://") {
		reqUrl = gw
	} else {
		reqUrl = util.StringsJoin(BASE_URL, gw)
	}
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	bs, _ := json.Marshal(mp)
	return c.WithSpan(ctx, httpReq, resp, bs, c.fetch)
}

// PostHawkingLeads
func (c *SDKClient) PostHawkingLeads(ctx context.Context, req model.PostRequest) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, HAWKLING_LEADS_URL, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	debug.PrintJSONRequest("POST", HAWKLING_LEADS_URL, httpReq.Header, reqBytes, c.debug)
	return c.WithSpan(ctx, httpReq, nil, reqBytes, c.fetch)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) (*http.Response, error) {
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	body, err := debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if httpResp.ContentLength <= 0 {
		httpResp.ContentLength = int64(len(body))
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, errors.Join(model.BaseResponse{
			ErrorCode: httpResp.StatusCode,
			ErrorMsg:  string(body),
		}, err)
	}
	if resp.IsError() {
		return httpResp, resp
	}
	return httpResp, nil
}

func (c *SDKClient) WithSpan(ctx context.Context, req *http.Request, resp model.Response, payload []byte, fn func(*http.Request, model.Response) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, req, resp, payload, fn)
}
