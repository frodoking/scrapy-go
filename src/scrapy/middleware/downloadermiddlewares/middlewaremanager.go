package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type Middleware interface {
	ProcessRequest(request *request.Request, spider spiders.Spider) (response.Response, error)
	ProcessResponse(request *request.Request, response response.Response, spider spiders.Spider)
	ProcessException(request *request.Request, exception interface{}, spider spiders.Spider)
}

type DownloaderMiddleware struct {
}

func (cm *DownloaderMiddleware) ProcessRequest(request *request.Request, spider spiders.Spider) (response.Response, error) {
	return nil, nil
}

func (cm *DownloaderMiddleware) ProcessResponse(request *request.Request, response response.Response, spider spiders.Spider) {

}
func (cm *DownloaderMiddleware) ProcessException(request *request.Request, exception interface{}, spider spiders.Spider) {

}

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func NewDownloaderMiddlewareManager() *DownloaderMiddlewareManager {
	instance := &DownloaderMiddlewareManager{}
	instance.AddMiddleware(&CookiesMiddleware{})
	instance.AddMiddleware(&DownloadTimeoutMiddleware{})
	instance.AddMiddleware(&RobotsTxtMiddleware{})
	return instance
}

func (dmm *DownloaderMiddlewareManager) Download(request *request.Request, spider spiders.Spider) (response.Response, error) {
	var respResult response.Response
	var errResult error

	for e := dmm.Middlewares.Front(); e != nil; e = e.Next() {
		mw := e.Value.(Middleware)
		resp, err := mw.ProcessRequest(request, spider)
		if err == nil {
			if resp != nil {
				respResult = resp
			}
		} else {
			mw.ProcessException(request, resp, spider)
			errResult = err
		}
	}

	if respResult != nil {
		for e := dmm.Middlewares.Back(); e != nil; e = e.Prev() {
			mw := e.Value.(Middleware)
			mw.ProcessResponse(request, respResult, spider)
		}
	}

	return respResult, errResult
}
