package middleware

import (
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware"
	"scrapy/spiders"
)

type Middleware interface {
	ProcessRequest(request *request.Request, spider spiders.Spider) chan interface{}
	ProcessResponse(request *request.Request, response response.Response, spider spiders.Spider)
	ProcessException(request *request.Request, exception interface{}, spider spiders.Spider)
}

type DownloaderMiddlewareManager struct {
	*middleware.MiddlewareManager
}

func (dmm *DownloaderMiddlewareManager) Download(request *request.Request, spider spiders.Spider) (response.Response, error) {
	var resp response.Response
	var err error
	for e := dmm.Middlewares.Front(); e != nil; e = e.Next() {
		mw := e.Value.(Middleware)
		responseChan := mw.ProcessRequest(request, spider)
		if responseChan != nil {
			select {
			case resultTmp := <-responseChan:
				if resultTmp != nil {
					switch resultTmp.(type) {
					case error:
						mw.ProcessException(request, resultTmp, spider)
						err = resultTmp.(error)
						break
					case response.Response:
						resp = resultTmp.(response.Response)
					}
				}
			}
		}
	}

	if resp != nil {
		for e := dmm.Middlewares.Back(); e != nil; e = e.Prev() {
			mw := e.Value.(Middleware)
			mw.ProcessResponse(request, resp, spider)
		}
	}

	return resp, err
}
