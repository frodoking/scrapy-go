package main

import (
	"container/list"
	"fmt"
	"log"
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
	a := 1                               //line 1
	b := 2                               //2
	defer calc("1", a, calc("10", a, b)) //3
	a = 0                                //4
	defer calc("2", a, calc("20", a, b)) //5
	b = 1                                //6
}

func testList() {
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  //首部元素之后插入一个值为10的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,061234
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,601234
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,460123
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出l2的值,460123
	}
	fmt.Println("")
}

type aa interface {
	MyPrint()
}

type bb struct {
	str string
}

func (b *bb) MyPrint() {
	println(b.str)
}

type cc struct {
	str string
}

func (c *cc) Myprint(a aa) {
	a.MyPrint()
}

func testInterface2() {
	b := &bb{"bbbbbb"}
	c := &cc{"cccccc"}
	c.Myprint(b)
	b.str = "xxxxxxx"
	c.Myprint(b)
}

func testChan(name string, callback chan interface{}) {
	callback <- name
}

func testChan2(name string) chan interface{} {
	result := make(chan interface{})
	go func() {
		result <- name
	}()
	return result
}

func testClosure() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		x := i
		s = append(s, func() {
			fmt.Println(&x, x)
		})
	}

	return s
}

func main() {
	testInterface()
	testDeferCall()
	testCalc()
	testList()
	testInterface2()

	callback := make(chan interface{})
	go testChan("frodo", callback)
	result := <-callback

	println("First callback >>>", result)

	callback2 := testChan2("frdoking")
	result2 := <-callback2
	aa := result2.(string)
	println("Second callback >>>", aa)

	for _, f := range testClosure() {
		f()
	}
}
