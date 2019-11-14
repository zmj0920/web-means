package main

import "fmt"

func main() {
	DoSomething(10)
	judgeType("技术菜")
}

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
	}
	fmt.Println("Unknow Type")
}

func judgeType(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("the type of a is int", v)
	case string:
		fmt.Println("the type of a is string", v)
	case float64:
		fmt.Println("the type of a is float", v)
	default:
		fmt.Println("unknown type")
	}
}
