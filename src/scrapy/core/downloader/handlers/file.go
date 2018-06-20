package handlers

import (
	"scrapy/http/request"
)

type FileDownloadHandler struct {
	*DownloadHandlers
}

func (fdh *FileDownloadHandler) DownloadRequest(request *request.Request) interface{} {
	return nil
}
