package scrapy

import (
	"os"
	"scrapy/actions"
	"scrapy/settings"
	"scrapy/common"
)

type Scrapy struct {
}

func (s *Scrapy) Execute(argv []string, mSettings *settings.Settings) {
	if argv == nil {
		argv = os.Args
	}

	if mSettings == nil {
		mSettings = settings.NewSettings()
	}

	if mSettings == nil {
		mSettings = GetProjectSettings()
	}

	crawlAction := &actions.CrawlAction{ScrapyAction: actions.ScrapyAction{Name: "", RequiresProject: false}}
	crawlAction.Settings = mSettings
	crawlAction.CrawlerProcess = common.NewCrawlerProcess()

	go crawlAction.Run(argv, nil)
}
