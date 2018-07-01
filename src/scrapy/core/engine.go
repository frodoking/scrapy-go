package core

import (
	"container/list"
	"reflect"
	"scrapy/core/downloader"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/settings"
	"scrapy/spiders"
	"time"
)

type CallLaterOnce struct {
	myFunc func(spider spiders.Spider)
	spider spiders.Spider
	timer  *time.Timer
}

func NewCallLaterOnce(myFunc func(spider spiders.Spider), spider spiders.Spider) *CallLaterOnce {
	return &CallLaterOnce{myFunc: myFunc, spider: spider}
}

func (clo *CallLaterOnce) schedule(delay uint) {
	clo.timer = time.AfterFunc(time.Duration(delay)*time.Second, func() {
		clo.myFunc(clo.spider)
	})
}

func (clo *CallLaterOnce) cancel() {
	clo.timer.Stop()
}

type EngineSlot struct {
	closing       bool
	inProgress    map[*request.Request]bool
	startRequests *list.List
	closeIfIdle   bool
	nextCall      *CallLaterOnce
	scheduler     *Scheduler
	heartbeat     *time.Ticker
}

func NewSlot(startRequests *list.List, closeIfIdle bool, nextCall *CallLaterOnce, scheduler *Scheduler) *EngineSlot {
	ticker := time.NewTicker(time.Second * 5)
	return &EngineSlot{startRequests: startRequests, closeIfIdle: closeIfIdle, nextCall: nextCall, scheduler: scheduler, heartbeat: ticker}
}

func (es *EngineSlot) AddRequest(req *request.Request) {
	es.inProgress[req] = true
}

func (es *EngineSlot) RemoveRequest(req *request.Request) {
	delete(es.inProgress, req)
}

func (es *EngineSlot) Close() {

}

func (es *EngineSlot) StartHeartbeat(delay uint) {
	go func() {
		for range es.heartbeat.C {
			es.nextCall.schedule(delay)
		}
	}()
}

type ExecutionEngine struct {
	settings   *settings.Settings
	scheduler  *Scheduler
	slot       *EngineSlot
	spider     spiders.Spider
	running    bool
	paused     bool
	downloader *downloader.Downloader
	scraper    *Scraper
}

func NewExecutionEngine(settings *settings.Settings) *ExecutionEngine {
	return &ExecutionEngine{settings: settings, scraper: NewScraper(settings)}
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

func (ee *ExecutionEngine) Crawl(req *request.Request, spider spiders.Spider) {
	ee.Schedule(req, spider)
	ee.slot.nextCall.schedule(0)
}

func (ee *ExecutionEngine) Schedule(req *request.Request, spider spiders.Spider) {
	if !ee.slot.scheduler.EnqueueRequest(req) {
		println("xxxxx")
	}
}

func (ee *ExecutionEngine) Download(req request.Request, spider spiders.Spider) response.Response {
	return nil
}

func (ee *ExecutionEngine) OpenSpider(spider spiders.Spider, startRequests *list.List, closeIfIdle bool) {
	nextCall := NewCallLaterOnce(ee.nextRequest, spider)
	scheduler := NewScheduler(ee.settings)
	startRequests = ee.scraper.spidermw.ProcessStartRequests(startRequests, spider)
	slot := NewSlot(startRequests, closeIfIdle, nextCall, scheduler)
	ee.slot = slot
	ee.spider = spider
	go scheduler.Open(spider)
	go ee.scraper.OpenSpider(spider)
	slot.nextCall.schedule(0)
	slot.StartHeartbeat(5)
}

func (ee *ExecutionEngine) CloseSpider(spider spiders.Spider, reason string) (bool, error) {
	if reason == "" {
		reason = "cancelled"
	}
	return true, nil
}

func (ee *ExecutionEngine) nextRequest(spider spiders.Spider) {
	slot := ee.slot
	if slot == nil {
		return
	}

	if ee.paused {
		return
	}

	for !ee.needsBackout(spider) {
		if ee.nextRequestFromScheduler(spider) == nil {
			break
		}
	}

	if slot.startRequests.Len() > 0 && !ee.needsBackout(spider) {
		req := slot.startRequests.Front().Value.(*request.Request)
		ee.Crawl(req, spider)
	}

	if ee.SpiderIsIdle(spider) && slot.closeIfIdle {
		ee.SpiderIsIdle(spider)
	}
}

func (ee *ExecutionEngine) needsBackout(spider spiders.Spider) bool {
	slot := ee.slot
	return ee.running || slot.closing || ee.downloader.NeedsBackout() || ee.scraper.slot.NeedsBackout()
}

func (ee *ExecutionEngine) nextRequestFromScheduler(spider spiders.Spider) chan interface{} {
	slot := ee.slot
	req := slot.scheduler.NextRequest()
	if req == nil {
		return nil
	}
	result := ee.download(req, spider)

	ee.handleDownloaderOutput((<-result).(response.Response), req, spider)
	return result
}

func (ee *ExecutionEngine) download(req *request.Request, spider spiders.Spider) chan interface{} {
	slot := ee.slot
	slot.AddRequest(req)

	return ee.downloader.Fetch(req, spider)
}

func (ee *ExecutionEngine) handleDownloaderOutput(resp response.Response, req *request.Request, spider spiders.Spider) interface{} {
	if reflect.TypeOf(resp).Name() == "Request" {
		ee.Crawl(req, spider)
		return nil
	}
	ee.scraper.EnqueueScrape(resp, req, spider)
	return nil
}
