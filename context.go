package res

import (
	"context"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Context interface {
	context.Context

	GetTraceID() string
	FormatMessage(emsg *i18n.Message, args map[string]interface{}) string

	GetRequest() *http.Request // 请求头
	GetHeader(string) string   // 获取请求头

	Next()  // 执行下一个
	Abort() // 终止执行

	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// GetAcceptLanguage 获取浏览器语言
func GetAcceptLanguage(c Context) string {
	if v := c.GetHeader("Accept-Language"); v != "" {
		if len := strings.Index(v, ","); len > 0 {
			v = v[:len]
		}
		if len := strings.Index(v, ";"); len > 0 {
			v = v[:len]
		}
		return v
	}
	return ""
}

// GetCtxValue 获取令牌加密方式
func GetCtxValue(c Context, key string) (interface{}, bool) {
	return c.Get(key)
}

// GetCtxValueToString 获取令牌加密方式
func GetCtxValueToString(c Context, key string) (string, bool) {
	if v, ok := c.Get(key); ok {
		if s, ok := v.(string); ok {
			return s, true
		}
	}
	return "", false
}

// SetCtxValue 配置令牌加密方式
func SetCtxValue(c Context, key string, value interface{}) {
	c.Set(key, value)
}
