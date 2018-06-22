package downloader

import (
	"scrapy/http"
	"scrapy/http/request"
	"scrapy/http/response"
)

type ScrapyHTTPPageGetter struct {
}

func (shpg *ScrapyHTTPPageGetter) ConnectionMade() {

}

func (shpg *ScrapyHTTPPageGetter) LineReceived(line string) {

}

func (shpg *ScrapyHTTPPageGetter) HandleHeader(key string, value interface{}) {

}

func (shpg *ScrapyHTTPPageGetter) HandleStatus(version string, status int, message string) {

}

func (shpg *ScrapyHTTPPageGetter) HandleEndHeaders(line string) {

}

func (shpg *ScrapyHTTPPageGetter) ConnectionLost(reason string) {

}

func (shpg *ScrapyHTTPPageGetter) HandleResponse(response response.Response) {

}

func (shpg *ScrapyHTTPPageGetter) Timeout() {

}

type ScrapyHTTPClientFactory struct {
	headersTime     int64
	startTime       int64
	responseHeaders map[string]interface{}
}

func (shcf *ScrapyHTTPClientFactory) buildResponse(body []byte, request *request.Request) response.Response {
	request.Meta["download_latency"] = shcf.headersTime - shcf.startTime
	headers := http.NewHeaders(shcf.responseHeaders)
	resp := response.ScrapyResponseTypes.FromHeaders(headers)
	switch resp.(type) {
	case response.TEXTResponse:
		break
	case response.HTMLResponse:
		break
	case response.XMLResponse:
		break
	}

	return resp
}
