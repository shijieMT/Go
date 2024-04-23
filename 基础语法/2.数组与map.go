package main

import (
	"fmt"
	"sort"
)

func func2() {
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	var array2 = [...]int{1, 2, 3}
	fmt.Println(array2)
	//如果要修改某个值，只能根据索引去找然后替换
	array1[0] = 10 // 根据索引找到对应的元素位置，然后替换
	fmt.Println(array1)
	// 有的语言支持负向索引，go不支持
	/*
	   a      b     c     d
	   -4    -3    -2    -1
	*/
	// 以定义一个字符串数组 a b c d 为例
	var sList = []string{"a", "b", "c", "d"}
	// 拿到倒数第二个元素
	fmt.Println(sList[len(sList)-2]) // 拿到a这个元素
	// 定义一个字符串切片
	var list []string
	list = append(list, "Go")
	list = append(list, "++")
	fmt.Println(list)
	fmt.Println(len(list)) // 切片长度
	// 修改第二个元素
	list[1] = "不知道"
	fmt.Println(list)
	var list2 []uint16
	list2 = append(list2, 66)
	fmt.Println(list2)
	//除了基本数据类型，其他数据类型如果只定义不赋值，那么实际的值就是nil
	// 定义一个字符串切片
	var list3 []string
	fmt.Println(list3 == nil) // true
	//make([]type, length, capacity) 可以通过make函数创建指定长度，指定容量的切片
	// 定义一个字符串切片
	var list4 = make([]string, 0)
	fmt.Println(list4, len(list4), cap(list4))
	fmt.Println(list4 == nil) // false

	list1 := make([]int, 2, 2)
	fmt.Println(list1, len(list1), cap(list1))
	//
	var list5 = [...]string{"a", "b", "c"}
	slices := list5[:] // 左一刀，右一刀  变成了切片
	fmt.Println(slices)
	fmt.Println(list5[1:2]) // b
	//切片排序
	var list6 = []int{4, 5, 3, 2, 7}
	fmt.Println("排序前:", list6)
	sort.Ints(list6)
	fmt.Println("升序:", list6)
	sort.Sort(sort.Reverse(sort.IntSlice(list6)))
	fmt.Println("降序:", list6)
	//Go语言中的map(映射、字典)是一种内置的数据结构，它是一个无序的key-value对的集合
	//map的key必须是基本数据类型，value可以是任意类型
	// 声明
	var m1 map[string]string
	// map使用之前一定要初始化
	//初始化1
	m1 = make(map[string]string)
	// 初始化2
	m1 = map[string]string{}
	// 设置值
	m1["name"] = "myName"
	fmt.Println(m1)
	// 取值
	fmt.Println(m1["name"])
	// 删除值
	delete(m1, "name")
	fmt.Println(m1)
	// 声明并赋值
	var m2 = map[string]string{}
	fmt.Println(m2)
	var m3 = make(map[string]string)
	fmt.Println(m3)
	//如果只有一个参数接，那这个参数就是值，如果没有，这个值就是类型的零值
	//如果两个参数接，那第二个参数就是布尔值，表示是否有这个元素
	// 声明并赋值
	var m4 = map[string]int{
		"age": 21,
	}
	age1 := m4["age1"] // 取一个不存在的
	fmt.Println(age1)
	age2, ok := m4["age1"]
	fmt.Println(age2, ok)
}
