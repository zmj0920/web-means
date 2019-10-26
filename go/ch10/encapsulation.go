package main

import (
	"fmt"
	"unsafe"
)

func main() {
	testStructOperations()
	testCreateEmployeeObj()
}

//结构体定义  包内使用首字母小写，包外使用首字母大写
type employee struct {
	id   string
	name string
	age  int
}

//实例创建及初始化
func testCreateEmployeeObj() {
	e := employee{"0", "Bob", 20}
	e1 := employee{name: "Mike", age: 30}
	e2 := new(employee) //注意这⾥返回的引⽤/指针，相当于 e := &Employee{}
	e2.id = "2"         //与其他主要编程语⾔的差异：通过实例的指针访问成员不需要使⽤箭头函数>
	e2.age = 22
	e2.name = "Rose"
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e1.id)
	fmt.Println(e2)
	fmt.Printf("e is %T", e) //实例类型
	fmt.Println()
	fmt.Printf("e2 is %T", e2) //指针类型
}

//指针定义通常情况下为了避免内存拷⻉我们使⽤第⼆种定义⽅式 new
func (e *employee) test() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.id, e.name, e.age)
}

//第⼀种定义⽅式在实例对应⽅法被调⽤时，实例的成员会进⾏值复制
// func (e employee) test() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.id, e.name, e.age)
// }

func testStructOperations() {
	e := employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.name))
	fmt.Println(e.test()) //指针实例访问都是点调用
}
