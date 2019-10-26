package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// testFn()
	// testVarParam()
	testDefer()
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

//测试函数
func testFn() {
	a, b := returnMultiValues()
	fmt.Println(a, b)
	//传入一个函数
	tsSF := timeSpent(slowFun)
	//传10返回10
	fmt.Println(tsSF(20))
}

//可变参数
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

//可变长参数测试
func testVarParam() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func clear() {
	fmt.Println("Clear resources.")
}

//defer函数延迟执行 释放放资源释放锁使用
func testDefer() {
	defer clear()

	defer func() {
		fmt.Println(666)
	}()

	fmt.Println("Start")
	panic("err") //defer仍会执⾏
	// fmt.Println("End")
}
