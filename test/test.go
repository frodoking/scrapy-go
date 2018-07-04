package main

import (
	"runtime"
	"scrapy/common"
	"scrapy/core"
)

func main() {
	logger := common.WithLogger("test")
	logger.Info(core.NewCrawlerProcess().Print())
	v := []int{100, 200, 300, 400}

	for i, item := range v {
		logger.Info("-----", i, item)
	}

	a := append(v, 123)
	for i, item := range v {
		logger.Info("+++++", i, item)
	}

	for i, item := range a {
		logger.Info(".....", i, item)
	}

	num := runtime.NumCPU() //本地机器的逻辑CPU个数
	runtime.GOMAXPROCS(num) //设置可同时执行的最大CPU数，并返回先前的设置
	logger.Println(num)
	logger.WithField("cc", "dfasdfsd").Info("xxxx")
	logger.WithField("bb", "eeee").Info("xxxx")
}
