package spiders

type Spider struct {
	concurrency    uint
	delay          uint
	randomizeDelay uint
	active         map[int]bool
	queue          []string
	transferring   map[int]bool
	lastSeen       uint
	laterCall      interface{}
}
