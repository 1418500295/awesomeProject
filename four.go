package main

import "fmt"

func main() {
	/**
	无限循环
	 */
	//for true {
	//	fmt.Print("这是无限循环")
	//}
	/**
	计算1到10之间数字的和
	 */
	sum := 0
	for i := 0 ;i <= 10;i++ {
		//fmt.Println(sum)
		sum += i
	}
	fmt.Println(sum)

	
}
