package res

import "fmt"

// Success 正常请求结构体
type Success struct {
	Success bool        `json:"success"`        // 请求成功, false
	Data    interface{} `json:"data,omitempty"` // 响应数据
	TraceID string      `json:"traceId"`        // 方便进行后端故障排除：唯一的请求ID
}

func (e *Success) Error() string {
	return "success"
}

// ErrorInfo 异常的请求结果体
type ErrorInfo struct {
	Success      bool        `json:"success"`        // 请求成功, false
	Data         interface{} `json:"data,omitempty"` // 响应数据
	ErrorCode    string      `json:"errorCode"`      // 错误代码
	ErrorMessage string      `json:"errorMessage"`   // 向用户显示消息
	ShowType     int         `json:"showType"`       // 错误显示类型：0静音； 1条消息警告； 2消息错误； 4通知； 9页
	TraceID      string      `json:"traceId"`        // 方便进行后端故障排除：唯一的请求ID
	//Status       int         `json:"-"`
}

func (e *ErrorInfo) Error() string {
	return fmt.Sprintf("[%s]%s", e.ErrorCode, e.ErrorMessage)
}

// ErrorRedirect 重定向
type ErrorRedirect struct {
	Status   int    // http.StatusSeeOther
	State    string // 状态, 用户还原缓存现场
	Location string
}

func (e *ErrorRedirect) Error() string {
	return "Redirect: " + e.Location
}

// ErrorNone 返回值已经被处理,无返回值
type ErrorNone struct {
}

func (e *ErrorNone) Error() string {
	return "none"
}
