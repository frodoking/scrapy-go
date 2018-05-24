package downloader

import "scrapy/core/downloader/handlers"
import "scrapy/http"
import "scrapy/spider"

type Slot struct {
	concurrency    uint
	delay          uint
	randomizeDelay uint
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
	handlers         *downloader.DownloadHandlers
	totalConcurrency uint
	status           int
	body             []byte
	request          *http.Request
	flags            []string
}

func (d *Downloader) Fetch(request *http.Request, spider *spider.Spider) map[int]bool {
	return nil
}

func (d *Downloader) NeedsBackout() bool {
	return false
}
