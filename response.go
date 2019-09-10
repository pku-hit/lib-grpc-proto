package libgrpc

import "fmt"

var (
	SUCCESS        = &Response{Code: "RC00000", Info: "Success"}
	LOGIN_ERROR    = &Response{Code: "RC10000", Info: "Logic error."}
	CALL_SVC_ERROR = &Response{Code: "RC70000", Info: "Calling service %s error %s."}
	DB_ERROR       = &Response{Code: "RC80000", Info: "Database invoke error: %s."}
	SYS_ERROR      = &Response{Code: "RC90000", Info: "System error: %s."}
)

func GenRespCode(srcResp *Response, message ...interface{}) (resp *Response) {
	resp = &Response{}
	resp.Code = srcResp.Code
	if len(message) == 0 {
		resp.Info = srcResp.Info
	} else {
		resp.Info = fmt.Sprintf(srcResp.Info, message)
	}
	return
}

func (resp *Response) IsSuccess() (result bool) {
	return resp.Is(SUCCESS)
}

func (resp *Response) Is(srcResp *Response) (result bool) {
	if resp == srcResp {
		result = true
	} else {
		result = resp != nil && srcResp != nil && resp.Code == srcResp.Code
	}
	return
}
