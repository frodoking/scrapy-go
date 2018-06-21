package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type Middleware interface {
	ProcessRequest(request *request.Request, spider *spiders.Spider)
	ProcessResponse(request *request.Request, response *response.Response, spider *spiders.Spider)
	ProcessException(request *request.Request, exception interface{}, spider *spiders.Spider)
}

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (m *DownloaderMiddlewareManager) Download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) chan interface{} {
	result := make(chan interface{})

	go func() {
		result <- &response.Response{}
	}()

	return result
}
