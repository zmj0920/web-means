package main

import "fmt"

func main() {
 
}

//切⽚声明
func test3() {
	var s0 []int
	s0 = append(s0, 1) //填充元素
	// s := []int{}
	// s1 := []int{1, 2, 3}
	// s2 := make([]int, 2, 4)
	/*[]type, len, cap
	其中len个元素会被初始化为默认零值，未初始化元素不可以访问
	*/
	fmt.Println(cap(s0))
}

//cap可伸缩
func test4() {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
	//当存储空间不够用了cap*2增长
	// 1 1
	// 2 2
	// 3 4
	// 4 4
	// 5 8
	// 6 8
	// 7 8
	// 8 8
	// 9 16
	// 10 16
}

func test5() {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2)) //[Apr May Jun] 3 9  cap长度不包含前三个
	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer)) //[Jun Jul Aug] 3 7   cap长度不包含前三个
	summer[0] = "Unknow"
	fmt.Println(Q2)
	fmt.Println(year)
}

//数组比较
func test6() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { //切片只能和nil比较
	// 	t.Log("equal")
	// }
	fmt.Println(a, b)
}
