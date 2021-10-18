package service_proxy

import "sync"

var proxyContextPool = sync.Pool{
	New: func() interface{} { return &ServiceProxyContext{} },
}

func BorrowProxyContext() *ServiceProxyContext {
	ctx := proxyContextPool.Get().(*ServiceProxyContext)
	ResetProxyContext(ctx)
	return ctx
}

func ResetProxyContext(ctx *ServiceProxyContext) {
	ctx.RequestContext = nil
	ctx.proxy = nil
}

func ReturnProxyContext(ctx *ServiceProxyContext) {
	ResetProxyContext(ctx)
	proxyContextPool.Put(ctx)
}
