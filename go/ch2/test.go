package main

import "fmt"

//数据类型
//不支持隐身转化

func main() {
	test()
	test1()
	testString()
}

//自定义类型
//type Myint int64

func test() {
	var a int32 = 1
	var b int64

	//显式转换
	b = int64(a)
	//	b = Myint(a)
	fmt.Println(a, b)
}

//指针
func test1() {
	a := 1
	aPtr := &a
	//不支持指针运算
	//	aPtr = aPtr + 1
	fmt.Println(a, aPtr)
}

//空字符串
func testString() {
	var s string
	fmt.Println("*" + s + "*")
	//非空判断直接可以等于空字符串
	if s == "" {
		fmt.Println(len(s))
	}
}
