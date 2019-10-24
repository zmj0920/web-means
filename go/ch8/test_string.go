package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test()
	TestConv()
}

//Split切片
func test() {
	s := "A,C,D"
	//截取
	parts := strings.Split(s, ",")

	for _, part := range parts {
		fmt.Println(part)
	}
	//拼接
	fmt.Println(strings.Join(parts, "-"))
}

func TestConv() {
	s := strconv.Itoa(10)
	fmt.Println("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + i)
	}
}
