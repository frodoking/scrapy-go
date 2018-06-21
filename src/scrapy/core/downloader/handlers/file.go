package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type FileDownloadHandler struct {
}

func (fdh *FileDownloadHandler) DownloadRequest(request *request.Request, spider *spiders.Spider) chan interface{} {
	return nil
}
