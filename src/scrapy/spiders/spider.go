package spiders

type Spider struct {
	name           string
	customSettings map[string]string
	concurrency    uint
	delay          uint
	randomizeDelay uint
	active         map[int]bool
	queue          []string
	transferring   map[int]bool
	lastSeen       uint
	laterCall      interface{}
}

func NewDefaultSpider() *Spider {
	spider := &Spider{}
	return spider
}
