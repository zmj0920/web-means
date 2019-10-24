package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test()
}

//多返回值函数
func returnMultiValues() (int, int) {
	//返回随机数
	return rand.Intn(10), rand.Intn(20)
}

//计时函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		//开始时间
		start := time.Now()
		//调用内部函数
		ret := inner(n)
		//输入内部函数运行的时间
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

//非常慢的函数
func slowFun(op int) int {
	//延迟1秒
	time.Sleep(time.Second * 1)
	return op
}

func test() {
	a, b := returnMultiValues()
	fmt.Println(a, b)
	//传入一个函数
	tsSF := timeSpent(slowFun)
	//传10返回10
	fmt.Println(tsSF(10))
}
