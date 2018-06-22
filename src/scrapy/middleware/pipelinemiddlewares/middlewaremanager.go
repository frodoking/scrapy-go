package middleware

import (
	"container/list"
	"scrapy/middleware"
	"scrapy/spiders"
)

type Middleware interface {
}

type ItemPipelineManager struct {
	*middleware.MiddlewareManager
}

func (ipm *ItemPipelineManager) ProcessStartRequests(reqs *list.List, spider *spiders.Spider) *list.List {
	return reqs
}

func (ipm *ItemPipelineManager) ProcessItem(reqs *list.List, spider *spiders.Spider) *list.List {
	return reqs
}
