package spiders

import (
	"container/list"
	"scrapy/http/request"
	"scrapy/http/response"
)

type Spider interface {
	StartRequests() *list.List

	MakeRequestsFromUrl(url string) *request.Request

	Parse(response response.Response)

	HandlesRequest(request *request.Request) interface{}
}
