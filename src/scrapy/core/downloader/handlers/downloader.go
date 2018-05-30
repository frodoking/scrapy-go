package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloadHandlers struct {
	crawler       interface{}
	schemes       map[string]string
	handlers      map[string]DownloadHandlers
	notConfigured map[string]string
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request, spider *spiders.Spider) interface{} {
	return nil
}
