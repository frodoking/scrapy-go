package main

import (
	"scrapy/common"
	"fmt"
)

func main() {
	//testScrapy := scrapy.Scrapy{}
	//testScrapy.Execute(nil, nil)

	for i := 1; i <= 10; i++ {
		listener := common.ScrapySignal.Connect("aa")
		go func(position int) {
			for {
				select {
				case event := <-listener:
					if event !=nil {
						println("[", position, "] received : ", event.(string))
						if event == "stop" {
							return
						}
					}
				}
			}
		}(i)
	}

	result1 := common.ScrapySignal.Send("aa", "send message 111")
	println("send message 111 success: ", result1)
	result2 := common.ScrapySignal.Send("aa", "send message 222")
	println("send message 222 success: ", result2)
	result3 := common.ScrapySignal.Send("aa", "send message 333")
	println("send message 333 success: ", result3)
	resultxx := common.ScrapySignal.DisconnectAll("aa")
	println("close channel aa success: ", resultxx)
	result4 := common.ScrapySignal.Send("aa", "send message 444")
	println("send message 444 success: ", result4)

	fmt.Scanln()
}
