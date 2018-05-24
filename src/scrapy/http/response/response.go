package src

import "scrapy/http"

type Response struct {
	url     string
	headers *headers.Headers
	status  int
	body    []byte
	request *request.Request
	flags   []string
}

func (response *Response) Copy() {

}

func (response *Response) Replace() {

}

func (response *Response) Xpath() interface{} {
	return nil
}

func (response *Response) Follow() *Request {
	return nil
}
