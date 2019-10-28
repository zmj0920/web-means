package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

type Pupil struct {
	Student //嵌入Student匿名结构体
}

type Graduate struct {
	Student //嵌入Student匿名结构体
}

func (stu *Student) Showinfo() {
	fmt.Printf("学生姓名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (p *Pupil) test() {
	fmt.Println("小学生正在考试...")
}

func (p *Graduate) test() {
	fmt.Println("大学生正在考试...")
}

func main() {
	//当对结构体嵌入使用方法发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 10
	pupil.test()
	pupil.Student.SetScore(90)
	pupil.Student.Showinfo()

	g := &Graduate{}
	g.Student.Name = "小黑"
	g.Student.Age = 20
	g.test()
	g.Student.SetScore(80)
	g.Student.Showinfo()

}
