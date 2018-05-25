package main

import "log"
import "scrapy"

func main() {
	log.Println(scrapy.NewCrawlerProcess().Print())
}
