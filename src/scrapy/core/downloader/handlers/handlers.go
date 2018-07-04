package handlers

import (
	"errors"
	"fmt"
	"net/url"
	"scrapy/common"
	"scrapy/http/request"
	"scrapy/spiders"
)

var Classes = map[string]string{
	"data":  "DataURIDownloadHandler",
	"file":  "FileDownloadHandler",
	"http":  "HTTPDownloadHandler",
	"https": "HTTPDownloadHandler",
	"s3":    "S3DownloadHandler",
	"ftp":   "FTPDownloadHandler",
}

type DownloadHandlers struct {
	schemes       map[string]string
	handlers      map[string]DownloadHandler
	notConfigured map[string]string
}

func NewDownloadHandlers() *DownloadHandlers {
	schemes := make(map[string]string)
	handlers := make(map[string]DownloadHandler)
	notConfigured := make(map[string]string)

	for scheme, class := range Classes {
		schemes[scheme] = scheme
		switch class {
		case "http":
		case "https":
			handlers[scheme] = &HttpDownloadHandler{}
			break
		case "file":
			handlers[scheme] = &FileDownloadHandler{}
			break
		case "ftp":
			handlers[scheme] = &FtpDownloadHandler{}
			break
		default:
			break
		}
	}

	instance := &DownloadHandlers{schemes, handlers, notConfigured}

	logger := common.WithLogger("handlers")
	listener := common.ScrapySignal.Connect(common.EngineStopped)
	go func() {
		for {
			select {
			case event := <-listener:
				if event != nil {
					logger.Info("received : %s ", event.(string))
					instance.close()
					return
				}
			}
		}
	}()

	return instance
}

func (dh *DownloadHandlers) getHandler(scheme string) DownloadHandler {
	if value, ok := dh.handlers[scheme]; ok {
		return value
	}

	if _, ok := dh.notConfigured[scheme]; ok {
		return nil
	}

	if _, ok := dh.schemes[scheme]; !ok {
		return nil
	}

	dhcls := dh.handlers[scheme]

	var downloadHandler DownloadHandler = nil
	switch dhcls.(type) {
	case *HttpDownloadHandler:
		downloadHandler = &HttpDownloadHandler{}
		break
	case *FileDownloadHandler:
		downloadHandler = &FileDownloadHandler{}
		break
	case *FtpDownloadHandler:
		downloadHandler = &FtpDownloadHandler{}
		break
	default:
		dh.notConfigured[scheme] = fmt.Sprintf("Loading download handler for scheme:[%s] not found", scheme)
	}

	if downloadHandler != nil {
		dh.handlers[scheme] = downloadHandler
	}

	return dh
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request, spider spiders.Spider) chan interface{} {
	requestUrl, err := url.Parse(request.Url)
	if err != nil {
		panic(fmt.Sprintf("parse Url[%s] error", request.Url))
	}
	scheme := requestUrl.Scheme
	handler := dh.getHandler(scheme)
	if handler == nil {
		errors.New(fmt.Sprintf("Unsupported URL scheme %s: %s", scheme, handler))
	}

	return handler.DownloadRequest(request, spider)
}

func (dh *DownloadHandlers) close() {
	for _, handler := range dh.handlers {
		handler.close()
	}
}
