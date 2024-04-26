package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

/*
向上抛
将错误交给上一级处理
一般是用于框架层，有些错误框架层面不能擅做决定，将错误向上抛不失为一个好的办法
*/
// Parent /* This is Parent
func Parent() error {
	err := method() // 遇到错误向上抛
	return err
}
func method() error {
	return errors.New("出错了")
}
func 向上抛() {
	fmt.Println(Parent())
}

/*
中断程序
遇到错误直接停止程序
这种一般是用于初始化，一旦初始化出现错误，程序继续走下去也意义不大了，还不如中断掉
*/
/*
在Go语言中，init 函数是一个特殊的函数，它在包初始化时自动被调用。
每个包可以有多个 init 函数，它们按照包中源文件的顺序被调用。init 函数不能被显式地调用，也不能被引用。
*/
/*
specialization 特化
specified 明确规定
*/
/*func init() {
	// 读取配置文件中，结果路径错了
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}*/
func 中断程序() {
	fmt.Println("啦啦啦")
}

/*
恢复程序
我们可以在一个函数里面，使用一个defer，可以实现对panic的捕获
以至于出现错误不至于让程序直接崩溃
这种一般也是框架层的异常处理所做的
*/
/*
在Go程序中，所有的goroutine都是并发执行的，但是 main 函数所在的goroutine具有特殊的地位。
如果 main 函数返回，或者在 main 函数中发生panic且没有被recover，那么程序将结束执行。
其他goroutine可能会继续运行，但最终也会被强制结束，因为程序的主goroutine已经退出。
*/
func read() {
	// 在Go语言中，panic 会导致当前函数立即停止执行，并开始向上搜索最近的 defer 语句。
	// 所以，不能放在异常位置的后面
	defer func() {
		// 短变量声明（:=）来声明 err 变量，并且将 recover() 函数的返回值赋给了这个变量
		if err := recover(); err != nil {
			fmt.Println("捕获异常，打印错误信息:\n", err) // 捕获异常，打印错误信息
			// 打印错误的堆栈信息
			s := string(debug.Stack())
			fmt.Println("打印错误的堆栈信息:\n", s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) // 肯定会有一个panic
}

//Process finished with the exit code 2 表示程序以退出码2结束。
//这个退出码是Go语言运行时用来表示程序遇到了一个未被捕获的 panic。
//当 panic 发生且没有被 recover 捕获时，Go运行时会打印出 panic 的错误信息和堆栈跟踪，然后以非零退出码结束程序。
/*
当然，这个用于捕获异常的defer的延迟函数可以在调用链路上的任何一个函数上
一般用于在最上层函数，捕获所有异常
*/
func 恢复程序() {
	read()
	fmt.Println("继续执行。。。")
}
func func9() {
	向上抛()
	中断程序()
	恢复程序()
}
