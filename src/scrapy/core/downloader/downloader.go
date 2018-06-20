package downloader

import (
	"container/list"
	"math/rand"
	"reflect"
	"scrapy/core/downloader/handlers"
	"scrapy/crawler"
	"scrapy/http/request"
	"scrapy/middleware/downloadermiddlewares"
	"scrapy/spiders"
	"time"
)

type DownloaderSlot struct {
	concurrency    int
	delay          float64
	randomizeDelay bool
	active         map[*request.Request]bool
	queue          *list.List
	transferring   map[*request.Request]bool
	LastSeen       float64
	laterCall      interface{}
}

func (ds *DownloaderSlot) freeTransferSlots() int {
	return ds.concurrency - len(ds.transferring)
}

func (ds *DownloaderSlot) downloadDelay() float64 {
	if ds.randomizeDelay {
		return (rand.Float64() + 0.5) * ds.delay
	}
	return ds.delay
}

func (ds *DownloaderSlot) close() {

}

type downloaderRequestDeferred struct {
	request  *request.Request
	deferred chan interface{}
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
	handlers := handlers.NewDownloadHandlers()
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

func (d *Downloader) enqueueRequest(req *request.Request, spider *spiders.Spider) chan interface{} {
	key, slot := d.getSlot(req, spider)
	req.Meta["download_slot"] = key

	slot.active[req] = true

	callback := make(chan interface{}, 1)
	go func() {
		select {
		case result := <-callback:
			delete(slot.active, req)
			callback <- result
		}
	}()

	slot.queue.PushBack(&downloaderRequestDeferred{req, callback})

	return callback
}

func (d *Downloader) processQueue(spider *spiders.Spider, slot *DownloaderSlot) {
	now := float64(time.Now().UnixNano())
	delay := slot.downloadDelay()
	if delay != 0 {
		penalty := delay - now + slot.LastSeen
		if penalty > 0 {
			time.AfterFunc(penalty*time.Millisecond, func() {
				d.processQueue(spider, slot)
			})
			return
		}
	}

	for slot.queue.Len() > 0 && slot.freeTransferSlots() > 0 {
		slot.LastSeen = now
		front := slot.queue.Front()
		slot.queue.Remove(front)
		requestDeferred := front.Value.(downloaderRequestDeferred)
		dfd := d.download(slot, requestDeferred.request, spider)
		select {
		case deferred := <-requestDeferred.deferred:
			if reflect.TypeOf(deferred).Name() == "Response" {
				dfd <- deferred
			}
		}
		if delay != 0 {
			d.processQueue(spider, slot)
			break
		}
	}
}

func (d *Downloader) download(slot *DownloaderSlot, req *request.Request, spider *spiders.Spider) chan interface{} {
	newResult := make(chan interface{}, 1)
	result := d.handlers.DownloadRequest(req, spider)
	select {
	case resp := <-result:
		delete(slot.transferring, req)
		d.processQueue(spider, slot)
		go func() {
			newResult <- resp
		}()
	}
	return newResult
}

func (d *Downloader) getSlot(req *request.Request, spider *spiders.Spider) (string, *DownloaderSlot) {
	return "", nil
}
