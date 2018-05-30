package downloader

import (
	"scrapy"
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloaderMiddlewareManager struct {
	*scrapy.MiddlewareManager
}

func (m *DownloaderMiddlewareManager) download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) {

}
