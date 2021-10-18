package router

import (
	"context"
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/valyala/fasthttp"
)

type RouteInfoFinder interface {
	Find(ctxt context.Context, requestCtx *fasthttp.RequestCtx) RouteInfo

	Iterate(f func(RouteInfo))
}

type RouteInfoConfigurator interface {
	// Add a route info into this configurator.
	// The route info will then be able to be found with the corresponding RouteInfoFinder.
	Add(ctxt context.Context, info RouteInfo) error

	// Mark all RouteInfo has been added via the Add method.
	// This method should be called before put into use, because some implementations will have to
	// do some extra work on the added RouteInfos before it is put into use.
	ConfigDone(ctxt context.Context) error
}

type ConfigurableRouter interface {
	RouteInfoConfigurator
	RouteInfoFinder
}

// 路由查找的结果
type RouteInfo interface {
	// 最终路由的目标 psm, 仅用于调试日志
	PSM() string
	// 最终路由的目标 serviceId
	ServiceId() string
	// 服务的配置
	GetProxyConfig() *config.ProxyServiceConfig
	// 路由的配置
	RouteConfig() *config.RouteConfig
}
