package downloader

import (
	"scrapy/core/downloader/handlers"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/middleware/downloadermiddlewares"
	"scrapy/spiders"
	"scrapy/http/response"
)

type DownloaderSlot struct {
	concurrency    uint8
	delay          uint16
	randomizeDelay uint8
	active         map[*request.Request]bool
	queue          []string
	transferring   map[int]bool
	lastSeen       uint
	laterCall      interface{}
}

type Downloader struct {
	settings          map[string]string
	signals           string
	slots             map[string]*DownloaderSlot
	active            map[*request.Request]bool
	handlers          *handlers.DownloadHandlers
	totalConcurrency  uint8
	domainConcurrency uint8
	ipConcurrency     uint8
	randomizeDelay    bool
	middleware        *middleware.DownloaderMiddlewareManager
}

func NewDownloader(crawler *crawler.Crawler) *Downloader {
	handlers := handlers.NewDownloadHandlers(crawler)
	middleware := &middleware.DownloaderMiddlewareManager{}

	downloader := &Downloader{handlers: handlers, middleware: middleware}

	return downloader
}

func (d *Downloader) Fetch(req *request.Request, spider *spiders.Spider) chan interface{} {
	delete(d.active, req)
	result := d.middleware.Download(d.enqueueRequest, req, spider)
	return result
}

func (d *Downloader) NeedsBackout() bool {
	return false
}

func (d *Downloader) enqueueRequest(req *request.Request, spider *spiders.Spider) func(resp *response.Response) *response.Response {
	key, slot:= d.getSlot(req, spider)
	req.Meta["download_slot"] = key

	slot.active[req] = true

	deactivateCallback := func(resp *response.Response) *response.Response {
		delete(d.active, req)
		return resp
	}
	return deactivateCallback
}

func (d *Downloader) getSlot(req *request.Request, spider *spiders.Spider) (string, *DownloaderSlot) {
	return "", nil
}
