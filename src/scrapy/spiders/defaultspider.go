package spiders

import (
	"container/list"
	"scrapy/common"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/settings"
)

type DefaultSpider struct {
	name           string
	startUrls      []string
	concurrency    uint
	delay          uint
	randomizeDelay uint
	active         map[int]bool
	queue          []string
	transferring   map[int]bool
	lastSeen       uint
	laterCall      interface{}
	settings       *settings.Settings
}

func NewDefaultSpider() Spider {
	spider := &DefaultSpider{}
	listener := common.ScrapySignal.Connect(common.SpiderClosed)
	logger := common.WithLogger("spiders")
	go func() {
		for {
			select {
			case reason := <-listener:
				if reason != nil {
					logger.Info("received : %s ", reason.(string))
					spider.Close(reason.(string))
					return
				}
			}
		}
	}()
	return spider
}

func (s *DefaultSpider) StartRequests() *list.List {
	requests := list.New()
	for _, url := range s.startUrls {
		req := s.MakeRequestsFromUrl(url)
		requests.PushBack(req)
	}
	return requests
}

func (s *DefaultSpider) MakeRequestsFromUrl(url string) *request.Request {
	return request.NewRequest(url, "utf-8")
}

func (s *DefaultSpider) Parse(response response.Response) {
}

func (s *DefaultSpider) UpdateSettings(settings *settings.Settings) {
	settings.Set("", "")
}

func (s *DefaultSpider) HandlesRequest(request *request.Request) interface{} {
	return nil
}

func (s *DefaultSpider) GetAttrFromSettings(key string) interface{} {
	return s.settings.Get(key)
}

func (s *DefaultSpider) Close(reason string) {
	logger := common.WithLogger("spiders")
	logger.Info("DefaultSpider Closed, because of ", reason)
}
