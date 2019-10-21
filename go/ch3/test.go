package main

import "fmt"

//运算符的使用
func main() {
	test()
	test1()
	test2()
}

// + - * / % ++  --   == !=  > <  >= <=  逻辑运算符 && || !

//位运算符  & |  ^ 双目运算符 << 左移运算符  >> 右移运算符  &^ 按位清零运算符

//不支持前置++ -- ++a --a 是错误的
func test() {

	a := 10
	b := 2

	fmt.Println("a + b=", a+b) //12

	fmt.Println("a - b=", a-b) //8

	fmt.Println("a * b=", a*b) //20

	fmt.Println("a / b=", a/b) //5

	fmt.Println("a % b=", a%b) //0

	//fmt.Println(a++) //

	//fmt.Println(a--) //
}

//数组 ==比较运算
func test1() {

	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	//	c := [...]int{1, 2, 3, 4}
	fmt.Println(a == b)
	//长度不等的数组不能进行比较
	//fmt.Println(c == b)

}

//连续位常量
func test2() {
	// iota常量计数器
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	// 0111 即 7,表示可读,可写,可执行
	accessCode := 7
	//清除可读
	accessCode = accessCode &^ Readable
	//清除可执行
	accessCode = accessCode &^ Executable

	fmt.Println(accessCode&Readable == Readable, accessCode&Writable == Writable, accessCode&Executable == Executable)
}
