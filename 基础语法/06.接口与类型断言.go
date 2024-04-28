package main

import (
	"encoding/json"
	"fmt"
)

func func6() {

	接口()
	类型断言()
	空接口()
	//使用技巧
	// 解析JSON()
}

// 接口是一组仅包含方法名、参数、返回值的未具体实现的方法的集合
/*
	实现接口：
	一个类型实现了接口的所有方法
	即实现了该接口
*/
// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println(c.Name + " 唱")
}

func (c Chicken) jump() {
	fmt.Println(c.Name + " 跳")
}

func (c Chicken) rap() {
	fmt.Println(c.Name + " rap")
}

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// 全部实现完之后，chicken就不再是一只普通的鸡了
func 接口() {

	var animal Animal
	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()
}

func 解析JSON() {
	// 假设我们有一个JSON字符串
	jsonStr := `{"name": "Alice", "age": 30, "active": true, "skills": ["coding", "reading"]}`

	// 解析JSON字符串到map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		panic(err)
	}

	// 访问map中的值
	name := data["name"].(string)
	// 此处改 float64 为 int 会报错
	// https://blog.csdn.net/u013474436/article/details/109177763
	age := data["age"].(float64)
	active := data["active"].(bool)
	skills := data["skills"].([]interface{})

	// 输出结果
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Active:", active)
	fmt.Println("Skills:", skills)
}

func judge_animal(obj Animal) {
	// 通过断言来获取此时的具体类型
	switch obj.(type) {
	case Chicken:
		fmt.Println("鸡")
	case Cat:
		fmt.Println("猫")
	}
	obj.sing()
}

// Cat 需要全部实现这些接口
type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Println(c.Name + " 唱")
}
func (c Cat) jump() {
	fmt.Println(c.Name + " 跳")
}
func (c Cat) rap() {
	fmt.Println(c.Name + " rap")
}
func 类型断言() {
	var c = Cat{"叮当猫"}
	judge_animal(c)
}

//

func PrintAnyValue1(v interface{}) {
	fmt.Println(v)
}
func PrintAnyValue2(v any) {
	fmt.Println(v)
}
func 空接口() {
	// interface{}
	PrintAnyValue1("123")
	PrintAnyValue2("123")
}
