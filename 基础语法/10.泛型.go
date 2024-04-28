package main

import (
	"encoding/json"
	"fmt"
)

func func10() {
	泛型函数()
	泛型结构体()
	泛型切片()
	泛型map()
}

func template_add[T int | float64 | int32](a, b T) T {
	return a + b
}
func 泛型函数() {
	fmt.Println(template_add(4, 3))
}

func 泛型结构体() {
	type Response[T any] struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data T      `json:"data"`
	}
	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"超级无敌大王"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)
	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"超级无敌大王","age":24}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}
func 泛型切片() {
	// 定义 泛型切片类型 type TypeName Type
	type myint int32
	type MySlice[T any] []T
	// 声明
	var mySlice MySlice[string]
	mySlice = append(mySlice, "超级无敌大王")
	var intSlice MySlice[int]
	intSlice = append(intSlice, 2)
}
func 泛型map() {
	type mymap map[string]int
	type MyMap[K string | int, V any] map[K]V
}
