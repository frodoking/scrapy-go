package core

import (
	"reflect"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware/spidermiddlewares"
	"scrapy/spiders"
)

const MIN_RESPONSE_SIZE = 1024

type ResponseRequestDefer struct {
	Resp  *response.Response
	Req   *request.Request
	Defer chan bool
}

type ScraperSlot struct {
	maxActiveSize uint
	queue         []*ResponseRequestDefer
	active        map[*request.Request]bool
	activeSize    int
	itemProcSize  uint
	closing       bool
}

func NewScraperSlot(maxActiveSize uint) *ScraperSlot {
	if maxActiveSize == -1 {
		maxActiveSize = 5000000
	}
	return &ScraperSlot{maxActiveSize, make([]*ResponseRequestDefer, 10), make(map[*request.Request]bool), 0, 0, nil}
}

func (ss *ScraperSlot) AddResponseRequest(resp *response.Response, req *request.Request) chan bool {
	deferred := make(chan bool, 1)
	ss.queue = append(ss.queue, &ResponseRequestDefer{resp, req, deferred})

	if reflect.TypeOf(resp).Name() == "Response" {
		max := len(resp.Body)
		if max < MIN_RESPONSE_SIZE {
			max = MIN_RESPONSE_SIZE
		}
		ss.activeSize += max
	} else {
		ss.activeSize += MIN_RESPONSE_SIZE
	}
	return deferred
}

func (ss *ScraperSlot) NextResponseRequestDeferred() *ResponseRequestDefer {
	defered := ss.queue[0]
	ss.queue = ss.queue[1:]
	ss.active[defered.Req] = true
	return defered
}

func (ss *ScraperSlot) FinishResponse(resp *response.Response, req *request.Request) {
	delete(ss.active, req)
	if reflect.TypeOf(resp).Name() == "Response" {
		max := len(resp.Body)
		if max < MIN_RESPONSE_SIZE {
			max = MIN_RESPONSE_SIZE
		}
		ss.activeSize -= max
	} else {
		ss.activeSize -= MIN_RESPONSE_SIZE
	}
}

func (ss *ScraperSlot) IsIdle() bool {
	return len(ss.queue) == 0 && len(ss.active) == 0
}

type Scraper struct {
	slot            *ScraperSlot
	spidermw        *middleware.SpiderMiddlewareManager
	itemProc        interface{}
	concurrentItems interface{}
	crawler         *crawler.Crawler
}

func NewScraper(crawler *crawler.Crawler) {

}

func (s *Scraper) OpenSpider(spider *spiders.Spider) {
	s.slot = NewScraperSlot(-1)
	//itemProc := utils.LoadObject(s.crawler.Settings.Get("ITEM_PROCESSOR"))
}

func (s *Scraper) CloseSpider(spider *spiders.Spider) bool {
	return s.slot.closing
}

func (s *Scraper) IsIdle() bool {
	return s.slot == nil
}

func (s *Scraper) EnqueueScrape(resp *response.Response, req *request.Request, spider *spiders.Spider) chan bool {
	slot := s.slot
	defered := slot.AddResponseRequest(resp, req)
	s.scrapeNext(spider, slot)
	return defered
}

func (s *Scraper) scrapeNext(spider *spiders.Spider, slot *ScraperSlot) {

}

func (s *Scraper) HandleSpiderError(failure string, resp *response.Response, req *request.Request, spider *spiders.Spider) {

}

func (s *Scraper) HandleSpiderOutput(result string, resp *response.Response, req *request.Request, spider *spiders.Spider) {

}
