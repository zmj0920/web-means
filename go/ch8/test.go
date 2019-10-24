package main

import "fmt"

func main() {
	test()
	test1()
}

func test() {
	var s string ////初始化为默认零值“”

	s = "hello"

	fmt.Println(len(s)) //五个字符
	//s[1] = '3' //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //严可以存储任何二进制数据
	//s = "\xE4\xBA\xBB\xFF"
	fmt.Println(s)

	fmt.Println(len(s)) //汉子3个字符

	s = "中"
	fmt.Println(len(s)) //是byte数

	c := []rune(s)

	fmt.Println(len(c))

	fmt.Printf("中 unicode %x", c[0])
	fmt.Println()
	fmt.Printf("中 UTF8 %x", s)
	fmt.Println()
}

func test1() {
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c %[1]x", c)
		fmt.Println()
	}
}
