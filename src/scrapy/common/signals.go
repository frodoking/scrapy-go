package common

type Signal int

const (
	EngineStarted Signal = 0x1000
	EngineStopped Signal = 0x1001

	SpiderOpened Signal = 0x2000
	SpiderIdle   Signal = 0x2001
	SpiderClosed Signal = 0x2002
	SpiderError  Signal = 0x2003

	RequestScheduled Signal = 0x3003
	RequestDropped   Signal = 0x3003

	ResponseReceived   Signal = 0x4003
	ResponseDownloaded Signal = 0x4003

	ItemScraped Signal = 0x5003
	ItemDropped Signal = 0x5003
)
