package core

import (
	"scrapy/crawler"
	"scrapy/dupefilter"
	"strconv"
	"scrapy/http/request"
)

type Scheduler struct {
	df       dupefilter.BaseDupeFilter
	dqdir    string
	pqclass  interface{}
	dqclass  interface{}
	mqclass  interface{}
	logunser bool
	stats    uint8
}

func NewScheduler(crawler *crawler.Crawler) *Scheduler {
	df := dupefilter.NewRFPDupeFilter(crawler.Settings)
	//pqclass := crawler.Settings.Get("SCHEDULER_PRIORITY_QUEUE")
	//dqclass := crawler.Settings.Get("SCHEDULER_DISK_QUEUE")
	//mqclass := crawler.Settings.Get("SCHEDULER_MEMORY_QUEUE")
	debugString := crawler.Settings.Get("SCHEDULER_DEBUG")
	debug, err := strconv.ParseBool(debugString)
	if err != nil {
		debug = false
	}

	logunserString := crawler.Settings.Get("LOG_UNSERIALIZABLE_REQUESTS")
	logunser, err := strconv.ParseBool(logunserString)
	if err != nil {
		logunser = false
	} else {
		logunser = debug
	}

	return &Scheduler{df, nil, nil, nil, nil, logunser, 1}
}

func (s *Scheduler) HasPendingRequests() bool {
	return false
}

func (s *Scheduler) Open() chan string {
	return s.df.Open()
}

func (s *Scheduler) Close() chan string {
	return s.df.Close()
}

func (s *Scheduler) EnqueueRequest(request *request.Request) bool {

	return false
}

func (s *Scheduler) NextRequest() *request.Request {
	return nil
}