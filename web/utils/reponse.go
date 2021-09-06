package utils

// Response 响应
type Response struct {
	Errno  errorCode   // 错误码
	Errmsg string      // 错误消息
	Data   interface{} // 数据
}


//
// OK 响应成功
//  @Description:
//  @param data
//  @return *Response
//
func OK(data interface{}) *Response {
	return &Response{
		Errno:  RECODE_OK,
		Errmsg: RecodeText(RECODE_OK),
		Data:   data,
	}
}

//
// OKMsg
//  @Description: 响应成功 自定义提示
//  @param msg
//  @param data
//  @return *Response
//
func OKMsg(msg string, data interface{}) *Response {
	return &Response{
		Errno:  RECODE_OK,
		Errmsg: msg,
		Data:   data,
	}
}

//
// Fail
//  @Description: 响应错误
//  @param code 错误码
//  @return *Response 响应实体
//
func Fail(code errorCode) *Response {
	return &Response{
		Errno:  code,
		Errmsg: RecodeText(code),
		Data:   nil,
	}
}

//
// FailMsg
//  @Description: 响应错误：自定义消息errMsg
//  @param code
//  @param errMsg
//  @return *Response
//
func FailMsg(code errorCode, errMsg string) *Response {
	return &Response{
		Errno:  code,
		Errmsg: errMsg,
		Data:   nil,
	}
}
//
// FailData
//  @Description:  响应错误 携带数据
//  @param code
//  @param data
//  @return *Response
//
func FailData(code errorCode, data interface{}) *Response {
	return &Response{
		Errno:  code,
		Errmsg: RecodeText(code),
		Data:   data,
	}
}
