package middleware

import (
	"container/list"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type Middleware interface {
	ProcessSpiderInput(response response.Response, spider *spiders.Spider)
	ProcessSpiderOutput(response response.Response, result interface{}, spider *spiders.Spider)
	ProcessSpiderException(response response.Response, exception interface{}, spider *spiders.Spider)
}

type SpiderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (smm *SpiderMiddlewareManager) ProcessStartRequests(reqs *list.List, spider *spiders.Spider) *list.List {
	return reqs
}

func (smm *SpiderMiddlewareManager) ScrapeResponse(reqs *list.List, spider *spiders.Spider) *list.List {
	return reqs
}
