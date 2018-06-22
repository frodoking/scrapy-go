package middleware

import (
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type RefererMiddleware struct {
	*middleware.MiddlewareManager
}

func (cm *RefererMiddleware) ProcessSpiderInput(response response.Response, spider *spiders.Spider) {

}
func (cm *RefererMiddleware) ProcessSpiderOutput(response response.Response, result interface{}, spider *spiders.Spider) {

}
func (cm *RefererMiddleware) ProcessSpiderException(response response.Response, exception interface{}, spider *spiders.Spider) {

}
