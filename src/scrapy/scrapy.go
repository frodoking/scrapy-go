package scrapy

import (
	"os"
	"scrapy/actions"
	"scrapy/settings"
)

type Scrapy struct {
}

func (s *Scrapy) Execute(argv []string, mSettings *settings.CrawlerSettings) {
	if argv == nil {
		argv = os.Args
	}

	if mSettings == nil {
		mSettings = settings.NewCrawlerSettings()
	}

	if mSettings == nil {
		mSettings = GetProjectSettings()
	}

	crawlAction := &actions.CrawlAction{ScrapyAction: actions.ScrapyAction{Name: "", RequiresProject: false}}
	crawlAction.Run(argv, nil)
	print("xxxx")
}
