package downloader

import (
	"container/list"
	"math/rand"
	"reflect"
	"scrapy/core/downloader/handlers"
	"scrapy/http/request"
	"scrapy/middleware/downloadermiddlewares"
	"scrapy/spiders"
	"time"
)

type Slot struct {
	concurrency    int
	delay          float64
	randomizeDelay bool
	active         map[*request.Request]bool
	queue          *list.List
	transferring   map[*request.Request]bool
	LastSeen       float64
	laterCall      interface{}
}

func (ds *Slot) freeTransferSlots() int {
	return ds.concurrency - len(ds.transferring)
}

func (ds *Slot) downloadDelay() float64 {
	if ds.randomizeDelay {
		return (rand.Float64() + 0.5) * ds.delay
	}
	return ds.delay
}

func (ds *Slot) close() {

}

type downloaderRequestDeferred struct {
	request  *request.Request
	deferred chan interface{}
}

type Downloader struct {
	settings          map[string]string
	signals           string
	slots             map[string]*Slot
	active            map[*request.Request]bool
	handlers          *handlers.DownloadHandlers `获取指定url数据的真正执行者`
	totalConcurrency  uint8
	domainConcurrency uint8
	ipConcurrency     uint8
	randomizeDelay    bool
	middleware        *middleware.DownloaderMiddlewareManager `downloader中间件`
}

func NewDownloader() *Downloader {
	handlers := handlers.NewDownloadHandlers()
	middleware := &middleware.DownloaderMiddlewareManager{}

	downloader := &Downloader{handlers: handlers, middleware: middleware}

	return downloader
}

func (d *Downloader) Fetch(req *request.Request, spider spiders.Spider) chan interface{} {
	d.active[req] = true
	_, err := d.middleware.Download(req, spider)
	if err != nil {

	} else {
		d.enqueueRequest(req, spider)
	}
	delete(d.active, req)
	return nil
}

func (d *Downloader) NeedsBackout() bool {
	return false
}

func (d *Downloader) enqueueRequest(req *request.Request, spider spiders.Spider) chan interface{} {
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

	d.processQueue(spider, slot)

	return callback
}

func (d *Downloader) processQueue(spider spiders.Spider, slot *Slot) {
	now := float64(time.Now().UnixNano())
	delay := slot.downloadDelay()
	if delay != 0 {
		penalty := delay - now + slot.LastSeen
		if penalty > 0 {
			time.AfterFunc(time.Duration(penalty)*time.Millisecond, func() {
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

func (d *Downloader) download(slot *Slot, req *request.Request, spider spiders.Spider) chan interface{} {
	slot.transferring[req] = true
	newResult := make(chan interface{}, 1)
	result := d.handlers.DownloadRequest(req, spider)
	select {
	case resp := <-result:
		if reflect.TypeOf(resp).Name() == "Response" {
			// Notify response_downloaded listeners about the recent download
			// before querying queue for next request
		}
		delete(slot.transferring, req)
		d.processQueue(spider, slot)
		go func() {
			newResult <- resp
		}()
	}
	return newResult
}

func (d *Downloader) getSlot(req *request.Request, spider spiders.Spider) (string, *Slot) {
	return "", nil
}
