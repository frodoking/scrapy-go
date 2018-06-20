package handlers

import (
	"scrapy/http/request"
)

type DownloadHandlers struct {
	schemes       map[string]string
	handlers      map[string]DownloadHandlers
	notConfigured map[string]string
}

func NewDownloadHandlers() *DownloadHandlers {
	return &DownloadHandlers{}
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request) interface{} {
	return nil
}
