package libgrpc

var (
	SUCCESS   = &Response{Code: "RC00000", Info: "Success"}
	SYS_ERROR = &Response{Code: "RC90000", Info: "System error"}
)

func (resp *Response) SetCode(srcResp *Response, message... string) {
	if resp == nil {
		resp = &Response{}
	}
	resp.Code = srcResp.Code
	if len(message) == 0 {
		resp.Info = srcResp.Info
	} else {
		resp.Info = message[0]
	}
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
