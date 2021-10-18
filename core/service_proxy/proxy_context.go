package service_proxy

import (
	"github.com/Bailing1992/apigateway/core/service_proxy/protocol/http"
	"github.com/valyala/fasthttp"
	"sync"
)

type ServiceProxyContext struct {
	RequestContext *fasthttp.RequestCtx
	RequestData    *http.RequestData
	proxy          *ServiceProxy
	server         *ServiceInstance
}

func (proxyContext *ServiceProxyContext) Abort(errCode int) int {
	return 0
}

type ServiceInstance struct {
	Addr         string
	RawWeightStr string
	RawWeight    int
	Weight       int
	Status       int
	ConnCount    int
	SuccessCount int64
	ErrorCount   int64
	Tags         map[string]string
	lock         sync.Mutex
}

func (proxyContext *ServiceProxyContext) ServerAddr() (string, bool) {
	if proxyContext.server != nil {
		return proxyContext.server.Addr, true
	} else {
		return "", false
	}
}
