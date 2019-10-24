package main

import "fmt"

func main() {
	//test()
	mapForSet()
}

//map工厂模式
func test() {
	//Map 的 value 可以是⼀个⽅法
	m := map[int]func(op int) int{}

	m[1] = func(op int) int { return op }

	m[2] = func(op int) int { return op * op }

	m[3] = func(op int) int { return op * op * op }

	fmt.Println(m[1](2), m[2](2), m[3](2))

}

//使用map实现set
func mapForSet() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
	mySet[3] = true
	fmt.Println(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}

}
