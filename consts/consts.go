package consts

const (
	KeyMethod             = "call.method"          // 调用的方法
	KeyDirectUpstreamAddr = "direct_upstream_addr" // agw直连的下游地址,也就是rip
	KeyErrorCode          = "errorcode"
	KeyErrorTip           = "error.tip"
	KeyErrorType          = "errortype"
	KeyErrorMessage       = "error.msg"
)

const (
	StatusCodeBusinessError              = -1
	StatusCodeSuccess                    = 0 // SuccessCode Success
	StatusCodeNotAllowedByServiceCB      = 101
	StatusCodeNotAllowedByInstanceCB     = 102
	StatusCodeRPCTimeout                 = 103
	StatusCodeNotAllowedByDegradation    = 104
	StatusCodeGetDegradationPercentError = 105
	StatusCodeEmptyInstanceList          = 106
	StatusCodeNoMoreInstance             = 107
	StatusCodeConnRetry                  = 108
	StatusCodeRPCRetry                   = 110
	StatusCodeGetConnError               = 112
	StatusCodeServiceDiscoverError       = 113
)

const (
	HttpProtocolType   = "HTTP"
	ThriftProtocolType = "THRIFT"
)
