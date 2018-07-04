package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloadHandler interface {
	// return real response to chan
	DownloadRequest(request *request.Request, spider spiders.Spider) chan interface{}
	close()
}

type DefaultDownloadHandler struct {
}

func (dh *DefaultDownloadHandler) DownloadRequest(request *request.Request, spider spiders.Spider) chan interface{} {
	return nil
}

func (dh *DefaultDownloadHandler) close() {

}
