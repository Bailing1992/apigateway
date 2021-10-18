package service_proxy

import (
	"context"
	"github.com/Bailing1992/apigateway/consts"
)

func ProtocolHandler(ctx context.Context, proxyInstance *ServiceProxy, proxyContext *ServiceProxyContext) int {
	proxyContext.RequestContext.SetUserValue(consts.KeyErrorType, nil)
	proxyContext.RequestContext.SetUserValue(consts.KeyErrorCode, nil)

	// standalone代理模式下使用http协议做转发; sidecar代理模式在另外调用链处理了
	serviceProtocol := proxyInstance.serviceProtocol

	addr, _ := proxyContext.ServerAddr()
	errCode := serviceProtocol.Do(ctx, addr, proxyContext.RequestData, proxyContext.RequestContext)
	if errCode != consts.StatusCodeSuccess {
		return proxyContext.Abort(errCode)
	} else if proxyContext.RequestContext.Response.StatusCode() >= 400 {
		//proxyContext.RequestContext.SetUserValue(consts.KeyErrorType, consts.StatusDescOf(consts.StatusCodeBusinessError))
		proxyContext.RequestContext.SetUserValue(consts.KeyErrorCode, consts.StatusCodeBusinessError)
	}
	return consts.StatusCodeSuccess
}
