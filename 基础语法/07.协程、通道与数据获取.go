package main

import (
	"fmt"
	"sync"
	"time"
)

func sing1() {
	fmt.Println("a-a/a-a_a_a__a___")
	wait.Done()
}
func sing2() {
	fmt.Println("a-a/a-a_a_a__a___")
}

var (
	wait = sync.WaitGroup{}
)

func 协程() {

	wait.Add(4)
	go sing1()
	go sing1()
	go sing1()
	go func() {
		defer wait.Done() // 在 goroutine 结束时减少等待组的计数
		sing2()
	}()
	wait.Wait()

}

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道
func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan <- money
	fmt.Println("金额发送完成，数值：", money)
	wait.Done()
}
func 通道() {
	//在协程里面产生了数据，如何传递给主线程呢？
	var c chan int // 声明一个传递整形的通道
	// 初始化通道
	c = make(chan int, 1) //  初始化一个 有一个缓冲位的通道
	c <- 1
	//c <- 2 // 会报错 deadlock
	fmt.Println(<-c) // 取值
	//fmt.Println(<-c) // 再取也会报错  deadlock

	c <- 2
	n, ok := <-c
	fmt.Println(n, ok)
	close(c) // 关闭通道
	//c <- 3   // 关闭之后就不能再写或读了  send on closed channel
	fmt.Println(c)
	//
	var begin = time.Now()
	wait.Add(3)
	go pay("超级无敌购物大王", 100, &wait)
	go pay("超级无敌购物小王", 10, &wait)
	go pay("超级无敌购物中王", 50, &wait)
	go func() {
		defer close(moneyChan)
		// 在协程函数里面等待上面协程函数结束
		wait.Wait()
	}()
	var wait1 = sync.WaitGroup{}
	wait1.Add(1)
	go func() {
		defer wait1.Done()
		for {
			money, ok := <-moneyChan
			fmt.Println("金额接受完成，金额为：", money)
			if !ok {
				break
			}
		}
	}()
	wait1.Wait()
	fmt.Println(time.Since(begin))
}

var moneyChan1 = make(chan int)    // 声明并初始化一个长度为0的信道
var nameChan1 = make(chan string)  // 声明并初始化一个长度为0的信道
var doneChan = make(chan struct{}) // 声明一个用于关闭的信道

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

func 数据获取select() {
	var wait sync.WaitGroup
	startTime := time.Now()
	// 现在的模式，就是购物接力
	//shopping("张三")
	//shopping("王五")
	//shopping("李四")
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
	go send("张三", 2, &wait)
	go send("王五", 3, &wait)
	go send("李四", 5, &wait)

	//go func() {
	//	defer close(doneChan)
	//	defer time.Sleep(5 * time.Second)
	//	defer close(moneyChan1)
	//	defer close(nameChan1)
	//	wait.Wait()
	//}()
	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)

		wait.Wait()
	}()
	/*
		在Go语言中，close函数用于关闭一个通道。当你关闭一个通道时，它会向通道发送一个特殊的值，这个值被称为通道的关闭值。
		对于任何类型的通道，关闭值都是该类型的零值。对于doneChan这样的空结构体通道，关闭值是struct{}类型的零值，即struct{}本身。
		关闭通道的行为与发送一个值到通道是不同的。
		当你关闭一个通道时，所有等待从该通道接收数据的select语句都会立即执行case分支，而不是等待一个实际的值被发送。
		这意味着，即使通道中没有数据，关闭通道也会导致select语句中的case分支被执行。
	*/
	var moneyList []int
	var nameList []string
	// select随机选取避免饥饿
	var event = func() {
		for {
			time.Sleep(1 * time.Microsecond) // 因为随机选取，可能会有0值
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				return

			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)

	go event1()
	select {
	case <-done:
		fmt.Println("协程执行完毕")
	case <-time.After(1 * time.Second):
		fmt.Println("超时")
		//return
	}
}

var done = make(chan struct{})

func event1() {
	fmt.Println("event执行开始")
	time.Sleep(2 * time.Second)
	fmt.Println("event执行结束")
	close(done)
}

func func7() {
	//协程()
	//通道()
	数据获取select()
}
