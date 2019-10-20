package main

import "fmt"

func main() {
	test()
	test1()
	test2()
	test3()
	test4()
}

//斐波数
func test() {
	a := 1
	var b int = 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

//交换赋值
func test1() {
	// :=这个符号直接取代了var和type,这种形式叫做简短声明
	c := 3
	d := 4
	// 变量交换赋值
	c, d = d, c

	fmt.Println(c, d)
}

//连续位常量
func test2() {
	// iota常量计数器
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	// 0001 0010 0100 即 1 2 4
	fmt.Println(Readable, Writable, Executable)
	// 0111 即 7,表示可读,可写,可执行
	accessCode := 7
	fmt.Println(accessCode&Readable == Readable, accessCode&Writable == Writable, accessCode&Executable == Executable)
}

//连续常量赋值
func test3() {
	// iota常量计数器
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
	)
	fmt.Println(Monday, Tuesday, Wednesday)
}

//特殊
func test4() {
	// iota常量计数器
	const (
		Test1 = iota
		Test2 = iota
		//赋相同的值
		Test3, Test4, Test5 = iota, iota, iota
		Test6               = iota
	)
	fmt.Printf("%d\n", Test1)
	fmt.Printf("%d\n", Test2)
	fmt.Printf("%d\n", Test3)
	fmt.Printf("%d\n", Test4)
	fmt.Printf("%d\n", Test5)
	fmt.Printf("%d\n", Test6)
}
