package request

import "scrapy/http"

type Request struct {
	url        string
	encoding   string
	method     string
	headers    *http.Headers
	body       string
	cookies    *http.Cookies
	meta       string
	priority   int
	dontFilter bool
	flags      []string
}

func (request *Request) GetMeta() string {
	return request.meta
}
