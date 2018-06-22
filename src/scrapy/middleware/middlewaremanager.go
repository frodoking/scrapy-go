package middleware

import (
	"container/list"
	"scrapy/spiders"
)

type MiddlewareManager struct {
	Middlewares *list.List
}

func (mwm *MiddlewareManager) addMiddleware(mw interface{}) {
	mwm.Middlewares.PushBack(mw)
}

func (m *MiddlewareManager) OpenSpider(spider *spiders.Spider) chan struct{} {
	return nil
}

func (m *MiddlewareManager) CloseSpider(spider *spiders.Spider) chan struct{} {
	return nil
}
