package dupefilter

import (
	"scrapy/http/request"
	"scrapy/settings"
	"scrapy/spiders"
	"strconv"
)

type BaseDupeFilter interface {
	RequestSeen(request *request.Request) bool
	Open() chan string
	Close(reason string) chan string
	Log(request *request.Request, spider spiders.Spider) chan string
}

/**
  Request Fingerprint duplicates filter
*/
type RFPDupeFilter struct {
	file         string
	fingerprints map[string]bool
	logdupes     bool
	debug        bool
}

func NewRFPDupeFilter(settings *settings.Settings) *RFPDupeFilter {
	debugString := settings.Get("DUPEFILTER_DEBUG")
	debug, err := strconv.ParseBool(debugString)
	if err != nil {
		debug = false
	}
	fingerprints := make(map[string]bool)

	return &RFPDupeFilter{"", fingerprints, true, debug}
}

func (f *RFPDupeFilter) RequestSeen(request *request.Request) bool {
	return false
}

func (f *RFPDupeFilter) Open() chan string {
	defered := make(chan string, 1)

	return defered
}

func (f *RFPDupeFilter) Close(reason string) chan string {
	defered := make(chan string, 1)

	return defered
}

func (f *RFPDupeFilter) Log(request *request.Request, spider spiders.Spider) chan string {
	defered := make(chan string, 1)

	return defered
}
