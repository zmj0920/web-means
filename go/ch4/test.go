package main

import "fmt"

func main() {
	test()
	test1()
	test2()
}

//函数多返回值
func swap(x, y string) (string, string) {
	return x, y
}

//if 条件判断
func test() {
	if a := 9; a < 10 {
		fmt.Println(a)
	}
	//接收多返回值函数判断
	if c, d := swap("hello", "world"); c != "" || d != "" {
		fmt.Println(c, d)
	}

}

//switch条件判断支持多个值判断
func test1() {
	for i := 0; i < 5; i++ {
		switch i {
		//支持多个值条件判断只要一个满足就会执行 case 默认不加break
		case 0, 2:
			fmt.Println("Even")
		case 1, 3:
			fmt.Println("odd")
		default:
			fmt.Println("it is not 0-3")
		}
	}
}

//if 形式switch形式判断 switch后可以没有表达式，只在case后面直接打表达式即可
func test2() {
	for i := 0; i < 5; i++ {
		switch {
		//支持多个值条件判断只要一个满足就会执行 case 默认不加break
		case i%2 == 0:
			fmt.Println("Even")
		case i%2 == 1:
			fmt.Println("odd")
		default:
			fmt.Println("unknow")
		}
	}
}
