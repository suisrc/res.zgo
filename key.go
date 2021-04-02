package res

// 定义上下文中的键
const (
	prefix = "zgo"

	UserInfoKey = prefix + ":user-info"   // user info
	TraceIDKey  = prefix + ":tract-id"    // trace id
	ReqBodyKey  = prefix + ":req-body"    // request body
	ResBodyKey  = prefix + ":res-body"    // response body
	ResJwtKey   = prefix + ":res-jwt-kid" // jwt kid
	ResO2cKey   = prefix + ":res-o2c-kid" // o2c kid
	ResS2cKey   = prefix + ":res-s2c-kid" // s2c kid， 子母令牌， 标记母令牌ID

	XReqOriginHostKey   = "X-Request-Origin-Host"
	XReqOriginPathKey   = "X-Request-Origin-Path"
	XReqOriginMethodKey = "X-Request-Origin-Method"
)

const (
	// ShowNone 静音
	ShowNone = 0
	// ShowWarn 消息警告
	ShowWarn = 1
	// ShowError 消息错误
	ShowError = 2
	// ShowNotify 通知；
	ShowNotify = 4
	// ShowPage 页
	ShowPage = 9
)
