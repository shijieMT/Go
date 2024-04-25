package main

import (
	"errors"
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}
func add(n1 int, n2 int) {
	fmt.Println(n1, n2)
}

// 参数类型一样，可以合并在一起
func add1(n1, n2 int) {
	fmt.Println(n1, n2)
}

// 多个参数
func add2(numList ...int) {
	fmt.Println(numList)
}

// 无返回值
func fun1() {
	return // 也可以不写
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	//return // 相当于先定义再赋值
	return "abc"
}
func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int {
		var sum int
		for _, i2 := range numList {
			sum += i2
		}
		return sum
	}
}
func func4() {
	sayHello()
	add(1, 2)
	add1(1, 2)
	add2(1, 2)
	add2(1, 2, 3, 4)
	var a, b = fun3()
	fmt.Println(a, b)
	fmt.Printf("%#v %#v\n", a, b)
	var str = fun4()
	fmt.Println(str)
	//匿名函数
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	//根据用户输入的不同，执行不同的操作
	fmt.Println("请输入要执行的操作：")
	fmt.Println(`1：登录
2：个人中心
3：注销`)
	var num int
	fmt.Scan(&num)
	var funcMap = map[int]func(){
		1: func() {
			fmt.Println("登录")
		},
		2: func() {
			fmt.Println("个人中心")
		},
		3: func() {
			fmt.Println("注销")
		},
	}
	/*
		设计一个函数，先传一个参数表示延时，后面再次传参数就是将参数求和
		例如
		fun(2)(1,2,3) // 延时2秒求1+2+3
	*/
	funcMap[num]()
	var t = awaitAdd(3)(1, 3, 5)
	fmt.Println(t)
	//
	num1 := 20
	fmt.Println(&num1)
	add3(&num1)
	fmt.Println(num1) // 成功修改 2
}
func add3(num *int) {
	fmt.Println(num) // 内存值是一样的
	*num = 2         // 这里的修改会影响外面的num
}
