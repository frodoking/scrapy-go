package spiders

import (
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/http/response"
	"scrapy/settings"
)

type Spider struct {
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
}

func NewDefaultSpider() *Spider {
	spider := &Spider{}
	return spider
}

func (s *Spider) SetCrawler(crawler *crawler.Crawler) {

}

func (s *Spider) StartRequests() []*request.Request {
	requests := make([]*request.Request, len(s.startUrls))
	for i, url := range s.startUrls {
		req := s.MakeRequestsFromUrl(url)
		requests[i] = req
	}
	return requests
}

func (s *Spider) MakeRequestsFromUrl(url string) *request.Request {
	return request.NewRequest(url, "utf-8")
}

func (s *Spider) Parse(response *response.Response) {
}

func (s *Spider) UpdateSettings(settings *settings.Settings) {
	settings.Set("", "")
}

func (s *Spider) HandlesRequest(request request.Request) interface{} {
	return nil
}
