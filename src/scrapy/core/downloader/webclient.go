package downloader

import "scrapy/http/response"

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

func (shpg *ScrapyHTTPPageGetter) HandleResponse(response *response.Response) {

}

func (shpg *ScrapyHTTPPageGetter) Timeout() {

}

type ScrapyHTTPClientFactory struct {
}
