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

const PBX_PACK_PATH = "github.com/obase/go/pbx"

type Response struct {
	Code int         `json:"code" bson:"code"`                     // 响应代码
	Msg  string      `json:"msg,omitempty" bson:"msg,omitempty"`   // 响应消息
	Data interface{} `json:"data,omitempty" bson:"data,omitempty"` // 响应数据
	Tag  string      `json:"tag,omitempty" bson:"tag,omitempty"`   // 响应标签
}

func (rsp *Response) Error() string {
	bs, _ := json.Marshal(rsp)
	return string(bs)
}
