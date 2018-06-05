package middleware

import "scrapy/spiders"

type MiddlewareManager struct {
	middlewares []MiddlewareManager
	methods map[string]string
}

func (m *MiddlewareManager) OpenSpider(spider *spiders.Spider) chan struct{} {
	return nil
}

func (m *MiddlewareManager) CloseSpider(spider *spiders.Spider) chan struct{} {
	return nil
}