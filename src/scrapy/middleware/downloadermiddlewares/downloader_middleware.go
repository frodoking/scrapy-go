package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

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
