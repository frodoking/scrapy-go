package middleware

import (
	"container/list"
	"scrapy/middleware"
	"scrapy/spiders"
)

type SpiderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (smm *SpiderMiddlewareManager) ProcessStartRequests(reqs *list.List, spider *spiders.Spider) *list.List {
	return reqs
}
