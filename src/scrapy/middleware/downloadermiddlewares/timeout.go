package middleware

import (
	"scrapy/common"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/spiders"
)

type DownloadTimeoutMiddleware struct {
	*DownloaderMiddleware
	timeout int
}

func NewDownloadTimeoutMiddleware() *DownloadTimeoutMiddleware {
	middleware := &DownloadTimeoutMiddleware{}
	listener := common.ScrapySignal.Connect(common.SpiderOpened)
	go func() {
		for {
			select {
			case event := <-listener:
				if event != nil {
					spider := event.(spiders.Spider)
					middleware.timeout = spider.GetAttrFromSettings("download_timeout").(int)
					return
				}
			}
		}
	}()

	return middleware
}

func (cm *DownloadTimeoutMiddleware) ProcessRequest(request *request.Request, spider spiders.Spider) (response.Response, error) {
	if cm.timeout > 0 {
		request.Meta["download_timeout"] = cm.timeout
	}
	return nil, nil
}
