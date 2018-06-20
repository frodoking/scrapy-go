package handlers

import (
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloadHandlers struct {
	schemes       map[string]string
	handlers      map[string]DownloadHandlers
	notConfigured map[string]string
}

func NewDownloadHandlers() *DownloadHandlers {
	return &DownloadHandlers{}
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request, spider *spiders.Spider) chan interface{} {
	result := make(chan interface{}, 1)
	return result
}
