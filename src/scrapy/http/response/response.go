package response

import "scrapy/http"

type Response struct {
	url     string
	headers *headers.Headers
	status  int
	body    []byte
	request *request.Request
	flags   []string
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
