package request

import (
	"scrapy/http"
)

type Request struct {
	Url        string
	encoding   string
	method     string
	headers    *http.Headers
	body       string
	cookies    *http.Cookies
	Meta       map[string]interface{}
	priority   int
	DontFilter bool
	flags      []string
	callback   func(response interface{})
	errorBack  func(err error)
}

func NewRequest(url string, encoding string) *Request {
	request := &Request{}
	request.Url = url
	request.callback = nil
	request.method = "GET"
	request.encoding = encoding
	request.headers = http.NewHeaders(make(map[string]interface{}))
	request.body = ""
	request.cookies = nil
	request.Meta = make(map[string]interface{})
	request.priority = 0
	request.DontFilter = false
	request.errorBack = nil
	request.flags = make([]string, 0)
	return request
}
