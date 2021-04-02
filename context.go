package res

import (
	"context"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Context interface {
	GetTraceID() string
	FormatMessage(emsg *i18n.Message, args map[string]interface{}) string
}

// 跳过控制器

type ReqContext interface {
	context.Context
	GetRequest() *http.Request // 请求头
	GetHeader(string) string   // 获取请求头
	Next()                     // 执行下一个
	Abort()                    // 终止执行
}
