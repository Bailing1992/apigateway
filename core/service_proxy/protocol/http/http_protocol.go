package http

import (
	"context"
	"github.com/Bailing1992/apigateway/consts"
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/Bailing1992/apigateway/core/errors"
	"github.com/Bailing1992/apigateway/utils"
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

type HttpProtocol struct {
	client        *fasthttp.Client
	timeout       uint32
	timeoutConfig *config.TimeoutConfig
}

type HttpTrafficEnv struct {
	Open bool   `json:"open"`
	Env  string `json:"env"`
}

type HttpExtra map[string]string

type HttpRequestBase struct {
	LogID      string         `json:"log_id"`
	Caller     string         `json:"caller"`
	Addr       string         `json:"addr"`
	Client     string         `json:"client"`
	TrafficEnv HttpTrafficEnv `json:"traffic_env,omitempty"`
	Extra      HttpExtra      `json:"extra,omitempty"`
}

func NewHttpProtocol(proxyConfig *config.TimeoutConfig) *HttpProtocol {
	return &HttpProtocol{
		client:        buildHttpClient(proxyConfig),
		timeout:       getTimeout(proxyConfig),
		timeoutConfig: proxyConfig,
	}
}

func buildHttpClient(proxyConfig *config.TimeoutConfig) *fasthttp.Client {
	maxIdleConnDur := 5 * time.Second
	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			conn, err := net.DialTimeout("tcp", addr, utils.Uint32ToTimeMillisecond(proxyConfig.ConnectTimeout))
			if err != nil {
				return nil, errors.NewProxyError("Dial to http backend server failed, addr: %s, error: %v", addr, err)
			}
			return conn, nil
		},
		ReadTimeout:         utils.Uint32ToTimeMillisecond(proxyConfig.ReadTimeout),
		WriteTimeout:        utils.Uint32ToTimeMillisecond(proxyConfig.WriteTimeout),
		MaxConnsPerHost:     10000,
		MaxIdleConnDuration: maxIdleConnDur,
		ReadBufferSize:      32 * 1024, // 不能改，因为读响应的时候需要该buffer存放完整的header，否则就会报错，所以会导致header过大的响应报错
	}
}

func getTimeout(serviceTimeout *config.TimeoutConfig) uint32 {
	connectTimeout := serviceTimeout.ConnectTimeout
	writeTimeout := serviceTimeout.WriteTimeout
	readTimeout := serviceTimeout.ReadTimeout
	return connectTimeout + writeTimeout + readTimeout
}

func (h *HttpProtocol) Do(ctx context.Context, addr string, requestData *RequestData, requestContext *fasthttp.RequestCtx) (errCode int) {
	fastHttpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(fastHttpReq)

	requestData.SetUserValue(consts.KeyMethod, string(requestContext.Request.Header.Method()))

	p := &requestContext.Request
	p.CopyTo(fastHttpReq)

	requestData.ForAllExtraHeader(func(name, value string) {
		fastHttpReq.Header.Set(name, value)
	})

	fastHttpResp := &requestContext.Response

	requestContext.SetUserValue(consts.KeyDirectUpstreamAddr, addr)
	fastHttpReq.SetHost(addr)
	fastHttpReq.URI().SetScheme("http")
	// 底层重试最多1次
	if err := h.client.Do(fastHttpReq, fastHttpResp); err != nil {

		requestData.SetUserValue(consts.KeyErrorMessage, err.Error())
		// 无结果
		switch err {
		case fasthttp.ErrNoFreeConns:
			return consts.StatusCodeGetConnError
		case fasthttp.ErrTimeout:
			return consts.StatusCodeRPCTimeout
		default:
			return consts.StatusCodeBusinessError
		}
	}
	return consts.StatusCodeSuccess

}
