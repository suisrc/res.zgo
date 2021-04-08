package res

import (
	"fmt"
	"strings"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(Context) bool

// HandlerFunc ...
type HandlerFunc func(Context)

// SkipHandler 统一处理跳过函数
func SkipHandler(c Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// EmptyMiddleware 不执行业务处理的中间件
func EmptyMiddleware() HandlerFunc {
	return func(c Context) {
		c.Next() // Pass, 跳过
	}
}

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c Context) bool {
		path := c.GetRequest().URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c Context) bool {
		path := c.GetRequest().URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果不包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c Context) bool {
		creq := c.GetRequest()
		path := JoinRouter(creq.Method, creq.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinPath 拼接路由
func JoinPath(paths ...string) string {
	path := "/"
	for _, p := range paths {
		if lr := len(p); lr > 0 {
			ll := len(path)
			if p[0] == '/' && path[ll-1] == '/' {
				path += p[1:]
			} else if p[0] != '/' && path[ll-1] != '/' {
				path += "/" + p
			} else {
				path += p
			}
		}
	}
	return path
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
