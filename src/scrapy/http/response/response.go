package response

import "scrapy/http"
import "scrapy/http/request"

type Response interface {
	Url() string
	Body() []byte
	Meta() map[string]interface{}
	Copy()
	Replace()
	UrlJoin(url string)
	Text() string
	CSS() string
	XPath(args []string, kwargs []string) interface{}
	Follow() *request.Request
}

type BaseResponse struct {
	url     string
	headers *http.Headers
	status  int
	body    []byte
	Request *request.Request
	flags   []string
}

func NewResponse(url string) *BaseResponse {
	response := &BaseResponse{}
	response.url = url
	response.status = 200
	response.headers = http.NewHeaders(make(map[string]interface{}))
	response.body = make([]byte, 0)
	response.flags = make([]string, 0)
	return response
}

func (r *BaseResponse) Url() string {
	return r.url
}

func (r *BaseResponse) Meta() map[string]interface{} {
	return r.Request.Meta
}

func (r *BaseResponse) Body() []byte {
	return r.body
}

func (r *BaseResponse) Copy() {
	r.Replace()
}

func (r *BaseResponse) Replace() {

}

func (r *BaseResponse) UrlJoin(url string) {

}

func (r *BaseResponse) Text() string {
	return ""
}

func (r *BaseResponse) CSS() string {
	return ""
}

func (r *BaseResponse) XPath(args []string, kwargs []string) interface{} {
	return nil
}

func (r *BaseResponse) Follow() *request.Request {
	return nil
}
