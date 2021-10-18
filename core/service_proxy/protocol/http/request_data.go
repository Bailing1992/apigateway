package http

import (
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/bitly/go-simplejson"
	"github.com/valyala/fasthttp"
	"reflect"
	"sync"
)

type RequestData struct {
	// http使用copy的方式,需要将新的header设置到RequestCtx中
	// TODO 统一获取header的方式
	requestContext *fasthttp.RequestCtx
	// 请求基本信息结构化
	URI         string `json:"uri"`
	Method      string `json:"method"`
	proxyConfig *config.TimeoutConfig
	Headers     map[string]string   `json:"headers"`
	Cookies     map[string]string   `json:"cookies"`
	QueryArgs   map[string][]string `json:"query_args"`
	PostArgs    map[string][]string `json:"post_args"`
	Body        []byte              `json:"body"`
	ClientIP    string              `json:"client_ip"`
	RequestTime int64               `json:"request_time"`
	SessionJson *simplejson.Json    `json:"session"`
	UriMatch    *UriMatch

	userValues     sync.Map
	ExtraHeaders   map[string]string
	mutexSetHeader sync.RWMutex `json:"-"`
	HeaderAdapter
}

func NewRequestData(requestContext *fasthttp.RequestCtx) *RequestData {
	req := &RequestData{
		requestContext: requestContext,
		Headers:        make(map[string]string),
		Cookies:        make(map[string]string),
		QueryArgs:      make(map[string][]string),
		PostArgs:       make(map[string][]string),
		UriMatch:       &UriMatch{},
	}

	req.HeaderAdapter.NewHeaderAdapter(&req.Headers, &req.ExtraHeaders)
	return req
}

type UriMatch struct {
	Path   string
	Params map[string]string
}

func (requestData *RequestData) ForAllExtraHeader(f func(name, value string)) {
	requestData.mutexSetHeader.RLock()
	defer requestData.mutexSetHeader.RUnlock()

	for name, value := range requestData.HeaderAdapter.ExtraHeaders() {
		f(name, value)
	}
}

func (requestData *RequestData) SetUserValue(key, value interface{}) {
	// the logic of entrance test is borrowed from Context
	if key == nil {
		panic("nil key unsupported")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	requestData.userValues.Store(key, value)
}
