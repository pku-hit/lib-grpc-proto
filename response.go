package libgrpc

var (
	SUCCESS   = &Response{Code: "RC00000", Info: "Success"}
	SYS_ERROR = &Response{Code: "RC90000", Info: "System error"}
)

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
