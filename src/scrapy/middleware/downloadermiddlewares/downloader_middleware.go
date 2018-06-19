package middleware

import (
	"scrapy/http/request"
	"scrapy/middleware"
	"scrapy/spiders"
)

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (m *DownloaderMiddlewareManager) Download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) {

}
