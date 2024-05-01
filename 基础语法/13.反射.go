package main

import (
	"fmt"
	"reflect"
)

/*
如果是写一下框架，偏底层工具类的操作
不用反射确实不太好写，但是如果是在业务上，大量使用反射就不太合适了
因为反射的性能没有正常代码高，会慢个一到两个数量级
使用反射可读性也不太好，并且也不能在编译期间发生错误
*/
func func13() {
	类型判断()
	通过反射获取值()
	通过反射修改值()

	// todo 结构体反射 修改结构体中某些值 调用结构体方法 orm小案例
}
func refSetValue(obj any) {
	value := reflect.ValueOf(obj)
	elem := value.Elem()
	// 专门取指针反射的值
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("奇塔")
	}
}
func 通过反射修改值() {
	name := "寇尔芙"
	refSetValue(&name)
	fmt.Println(name)
}
func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	// 去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}
func 类型判断() {
	refType(struct {
		Name string
	}{Name: "莫辛纳甘"})
	name := "莫辛纳甘"
	refType(name)
	refType([]string{"莫辛纳甘"})
}
func refValue(obj any) {
	value := reflect.ValueOf(obj)
	// Type()方法返回的是reflect.Type类型的值，它代表了reflect.Value所持有的值的静态类型信息。
	fmt.Println(value, value.Type())
	// Kind()方法返回的是一个枚举值，它描述了值的类型类别，而不是值的完整类型信息。
	switch value.Kind() {
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())
	}
}
func 通过反射获取值() {
	refValue(struct {
		Name string
		age  int32
	}{"娜美西斯", 20})
}
