package scrapy

import (
	"scrapy/spiders"
	"fmt"
	"scrapy/settings"
)

type Crawler struct {
	spider *spiders.Spider
	settings *settings.Settings
}

func (c *Crawler) spiders() []spiders.Spider {
	return nil
}

func (c *Crawler) Crawl() {

}

func (c *Crawler) Stop() {

}

type CrawlerRunner struct {
	aa string
}

func (cr *CrawlerRunner) Crawl() {

}

func (cr *CrawlerRunner) Stop() {

}

func (cr *CrawlerRunner) Print() string {
	return fmt.Sprintf("%s .......", cr.aa)
}

type CrawlerProcess struct {
	*CrawlerRunner
	bb string
}

func (cp *CrawlerProcess) Start(stopAfterCrawl bool) {

}

func (cp *CrawlerProcess) Print() string {
	return fmt.Sprintf("%s ....... %s ", cp.aa, cp.bb)
}

func NewCrawlerProcess() *CrawlerProcess {
	return &CrawlerProcess{&CrawlerRunner{"aa"}, "bbb"}
}


