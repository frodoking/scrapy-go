package downloader

import (
	"scrapy/core/downloader/handlers"
	"scrapy/http/request"
	"scrapy/spiders"
)

type Slot struct {
	concurrency    uint8
	delay          uint16
	randomizeDelay uint8
	active         map[int]bool
	queue          []string
	transferring   map[int]bool
	lastSeen       uint
	laterCall      interface{}
}

type Downloader struct {
	settings         map[string]string
	signals          string
	slots            map[string]string
	active           map[int]bool
	handlers         *handlers.DownloadHandlers
	totalConcurrency uint
	status           uint8
	body             []byte
	request          *request.Request
	flags            []string
}

func (d *Downloader) Fetch(request *request.Request, spider *spiders.Spider) map[int]bool {
	return nil
}

func (d *Downloader) NeedsBackout() bool {
	return false
}
