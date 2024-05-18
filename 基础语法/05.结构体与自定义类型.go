package main

import (
	"encoding/json"
	"fmt"
)

/*
在 Go 语言中，方法（method）是一种特殊的函数，它与某个类型（即接收器类型）相关联。
方法允许我们将函数与特定的类型绑定在一起，使得该函数可以访问和操作该类型的实例。

func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error)
这是一个合法的方法声明，其中 GetVideoLogic 是一个结构体类型，l 是接收器变量名。
该方法使用指针接收器，可以修改接收器指向的实例的值。

func (l GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error)
这也是一个合法的方法声明，其中 GetVideoLogic 是一个结构体类型，l 是接收器变量名。
该方法使用值接收器，对接收器的修改不会影响原始实例的值。

func (GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error)
这不是一个合法的方法声明，因为缺少接收器变量名。
在 Go 语言中，方法声明必须指定一个接收器变量名，即使在方法内部不使用该变量。
*/

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

func 结构体() {
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

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7 // 校验失败的错误
	ServiceErrCode Code = 8 // 服务错误
)

func (c Code) GetMsg() string {
	if c == SuccessCode {
		return "Success"
	}
	if c == ValidCode {
		return "Valid"
	}
	if c == ServiceErrCode {
		return "Service"
	}
	return "error"
}

type myint int
type mystring string

/*
类型别名：
1. 不能绑定方法
2。 打印类型还是原始类型
3. 和原始类型比较，类型别名不用转换
*/

type AliasCode = int
type MyCode int

const (
	SuccessAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
//func (m AliasCode) MyAliasCodeMethod() {
//
//}

func 自定义类型() {
	fmt.Println(SuccessCode.GetMsg())
	// 类型别名，打印它的类型还是原始类型
	fmt.Printf("%T %T \n", SuccessCode, SuccessAliasCode) // main.Code int
	// 可以直接和原始类型比较
	var i int
	fmt.Println(SuccessAliasCode == i)
	fmt.Println(int(SuccessCode) == i) // 必须转换之后才能和原始类型比较
}
func func5() {
	结构体()
	自定义类型()
}
