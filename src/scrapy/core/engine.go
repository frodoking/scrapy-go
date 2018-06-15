package core

import (
	"scrapy/core/downloader"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/spiders"
)

type EngineSlot struct {
	closing    bool
	inprogress map[request.Request]bool
}

func NewSlot() {

}

func (es *EngineSlot) AddRequest(req request.Request) {
	es.inprogress[req] = true
}

func (es *EngineSlot) RemoveRequest(req request.Request) {
	delete(es.inprogress, req)
}

func (es *EngineSlot) Close() {

}

type ExecutionEngine struct {
	crawler    *crawler.Crawler
	scheduler  *Scheduler
	downloader *downloader.Downloader
}

func NewExecutionEngine(crawler *crawler.Crawler) *ExecutionEngine {
	return &ExecutionEngine{crawler: crawler}
}

func (ee *ExecutionEngine) Start() {

}

func (ee *ExecutionEngine) Stop() {

}

func (ee *ExecutionEngine) Close() {

}

func (ee *ExecutionEngine) Pause() {

}

func (ee *ExecutionEngine) UnPause() {

}

func (ee *ExecutionEngine) SpiderIsIdle(spider spiders.Spider) bool {
	return true
}

func (ee *ExecutionEngine) Crawl(re request.Request, spider spiders.Spider) {

}

func (ee *ExecutionEngine) Schedule(re request.Request, spider spiders.Spider) {

}

func (ee *ExecutionEngine) Download(re request.Request, spider spiders.Spider) *response.Response {
	return nil
}

func (ee *ExecutionEngine) OpenSpider(spider *spiders.Spider, startRequests []*request.Request, closeIfIdle bool) *response.Response {

	return nil
}

func (ee *ExecutionEngine) CloseSpider(spider spiders.Spider, reason string) (bool, error) {
	if reason == "" {
		reason = "cancelled"
	}
	return true, nil
}
