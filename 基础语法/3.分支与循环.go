package main

import (
	"fmt"
	"time"
)

func func3() {
	//分支()
	循环()
}
func 分支() {
	/*
	* 	<=0     未出生
	* 	1-18    未成年
	* 	18-35   青年
	* 	>=35    中年
	 */
	//fmt.Println("请输入你的年龄：")
	//var age int = 20
	//fmt.Scan(&age)
	//if age <= 0 {
	//	fmt.Println("未出生")
	//	return
	//}
	//if age <= 18 {
	//	fmt.Println("未成年")
	//	return
	//}
	//if age <= 35 {
	//	fmt.Println("青年")
	//	return
	//}
	//fmt.Println("中年")
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	// 嵌套式
	if age <= 18 {
		if age <= 0 {
			fmt.Println("未出生")
		} else {
			fmt.Println("未成年")
		}
	} else {
		if age <= 35 {
			fmt.Println("青年")
		} else {
			fmt.Println("中年")
		}
	}
	// 多条件式
	if age <= 0 {
		fmt.Println("未出生")
	}
	if age > 0 && age <= 18 {
		fmt.Println("未成年")
	}
	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}
	fmt.Println("请输入星期数字：")
	var week int
	fmt.Scan(&week)
	fmt.Println("今天是：")
	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
	//go的switch的多选一，满足其中一个结果之后，就结束switch了
	//fallthrough
	fmt.Println("本周还有：")
	switch week {
	case 1:
		fmt.Println("周一")
		fallthrough
	case 2:
		fmt.Println("周二")
		fallthrough
	case 3:
		fmt.Println("周三")
		fallthrough
	case 4:
		fmt.Println("周四")
		fallthrough
	case 5:
		fmt.Println("周五")
		fallthrough
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
}
func 循环() {
	//for 初始化;条件;操作{
	//}
	//没有前置++
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("0到100的和：", sum)
	//死循环 for{}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 年月日时分秒的固定格式
		break
	}
	//for 模拟 while
	i := 0
	for i < 5 {
		print(i)
		i++
	}
	//for 模拟 do while
	for {
		i--
		print(i)
		if i == 0 {
			break
		}
	}

	// 遍历切片
	// 第一个参数是索引，第二个参数是对应的值
	s := []string{"枫枫", "知道"}
	for index, s2 := range s {
		fmt.Println(index, s2)
	}
	//遍历map
	//第一个参数就是key，第二个就是value
	s1 := map[string]int{
		"age":   24,
		"price": 1000,
	}
	for key, val := range s1 {
		fmt.Println(key, val)
	}
	// break用于跳出当前循环
	// continue 跳过本轮循环
}
