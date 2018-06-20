package core

import (
	"reflect"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/middleware/spidermiddlewares"
	"scrapy/spiders"
)

const (
	MinResponseSize = 1024
)

type ResponseRequestDefer struct {
	Resp  *response.Response
	Req   *request.Request
	Defer chan interface{}
}

type ScraperSlot struct {
	maxActiveSize uint
	queue         []*ResponseRequestDefer
	active        map[*request.Request]bool
	activeSize    int
	itemProcSize  uint
	closing       chan interface{}
}

func NewScraperSlot(maxActiveSize uint) *ScraperSlot {
	if maxActiveSize == -1 {
		maxActiveSize = 5000000
	}
	return &ScraperSlot{maxActiveSize, make([]*ResponseRequestDefer, 10), make(map[*request.Request]bool), 0, 0, nil}
}

func (ss *ScraperSlot) AddResponseRequest(resp *response.Response, req *request.Request) chan interface{} {
	deferred := make(chan interface{})
	ss.queue = append(ss.queue, &ResponseRequestDefer{resp, req, deferred})

	if reflect.TypeOf(resp).Name() == "Response" {
		max := len(resp.Body)
		if max < MinResponseSize {
			max = MinResponseSize
		}
		ss.activeSize += max
	} else {
		ss.activeSize += MinResponseSize
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
		if max < MinResponseSize {
			max = MinResponseSize
		}
		ss.activeSize -= max
	} else {
		ss.activeSize -= MinResponseSize
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

func (s *Scraper) CloseSpider(spider *spiders.Spider) chan interface{} {
	slot := s.slot
	slot.closing = make(chan interface{})
	s.checkIfClosing(spider, slot)
	return slot.closing
}

func (s *Scraper) IsIdle() bool {
	return s.slot == nil
}

func (s *Scraper) checkIfClosing(spider *spiders.Spider, slot *ScraperSlot) {
	if <-slot.closing != nil && slot.IsIdle() {

	}
}

func (s *Scraper) EnqueueScrape(resp *response.Response, req *request.Request, spider *spiders.Spider) chan interface{} {
	slot := s.slot
	result := slot.AddResponseRequest(resp, req)

	for {
		select {
		case <-result:
			slot.FinishResponse(resp, req)
			s.checkIfClosing(spider, slot)
		}
	}
	s.scrapeNext(spider, slot)
	return result
}

func (s *Scraper) scrapeNext(spider *spiders.Spider, slot *ScraperSlot) {
	for _, defered := range slot.queue {
		nextDeferred := slot.NextResponseRequestDeferred()
		s.scrape(nextDeferred.Resp, nextDeferred.Req, spider)
	}
}

func (s *Scraper) scrape(resp *response.Response, req *request.Request, spider *spiders.Spider) {

}

func (s *Scraper) scrape2(resp *response.Response, req *request.Request, spider *spiders.Spider) chan interface{} {

}

func (s *Scraper) HandleSpiderError(failure string, resp *response.Response, req *request.Request, spider *spiders.Spider) {

}

func (s *Scraper) HandleSpiderOutput(result string, resp *response.Response, req *request.Request, spider *spiders.Spider) {

}
