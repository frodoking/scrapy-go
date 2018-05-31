package main

import (
	"log"
	"sync"
	"time"
)

var (
	domainSyncChan = make(chan int, 10)
)

func domainPut(num int) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("error to chan put.", num)
		}
	}()

	domainSyncChan <- num

	if num%2 == 0 {
		panic("error....")
	}

	for x := range domainSyncChan {
		log.Println(num, " xxxxx >>> ", x)
	}

	v, ok := <-domainSyncChan
	log.Println(v, ok)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func worker(done chan bool) {
	time.Sleep(time.Second * 3)
	// 通知任务已完成
	done <- true
}

func calc(w *sync.WaitGroup, i int) {
	log.Println("calc: ", i)
	time.Sleep(time.Second)
	w.Done()
}

func testChannel() {
	for i := 0; i < 10; i++ {
		domainName := i
		go domainPut(domainName)
	}
	time.Sleep(time.Second * 3)

	close(domainSyncChan)
}

func testMultiChannel() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	log.Println(x, y, x+y)
}

func testTimeTicker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			log.Println("Tick at", t)
		}
	}()
}

func testWork() {
	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	log.Println("Wait Task")
	<-done
	log.Println("Task Done")
}

func testMultiWork() {
	// Goroutine 例子（等待所有任务退出主程序再退出）
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}
	wg.Wait()
	log.Println("all goroutine finish")
}

func testTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		log.Println(res)
	case <-time.After(time.Second * 1):
		log.Println("timeout 1")
	}
}

func main() {
	testChannel()
	testMultiChannel()
	testTimeTicker()
	testWork()
	testMultiWork()
	testTimeout()
	time.Sleep(time.Second * 100)
}
