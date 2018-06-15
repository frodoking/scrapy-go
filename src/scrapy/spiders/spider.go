package spiders

import (
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/settings"
)

type Spider struct {
	name           string
	startUrls  		[]string
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

func (s *Spider) SetCrawler(crawler *crawler.Crawler)  {

}

func (s *Spider) StartRequests()  {

}

func (s *Spider) MakeRequestsFromUrl(url string) *request.Request {
	return request.NewRequest(url, "utf-8")
}

func (s *Spider) UpdateSettings(settings *settings.Settings) {
	settings.Set("", "")
}

func (s *Spider) HandlesRequest(request request.Request) interface{} {
	return nil
}


