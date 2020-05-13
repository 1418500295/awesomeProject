package main

import (
	"fmt"
	"time"
)
/**
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，
以一个不同的、新创建的 goroutine 来执行一个函数。
同一个程序中的所有 goroutine 共享同一个地址空间。
 */
func say(s string)  {
	for i:=0;i<10;i++{
		time.Sleep(1000)
		fmt.Println(s)
	}

}
func main() {

	go say("hello")
	say("world")

}
