package api

import (
	"encoding/json"
)

var JsonContentType = []string{"application/json; charset=utf-8"}

/*响应结果*/
const (
	SUCCESS               = 0
	UNKNOWN               = -1
	READING_REQUEST_ERROR = 601 // 读取request失败
	PARSING_REQUEST_ERROR = 602 // 解析request失败
	EXECUTE_SERVICE_ERROR = 603 // 执行service失败
)

type Response struct {
	Code int         `json:"code" bson:"code" yaml:"code"`                               // 响应代码
	Msg  string      `json:"msg,omitempty" bson:"msg,omitempty" yaml:"msg,omitempty"`    // 响应消息
	Data interface{} `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"` // 响应数据
	Tag  string      `json:"tag,omitempty" bson:"tag,omitempty" yaml:"tag,omitempty"`    // 响应标签
}

func (rsp *Response) Error() string {
	bs, _ := json.Marshal(rsp)
	return string(bs)
}

func SuccessResponse(data interface{}, tag ...string) *Response {
	rsp := &Response{
		Data: data,
	}
	if len(tag) > 0 {
		rsp.Tag = tag[0]
	}
	return rsp
}

func FailureResponse(code int, msg string, tag ...string) *Response {
	rsp := &Response{
		Code: code,
		Msg:  msg,
	}
	if len(tag) > 0 {
		rsp.Tag = tag[0]
	}
	return rsp
}

func Error(err error, tag ...string) *Response {
	rsp := &Response{
		Code: UNKNOWN,
		Msg:  err.Error(),
	}
	if len(tag) > 0 {
		rsp.Tag = tag[0]
	}
	return rsp
}
