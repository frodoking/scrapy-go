package handlers

import (
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/spiders"
)

type DownloadHandlers struct {
	crawler       *crawler.Crawler
	schemes       map[string]string
	handlers      map[string]DownloadHandlers
	notConfigured map[string]string
}

func NewDownloadHandlers(crawler *crawler.Crawler) *DownloadHandlers {
	return &DownloadHandlers{crawler: crawler}
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request, spider *spiders.Spider) interface{} {
	return nil
}
