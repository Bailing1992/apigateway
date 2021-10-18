package service_proxy

import (
	"context"
	"github.com/Bailing1992/apigateway/consts"
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/Bailing1992/apigateway/core/errors"
	"github.com/Bailing1992/apigateway/core/service_proxy/protocol"
	"github.com/Bailing1992/apigateway/core/service_proxy/protocol/http"
	"sync"
)

type ServiceProxy struct {
	PSM              string
	proxyConfig      *config.ProxyServiceConfig
	serviceProtocol  protocol.ServiceProtocolInterface
	proxyContextPool sync.Pool
}

func NewServiceProxy(ctx context.Context, proxyConfig *config.ProxyServiceConfig) (*ServiceProxy, error) {
	proxyInstance := &ServiceProxy{
		PSM:         proxyConfig.PSM,
		proxyConfig: proxyConfig,
	}
	proxyInstance.proxyContextPool.New = func() interface{} { return &ServiceProxyContext{} }

	// protocol
	if err := proxyInstance.initProtocol(ctx); err != nil {
		return nil, err
	}
	return proxyInstance, nil
}

func (proxyInstance *ServiceProxy) initProtocol(ctx context.Context) error {
	switch proxyInstance.proxyConfig.ServiceProtocol {
	case consts.HttpProtocolType:
		proxyInstance.serviceProtocol = http.NewHttpProtocol(&proxyInstance.proxyConfig.ServiceTimeout)
	default:
		err := errors.NewProxyError("Unknown service protocol, protocol_type: %s", proxyInstance.proxyConfig.ServiceProtocol)
		return err
	}
	return nil

}
