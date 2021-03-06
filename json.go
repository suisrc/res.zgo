package res

import (
	//"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

// 定义JSON操作
var (
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal       = json.Marshal
	JSONUnmarshal     = json.Unmarshal
	JSONMarshalIndent = json.MarshalIndent
	JSONNewDecoder    = json.NewDecoder
	JSONNewEncoder    = json.NewEncoder
)

// 定义Type
var (
	// Response of request
	ResponseTypeJSON = "application/json; charset=utf-8"
	ResponseTypeTEXT = "text/plain; charset=utf-8"
)

// JSONMarshalToString JSON编码为字符串
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}
