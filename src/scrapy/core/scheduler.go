package core

import (
	"scrapy/dupefilter"
	"scrapy/http/request"
	"scrapy/spiders"
	"strconv"
	"scrapy/settings"
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

func NewScheduler(settings *settings.Settings) *Scheduler {
	df := dupefilter.NewRFPDupeFilter(settings)
	//pqclass := crawler.Settings.Get("SCHEDULER_PRIORITY_QUEUE")
	//dqclass := crawler.Settings.Get("SCHEDULER_DISK_QUEUE")
	//mqclass := crawler.Settings.Get("SCHEDULER_MEMORY_QUEUE")
	debugString := settings.Get("SCHEDULER_DEBUG")
	debug, err := strconv.ParseBool(debugString)
	if err != nil {
		debug = false
	}

	logunserString := settings.Get("LOG_UNSERIALIZABLE_REQUESTS")
	logunser, err := strconv.ParseBool(logunserString)
	if err != nil {
		logunser = false
	} else {
		logunser = debug
	}

	return &Scheduler{df, "", nil, nil, nil, logunser, 1}
}

func (s *Scheduler) HasPendingRequests() bool {
	return false
}

func (s *Scheduler) Open(spider spiders.Spider) chan string {
	return s.df.Open()
}

func (s *Scheduler) Close(reason string) chan string {
	return s.df.Close(reason)
}

func (s *Scheduler) EnqueueRequest(request *request.Request) bool {

	return false
}

func (s *Scheduler) NextRequest() *request.Request {
	return nil
}
