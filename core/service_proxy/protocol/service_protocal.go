package protocol

import (
	"context"
	"github.com/Bailing1992/apigateway/core/service_proxy/protocol/http"
	"github.com/valyala/fasthttp"
)

type ServiceProtocolInterface interface {
	Do(ctx context.Context, addr string, requestData *http.RequestData, requestContext *fasthttp.RequestCtx) int
}
