package main

import (
	"fmt"
	"time"
)

/**
channel用于goroutines之间的通信
，让它们之间可以进行数据交换。像管道一样，一个goroutine_A向channel_A中放数据，
另一个goroutine_B从channel_A取数据。
 */

func main() {
	ch := make(chan string)
	go send(ch)
	go recv(ch)
	time.Sleep(2000)
}

func recv(ch chan string) {
	var x string
	for {
		x = <- ch
		fmt.Println(x)
	}

}

func send(ch chan string) {
	ch <- "呵呵"
	ch <- "哈哈"
	ch <- "测试"

}
























