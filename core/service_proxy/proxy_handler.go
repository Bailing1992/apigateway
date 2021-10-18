package service_proxy

import "context"

type Handler func(context.Context, *ServiceProxy, *ServiceProxyContext) int

type NamedHandler struct {
	Name    string
	Handler Handler
}

var ProxyHandlers []NamedHandler

func init() {
	ProxyHandlers = []NamedHandler{
		{Name: "ProtocolHandler", Handler: ProtocolHandler},
	}
}
