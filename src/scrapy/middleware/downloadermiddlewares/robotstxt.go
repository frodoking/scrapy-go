package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/spiders"
)

type RobotsTxtMiddleware struct {
	*DownloaderMiddleware
}

func (cm *RobotsTxtMiddleware) ProcessRequest(request *request.Request, spider spiders.Spider) chan interface{} {
	return nil
}
func (cm *RobotsTxtMiddleware) ProcessResponse(request *request.Request, response response.Response, spider spiders.Spider) {

}
func (cm *RobotsTxtMiddleware) ProcessException(request *request.Request, exception interface{}, spider spiders.Spider) {

}
