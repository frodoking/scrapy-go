package main

import (
	"log"
	"fmt"
)

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string `json:"jsonop" xml:"xmlOpName"`
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func animal(animal Animal) {
	log.Println(animal.Speak())
}

func testInterface() {
	dog := &Dog{}
	animal(dog)
	cat := &Cat{}
	animal(cat)
}

// defer函数属延迟执行，延迟到调用者函数执行 return 命令前被执行。多个defer之间按LIFO先进后出顺序执行。
func testDeferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	// panic("触发异常")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer calc func中参数b必须先计算. defer是在函数末尾的return前执行，先进后执行
func testCalc() {
	a := 1                                             //line 1
	b := 2                                             //2
	defer calc("1", a, calc("10", a, b))  //3
	a = 0                                              //4
	defer calc("2", a, calc("20", a, b))  //5
	b = 1                                              //6
}

func main() {
	testInterface()
	testDeferCall()
	testCalc()
}
