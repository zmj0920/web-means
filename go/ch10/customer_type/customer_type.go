package main

import (
	"fmt"
	"time"
)

func main() {
	testFn()
}

type intConv func(op int) int

func timeSpent(inner intConv) intConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func testFn() {
	tsSF := timeSpent(slowFun)
	fmt.Println(tsSF(10))
}
