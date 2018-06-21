package middleware

import (
	"container/list"
	"scrapy/spiders"
)

type MiddlewareManager struct {
	middlewares *list.List
	methods     map[string]string
}

func (m *MiddlewareManager) OpenSpider(spider *spiders.Spider) chan struct{} {
	return nil
}

func (m *MiddlewareManager) CloseSpider(spider *spiders.Spider) chan struct{} {
	return nil
}
