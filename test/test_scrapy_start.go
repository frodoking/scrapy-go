package main

import (
	"fmt"
	"scrapy/common"
)

func main() {
	logger := common.WithLogger("ScrapyTest")
	//testScrapy := scrapy.Scrapy{}
	//testScrapy.Execute(nil, nil)

	var TestSignal common.Signal = 10
	for i := 1; i <= 10; i++ {
		listener := common.ScrapySignal.Connect(TestSignal)
		go func(position int) {
			for {
				select {
				case event := <-listener:
					if event != nil {
						logger.Info("[%s] received : %s ", position, event.(string))
						if event == "stop" {
							return
						}
					}
				}
			}
		}(i)
	}

	result1 := common.ScrapySignal.Send(TestSignal, "send message 111")
	logger.Info("send message 111 success: ", result1)
	result2 := common.ScrapySignal.Send(TestSignal, "send message 222")
	logger.Info("send message 222 success: ", result2)
	result3 := common.ScrapySignal.Send(TestSignal, "send message 333")
	logger.Info("send message 333 success: ", result3)
	resultxx := common.ScrapySignal.DisconnectAll(TestSignal)
	logger.Info("close channel aa success: ", resultxx)
	result4 := common.ScrapySignal.Send(TestSignal, "send message 444")
	logger.Info("send message 444 success: ", result4)

	fmt.Scanln()
}
