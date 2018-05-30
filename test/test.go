package main

import "log"
import (
	"scrapy/crawler"
)

func main() {
	log.Println(crawler.NewCrawlerProcess().Print())
	v := []int {100, 200, 300, 400}
	for i,item := range v {
		println(i, item)
	}
}
