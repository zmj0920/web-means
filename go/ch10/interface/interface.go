package main

import "fmt"

func main() {
	testClient()
}

//定义接口
type programmer interface {
	//方法
	writeHelloWorld() string
}

//接口实现
type goProgrammer struct {
}

//定义方法
func (g *goProgrammer) writeHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func testClient() {
	//定义接口变量
	var p programmer

	p = new(goProgrammer)
	fmt.Println(p.writeHelloWorld())
}
