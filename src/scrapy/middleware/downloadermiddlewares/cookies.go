package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type CookiesMiddleware struct {
	*middleware.MiddlewareManager
}

func (cm *CookiesMiddleware) ProcessRequest(request *request.Request, spider *spiders.Spider) chan interface{} {
	return nil
}
func (cm *CookiesMiddleware) ProcessResponse(request *request.Request, response response.Response, spider *spiders.Spider) {

}
func (cm *CookiesMiddleware) ProcessException(request *request.Request, exception interface{}, spider *spiders.Spider) {

}
