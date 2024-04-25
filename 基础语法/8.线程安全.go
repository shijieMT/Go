package main

import (
	"fmt"
	"sync"
	"time"
)

var num_func8 int
var wait_func8 sync.WaitGroup
var lock_func8 sync.Mutex

func add_func8() {
	lock_func8.Lock()
	for i := 0; i < 1000000; i++ {
		num_func8++
	}
	lock_func8.Unlock()

	wait_func8.Done()
}
func reduce_func8() {
	lock_func8.Lock()
	for i := 0; i < 1000000; i++ {
		num_func8--
	}
	lock_func8.Unlock()
	wait_func8.Done()
}
func 线程安全() {
	var t = time.Now()
	wait_func8.Add(2)
	go add_func8()
	go reduce_func8()
	wait_func8.Wait()
	fmt.Println(time.Since(t))
	fmt.Println(num_func8)
}

/*
我们不能在并发模式下读写map
如果要这样做
给读写操作加锁
使用sync.Map
加锁
*/
var mp = map[string]string{}

/*
如果我们在一个协程函数下，读写map就会引发一个错误
concurrent map read and map write
*/
func reader() {
	for i := 0; i < 3000000; i++ {
		fmt.Println(mp["time"])
	}
	wait_func8.Done()
}
func writer() {
	for i := 0; i < 3000000; i++ {
		mp["time"] = time.Now().Format("15:04:05")
	}
	wait_func8.Done()
}

/*
sync.Map 可以存储任意类型的键（key）和值（value）。在Go语言中，sync.Map 是一个并发安全的map，它允许你存储和检索键值对，而不需要担心并发访问的问题。
*/

var mp_sync = sync.Map{}

func reader_sync() {
	for {
		fmt.Println(mp_sync.Load("time"))
	}
	wait.Done()
}
func writer_sync() {
	for {
		mp_sync.Store("time", time.Now().Format("15:04:05"))
	}
	wait.Done()
}

func 线程安全下的map() {
	//wait_func8.Add(2)
	//go writer()
	//go reader()
	//wait_func8.Wait()

	wait_func8.Add(2)
	go reader_sync()
	go writer_sync()
	wait_func8.Wait()
}
func func8() {
	//线程安全()
	线程安全下的map()
}
