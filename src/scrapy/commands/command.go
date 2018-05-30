package commands

import "scrapy/settings"
import "scrapy"

type ScrapyCommand struct {
	Name            string
	RequiresProject bool
	settings        settings.Settings
	crawler         scrapy.Crawler
}

func (c *ScrapyCommand) SetCrawler(crawler scrapy.Crawler) {
	c.crawler = crawler
}

func (c *ScrapyCommand) Syntax() string {
	return ""
}

func (c *ScrapyCommand) ShortDesc() string {
	return ""
}

func (c *ScrapyCommand) LongDesc() string {
	return c.ShortDesc()
}

func (c *ScrapyCommand) Help() string {
	return c.LongDesc()
}

// Populate option parse with options available for this command
func (c *ScrapyCommand) AddOptions() {

}

func (c *ScrapyCommand) ProcessOptions() string {
	return ""
}

func (c *ScrapyCommand) Run(args []string, opts []string) {
}
