package scrapy

import (
	"os"
	"scrapy/actions"
	"scrapy/crawler"
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
	crawlAction.Settings = mSettings.Settings
	crawlAction.CrawlerProcess = crawler.NewCrawlerProcess()

	go crawlAction.Run(argv, nil)
}
