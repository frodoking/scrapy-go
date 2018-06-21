package handlers

import (
	"errors"
	"reflect"
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

type DownloadHandler interface {
	// return real response to chan
	DownloadRequest(request *request.Request, spider *spiders.Spider) chan interface{}
}

type DownloadHandlers struct {
	schemes       map[string]string
	handlers      map[string]DownloadHandler
	notConfigured map[string]DownloadHandler
}

func NewDownloadHandlers() *DownloadHandlers {
	schemes := make(map[string]string)
	handlers := make(map[string]DownloadHandlers)
	notConfigured := make(map[string]DownloadHandlers)

	for scheme, class := range Classes {
		schemes[scheme] = scheme
		switch class {
		case "http":
		case "https":
			schemes[scheme] = &HttpDownloadHandler{}
			break
		case "file":
			schemes[scheme] = &FileDownloadHandler{}
			break
		case "ftp":
			schemes[scheme] = &FtpDownloadHandler{}
			break
		default:
			break
		}
	}

	return &DownloadHandlers{schemes, handlers, notConfigured}
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

	//fv := reflect.ValueOf(NewDownloadHandlers)
	//params := make([]reflect.Value,1)
	//params[0] = "test"
	// return fv.Call(params)

	return reflect.New(reflect.TypeOf(dhcls))
}

func (dh *DownloadHandlers) DownloadRequest(request *request.Request, spider *spiders.Spider) chan interface{} {
	scheme := ""
	handler := dh.getHandler(scheme)
	if handler == nil {
		errors.New("Unsupported URL scheme %s: %s")
	}

	return handler.DownloadRequest(request, spider)
}
