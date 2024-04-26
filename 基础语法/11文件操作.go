package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"syscall"
)

// GetCurrentFilePath 获取当前文件路径
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
func 文件读取() {
	byteData, _ := os.ReadFile("./hello.txt")
	fmt.Println(string(byteData))
	//可以通过获取当前go文件的路径，然后用相对于当前go文件的路径去打开文件
	fmt.Println(GetCurrentFilePath())
	//分片读
	file, _ := os.Open("./hello.txt")
	defer file.Close()
	for {
		buf := make([]byte, 1)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", buf)
	}
}

const (
	O_RDONLY int = syscall.O_RDONLY // 只读
	O_WRONLY int = syscall.O_WRONLY // 只写
	O_RDWR   int = syscall.O_RDWR   // 读写

	O_APPEND int = syscall.O_APPEND // 追加
	O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
	O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 同步打开
	O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件
)

func 文件写入() {
	//直接写入
	/*
		// 如果文件不存在就创建
		os.O_CREATE|os.O_WRONLY
		// 追加写
		os.O_APPEND|os.O_WRONLY
		// 可读可写
		os.O_RDWR
	*/
	err := os.WriteFile("./input.txt", []byte("这是内容"), os.ModePerm)
	fmt.Println(err)
}
func 文件复制() {
	read, _ := os.Open("./hello.txt")
	write, _ := os.Create("./hello_copy.txt") // 默认是 可读可写，不存在就创建，清空文件
	n, err := io.Copy(write, read)
	fmt.Println(n, err)
}
func 目录操作() {
	dir, _ := os.ReadDir("./")
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.Name(), info.Size()) // 文件名，文件大小，单位比特
	}
}

func func11() {
	文件读取()
	文件写入()
	文件复制()
	目录操作()
}
