package baidu_push

type Request struct {
	Params map[string]interface{}
	Url    string
}

func NewRequest(path string) *Request {
	req := &Request{
		Url: BaseUrl + path,
	}
	req.Params = make(map[string]interface{})
	return req
}
