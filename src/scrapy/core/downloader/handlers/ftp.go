package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type FtpDownloadHandler struct {
}

func (fdh *FtpDownloadHandler) DownloadRequest(request *request.Request, spider spiders.Spider) chan interface{} {
	return nil
}
