package main

import "log"
import (
	"scrapy/crawler"
	"runtime"
)

func main() {
	log.Println(crawler.NewCrawlerProcess().Print())
	v := []int {100, 200, 300, 400}

	for i,item := range v {
		println("-----",i, item)
	}

	a:= append(v, 123)
	for i,item := range v {
		println("+++++", i, item)
	}

	for i,item := range a {
		println(">>>>>>", i, item)
	}

	num := runtime.NumCPU() //本地机器的逻辑CPU个数
	runtime.GOMAXPROCS(num) //设置可同时执行的最大CPU数，并返回先前的设置
	log.Println(num)
}
