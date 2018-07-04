package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type DataURIDownloadHandler struct {
	DefaultDownloadHandler
}

func (fdh *DataURIDownloadHandler) DownloadRequest(request *request.Request, spider spiders.Spider) chan interface{} {
	return nil
}
