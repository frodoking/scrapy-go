package commands

import (
	"regexp"
	"scrapy/http/request"
	"scrapy/spiders"
)

type FetchCommand struct {
	*ScrapyCommand
	crawlerProcess interface{}
}

func (c *FetchCommand) Syntax() string {
	return "[options] <url>"
}

func (c *FetchCommand) ShortDesc() string {
	return "Fetch a URL using the Scrapy downloader"
}

func (c *FetchCommand) LongDesc() string {
	return "to stdout. You may want to use --nolog to disable logging " +
		"Fetch a URL using the Scrapy downloader and print its content "
}

func (c *FetchCommand) Run(args []string, opts []string) {
	if len(args) != 1 || checkUrl(args[0]) {
		panic("xxxx")
	}

	var request = request.NewRequest(args[0], "")
	var spidercls = spiders.NewDefaultSpider()
	var spiderLoader = c.crawlerProcess

}

func checkUrl(url string) bool {
	if ok, _ := regexp.MatchString("^http?:/{2}w.+$", url); !ok {
		return false
	}
	return true
}
