package downloader

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
	handlers         *DownloadHandlers
	totalConcurrency uint
	status           int
	body             []byte
	request          *Request
	flags            []string
}

func (downloader *Downloader) Fetch(request *Request, spider *Spider) map[int]bool {
	return nil
}

func (downloader *Downloader) NeedsBackout() bool {
	return false
}
