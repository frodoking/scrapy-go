package response

import "scrapy/http"
import "scrapy/http/request"

type Response struct {
	url     string
	headers *http.Headers
	status  int
	Body    []byte
	Request *request.Request
	flags   []string
}

func NewResponse(url string) *Response {
	response := &Response{}
	response.url = url
	response.status = 200
	response.headers = http.NewHeaders(make(map[string]interface{}))
	response.Body = make([]byte, 0)
	response.flags = make([]string, 0)
	return response
}

func (r *Response) Meta() map[string]interface{} {
	return r.Request.Meta
}



func (r *Response) Copy() {
	r.Replace()
}

func (r *Response) Replace() {

}

func (r *Response) UrlJoin(url string) {

}

func (r *Response) Text() string {
	return ""
}

func (r *Response) CSS() string {
	return ""
}

func (r *Response) Xpath(args []string, kwargs []string) interface{} {
	return nil
}

func (r *Response) Follow() *request.Request {
	return nil
}
