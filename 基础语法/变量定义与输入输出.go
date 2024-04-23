package main

import "fmt"

func func2() {
	// 先定义，再赋值
	var name string
	name = "name"
	fmt.Println(name)
	// var 变量名 类型 = "变量值"
	var userName string = "shijie"
	fmt.Println(userName)
	//定义在函数体（包括main函数）内的变量都是局部变量，定义了就必须使用
	var name1, name2, name3 string // 定义多个变量
	fmt.Println(name1, name2, name3)
	// 定义多个变量并赋值
	var a1, a2 = "a1", "a2"
	// 简短定义多个变量并赋值
	a2, a3 := "a2", "a3"
	fmt.Println(a1, a2, a3)
	//
	var (
		name4     string = "name4"
		userName4        = "userName4"
	)
	print(name4, userName4)
	// 定义就要赋值,不能再修改了
	const nameConst string = "nameConst"
	fmt.Println(nameConst)
	//
	name5 := fmt.Sprintf("%v", "你好")
	fmt.Println(name5)
	//
	fmt.Printf("%v\n", "你好")        // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", 123, 123) // 打印类型
	fmt.Printf("%d\n", 3)           // 整数
	fmt.Printf("%7.2f\n", 6.66)     // 小数
	fmt.Printf("%s\n", "666")       // 字符串
	fmt.Printf("%#v\n", "")         // 用go的语法格式输出，很适合打印空字符串
	//
	name6 := fmt.Sprintf("%v", "你好")
	fmt.Println(name6)
	//
	//fmt.Println("输入您的名字：")
	//var name7 string
	//fmt.Scan(&name7) // 这里记住，要在变量的前面加个&
	//fmt.Println("你输入的名字是", name7)
	//
	fmt.Println("输入您的名字：name_yourRealName")
	var name8 string
	fmt.Scanf("name_%v", &name8)
	fmt.Println(name8)
}
