package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type HttpDownloadHandler struct {
}

func (fdh *HttpDownloadHandler) DownloadRequest(request *request.Request, spider *spiders.Spider) chan interface{} {
	return nil
}
