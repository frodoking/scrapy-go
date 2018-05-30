package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type FileDownloadHandler struct {
	*DownloadHandlers
}

func (fdh *FileDownloadHandler) DownloadRequest(request *request.Request, spider *spiders.Spider) interface{} {
	return nil
}
