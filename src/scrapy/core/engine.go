package core

import (
	"container/list"
	"go/types"
	"reflect"
	"scrapy/core/downloader"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/spiders"
)

type EngineSlot struct {
	closing       bool
	inProgress    map[request.Request]bool
	startRequests *list.List
	closeIfIdle   bool
	scheduler     *Scheduler
}

func NewSlot(startRequests *list.List, closeIfIdle bool, scheduler *Scheduler) *EngineSlot {
	return &EngineSlot{startRequests: startRequests, closeIfIdle: closeIfIdle, scheduler: scheduler}
}

func (es *EngineSlot) AddRequest(req *request.Request) {
	es.inProgress[req] = true
}

func (es *EngineSlot) RemoveRequest(req *request.Request) {
	delete(es.inProgress, req)
}

func (es *EngineSlot) Close() {

}

type ExecutionEngine struct {
	crawler    *crawler.Crawler
	scheduler  *Scheduler
	slot       *EngineSlot
	spider     *spiders.Spider
	running    bool
	paused     bool
	downloader *downloader.Downloader
	scraper    *Scraper
}

func NewExecutionEngine(crawler *crawler.Crawler) *ExecutionEngine {
	return &ExecutionEngine{crawler: crawler, scraper: &Scraper{}}
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

func (ee *ExecutionEngine) Crawl(req *request.Request, spider *spiders.Spider) {

}

func (ee *ExecutionEngine) Schedule(req request.Request, spider spiders.Spider) {

}

func (ee *ExecutionEngine) Download(req request.Request, spider spiders.Spider) *response.Response {
	return nil
}

func (ee *ExecutionEngine) OpenSpider(spider *spiders.Spider, startRequests *list.List, closeIfIdle bool) {
	scheduler := NewScheduler(ee.crawler)
	startRequests = ee.scraper.spidermw.ProcessStartRequests(startRequests, spider)
	slot := NewSlot(startRequests, closeIfIdle, scheduler)
	ee.slot = slot
	ee.spider = spider
	go scheduler.Open(spider)
	go ee.scraper.OpenSpider(spider)
}

func (ee *ExecutionEngine) CloseSpider(spider *spiders.Spider, reason string) (bool, error) {
	if reason == "" {
		reason = "cancelled"
	}
	return true, nil
}

func (ee *ExecutionEngine) nextRequest(spider *spiders.Spider) *request.Request {
	return nil
}

func (ee *ExecutionEngine) nextRequestFromScheduler(spider *spiders.Spider) chan interface{} {
	slot := ee.slot
	req := slot.scheduler.NextRequest()
	if req == nil {
		return nil
	}
	result := ee.download(req, spider)

	ee.handleDownloaderOutput(nil, req, spider)
	return result
}

func (ee *ExecutionEngine) download(req *request.Request, spider *spiders.Spider) chan interface{} {
	slot := ee.slot
	slot.AddRequest(req)

	return ee.downloader.Fetch(req, spider)
}

func (ee *ExecutionEngine) handleDownloaderOutput(resp *response.Response, req *request.Request, spider *spiders.Spider) interface{} {
	if reflect.TypeOf(resp).Name() == "Request" {
		ee.Crawl(req, spider)
		return nil
	}
	ee.scraper.EnqueueScrape(resp, req, spider)
	return nil
}
