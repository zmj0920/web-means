package main

import "fmt"

func main() {
	test()
}

func swap(x, y string) (string, string) {
	return x, y
}

func test() {
	if a := 9; a < 10 {
		fmt.Println(a)
	}
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
}
