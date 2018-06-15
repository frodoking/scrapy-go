package downloader

import (
	"scrapy/http/request"
	"scrapy/middleware"
	"scrapy/spiders"
)

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (m *DownloaderMiddlewareManager) download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) {

}
