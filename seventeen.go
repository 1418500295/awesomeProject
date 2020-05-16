package main

import (
	"fmt"
	"os"
)

func main() {
	//追加写入文件
	s := "C:\\Users\\ASUS\\go\\src\\awesomeProject\\cases\\a.txt"
	f, _ := os.OpenFile(s,os.O_APPEND|os.O_CREATE|os.O_WRONLY,777)
	c, _ := f.Write([]byte("今天天气不错"))
	fmt.Println(c)

}
