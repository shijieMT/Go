package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Time string
}

func (p People) Info() {
	fmt.Println("people ", p.Time)
}

// Student 定义结构体
type Student struct {
	People
	Name string
	Age  int
}

// PrintInfo 给机构体绑定一个方法
func (s Student) PrintInfo() {
	fmt.Printf("time:%s name:%s age:%d\n", s.Time, s.Name, s.Age)
}

// 想在函数里面或者方法里面修改结构体里面的属性
// 只能使用结构体指针或者指针方法
func SetAge(info Student, age int) {
	info.Age = age
}
func SetAge1(info *Student, age int) {
	info.Age = age
}

// json串字母默认小写，显式指定json串key
type Student1 struct {
	Name string `json:"NAME666666"`
	Age  int    `json:"AGE666666"`
}
type Student2 struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"` //空值省略
}

func func5() {
	var p = People{"2024_4_24"}
	var s Student
	s = Student{People: p, Name: "myName", Age: 20}
	s = Student{p, "myName", 20}
	s.PrintInfo()
	s.Info()                   // 可以调用父结构体的方法
	fmt.Println(s.People.Time) // 调用父结构体的属性
	fmt.Println(s.Time)        // 也可以这样
	SetAge(s, 18)
	fmt.Println(s.Age)
	SetAge1(&s, 17)
	fmt.Println(s.Age)
	//
	s1 := Student1{
		Name: "myName",
		Age:  20,
	}
	byteData, _ := json.Marshal(s1)
	fmt.Println(string(byteData))
	s2 := Student2{
		Name: "myName",
		Age:  0, // 空值会被省略
	}
	// byteDate1需要是新变量，不能用byteData
	byteData1, _ := json.Marshal(s2)
	fmt.Println(string(byteData1))
}
