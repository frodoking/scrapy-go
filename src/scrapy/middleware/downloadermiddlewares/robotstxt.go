package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type RobotsTxtMiddleware struct {
	*middleware.MiddlewareManager
}

func (cm *RobotsTxtMiddleware) ProcessRequest(request *request.Request, spider *spiders.Spider) {

}
func (cm *RobotsTxtMiddleware) ProcessResponse(request *request.Request, response response.Response, spider *spiders.Spider) {

}
func (cm *RobotsTxtMiddleware) ProcessException(request *request.Request, exception interface{}, spider *spiders.Spider) {

}
