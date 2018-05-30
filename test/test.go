package main

import "log"
import (
	"scrapy"
	"scrapy/settings"
)

func main() {
	log.Println(scrapy.NewCrawlerProcess().Print())
	settings.NewCrawlerSettings()
}
