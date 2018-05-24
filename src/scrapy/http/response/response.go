package response

import "scrapy/http"
import "scrapy/http/request"

type Response struct {
	url     string
	headers *http.Headers
	status  int
	body    []byte
	request *request.Request
	flags   []string
}

func NewResponse(url string) *Response {
	response := &Response{}
	response.url = url
	response.status = 200
	response.headers = http.NewHeaders(make(map[string]interface{}))
	response.body = make([]byte, 0)
	response.flags = make([]string, 0)
	return response
}

func (r *Response) Copy() {

}

func (r *Response) Replace() {

}

func (r *Response) Xpath() interface{} {
	return nil
}

func (r *Response) Follow() *request.Request {
	return nil
}
