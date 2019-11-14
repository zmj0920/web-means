package main

import "fmt"

func main() {
	slice := make([]int, 10)
	fmt.Println(len(slice), cap(slice)) //10  10

	slice1 := make([]int, 10, 20)
	fmt.Println(len(slice1), cap(slice1)) //10  20
}
