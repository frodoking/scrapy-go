package request

import "scrapy/http"

type Request struct {
	url        string
	encoding   string
	method     string
	headers    *http.Headers
	body       string
	cookies    *http.Cookies
	Meta       map[string]interface{}
	priority   int
	dontFilter bool
	flags      []string
	callback   interface{}
	errorBack  interface{}
}

func NewRequest(url string, encoding string) *Request {
	request := &Request{}
	request.url = url
	request.callback = nil
	request.method = "GET"
	request.encoding = encoding
	request.headers = http.NewHeaders(make(map[string]interface{}))
	request.body = ""
	request.cookies = nil
	request.meta = make(map[string]interface{})
	request.priority = 0
	request.dontFilter = false
	request.errorBack = nil
	request.flags = make([]string, 0)
	return request
}

func (request *Request) GetMeta() map[string]interface{} {
	return request.meta
}
