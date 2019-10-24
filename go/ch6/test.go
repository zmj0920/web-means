package main

import "fmt"

func main() {
	test()
	test1()
	test2()
	test3()
}

//make切片
func test() {
	//make可以创建切片，也可以用来创建Map
	m3 := make(map[int]int, 10)
	fmt.Printf("len m3=%d\n", len(m3))
	// 创建了key类型为string，value类型为int的空Map month
	month := make(map[string]int)
	month["January"] = 1
	month["February"] = 2
	month["March"] = 3

}

func test1() {
	//var month map[string]int
	//fmt.Println(month == nil) // 输出：true

	//map创建数组
	m := map[string]string{"id": "1", "name": "zmj", "age": "9"}
	//取值        长度
	fmt.Println(m["name"], len(m))

	m2 := map[int]int{}
	m2[4] = 16
	//有空Map 对于nil的map是不能存取键值对的 否则就会报错
	fmt.Println(m2 == nil) // false
	fmt.Printf("len m2=%d\n", len(m2))

	m1 := map[int][]int{}
	m1[2] = []int{1, 2, 3}
	fmt.Println(m1[2])
	m1[3] = []int{1, 2, 3}
	//删除指定的key 值
	delete(m1, 3)
	if v, ok := m1[3]; ok {
		fmt.Printf("Key 3's value is %d\n", v)
	} else {
		fmt.Println("key 3 is not existing.")
	}
}

//遍历map
func test2() {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

func test3() {

	slice := []int{1, 2, 3}
	m := map[string][]int{"slice": slice}
	fmt.Println(m)

}
