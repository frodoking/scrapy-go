package downloader

import (
	"scrapy/middleware"
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (m *DownloaderMiddlewareManager) download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) {

}
