package core

import (
	"scrapy/common"
	"scrapy/dupefilter"
	"scrapy/http/request"
	"scrapy/settings"
	"scrapy/spiders"
	"strconv"
)

type Scheduler struct {
	df       dupefilter.BaseDupeFilter
	dqDir    string
	pqClass  string
	dqClass  string
	mqClass  string
	logunser bool
	stats    uint8
	spider   spiders.Spider
	mqs      common.SerializableQueue
	dqs      common.SerializableQueue
}

func NewScheduler(settings *settings.Settings) *Scheduler {
	df := dupefilter.NewRFPDupeFilter(settings)
	pqClass := settings.Get("SCHEDULER_PRIORITY_QUEUE")
	dqClass := settings.Get("SCHEDULER_DISK_QUEUE")
	mqClass := settings.Get("SCHEDULER_MEMORY_QUEUE")
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

	return &Scheduler{df: df, pqClass: pqClass, dqClass: dqClass, mqClass: mqClass, logunser: logunser}
}

func (s *Scheduler) HasPendingRequests() bool {
	return false
}

func (s *Scheduler) Open(spider spiders.Spider) chan string {
	s.spider = spider
	s.mqs = &common.PriorityQueue{}
	s.dqs = &common.LifoMemoryQueue{}
	return s.df.Open()
}

func (s *Scheduler) Close(reason string) chan string {
	return s.df.Close(reason)
}

func (s *Scheduler) EnqueueRequest(request *request.Request) bool {
	if request.DontFilter && s.df.RequestSeen(request) {
		s.df.Log(request, s.spider)
		return false
	}
	dqok := s.dqPush(request)
	if dqok {

	} else {
		s.mqPush(request)
	}
	return true
}

func (s *Scheduler) NextRequest() *request.Request {
	req := s.mqs.Pop().(*request.Request)
	if req == nil {
		req = s.dqs.Pop().(*request.Request)
	}
	return req
}

func (s *Scheduler) dqPush(req *request.Request) bool {
	if s.dqs == nil {
		return false
	}
	s.dqs.Push(req)
	return true
}

func (s *Scheduler) mqPush(req *request.Request) bool {
	if s.mqs == nil {
		return false
	}
	s.mqs.Push(req)
	return true
}
